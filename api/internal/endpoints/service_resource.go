// Copyright 2022 Cisco Systems, Inc. and its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package endpoints

import (
	"context"
	"errors"
	"fmt"
	"github.com/cisco-developer/api-insights/api/internal/access"
	"github.com/cisco-developer/api-insights/api/internal/db"
	"github.com/cisco-developer/api-insights/api/internal/middleware"
	"github.com/cisco-developer/api-insights/api/internal/models"
	modelsanalyzer "github.com/cisco-developer/api-insights/api/internal/models/analyzer"
	"github.com/cisco-developer/api-insights/api/internal/models/diff"
	"github.com/cisco-developer/api-insights/api/pkg/analyzer"
	"github.com/cisco-developer/api-insights/api/pkg/apiclarity"
	apiclarityclient "github.com/cisco-developer/api-insights/api/pkg/apiclarity/client"
	"github.com/cisco-developer/api-insights/api/pkg/differ"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strconv"
	"time"
)

type serviceResource struct {
	config           *shared.AppConfig
	dao              db.ServiceDAO
	validate         *validator.Validate
	specDAO          db.SpecDAO
	specDiffDAO      db.SpecDiffDAO
	specAnalysisDAO  db.SpecAnalysisDAO
	analyzerDAO      db.AnalyzerDAO
	organizationDAO  db.OrganizationDAO
	analyzerSvc      analyzer.Service
	differSvc        differ.Service
	apiclarityClient *apiclarityclient.APIClarityAPIs
	accessChecker    access.Checker
	info             *models.Info
}

const (
	mimeTextMarkdown = "text/markdown"
	mimeTextPlain    = "text/plain"
)

// Register the API
// prefix: /v1/apiregistry/services
func (r *serviceResource) Register(config *shared.AppConfig, container *restful.Container, prefix string) {
	ws := &restful.WebService{}
	ws.Path(prefix).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON).ApiVersion(config.AppVersion).Doc("APIs for Service.")
	ws.Filter(middleware.ResourceAccessChecker(r.accessChecker))

	var service models.Service
	var services []models.Service
	var servicePatch models.ServicePatch
	var specAnalyses []models.SpecAnalysis
	var specDoc models.SpecDoc
	var id = ws.PathParameter("id", "unique identifier (UUID or Name ID) for service.").DataType("string")
	var oldSpecID = ws.PathParameter("oldSpecID", "old spec ID").DataType("string")
	var newSpecID = ws.PathParameter("newSpecID", "new spec ID").DataType("string")
	var specDiffFormat = ws.QueryParameter("format", "format of spec diff result").DataType("string").DefaultValue("json")
	var withDoc = ws.QueryParameter("withDoc", "flag indicating if doc should be included, by default false").DataType("boolean").DefaultValue("false")
	var download = ws.QueryParameter("download", "flag indicating if response content is to be served as a downloadable attachment (Content-Disposition), by default true").DataType("boolean").DefaultValue("true")
	var withFindings = ws.QueryParameter("withFindings", "flag indicating if result findings should be included, by default false").DataType("boolean").DefaultValue("false")
	var tags = ws.QueryParameter("tags", "tags for getting services").DataType("string")
	var q = ws.QueryParameter("q", "searching criteria for services").DataType("string")
	var limit = ws.QueryParameter("limit", "max items to return at one time").DataType("string")
	var offset = ws.QueryParameter("offset", "starting offset").DataType("string")

	ws.Route(
		ws.GET("").
			To(r.list).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteParams(offset)).
			Do(shared.RouteParams(limit, max)).
			Do(shared.RouteReturns(services, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteWrites(services)).
			Do(shared.RouteParams(q)).
			Do(shared.RouteParams(tags)).
			Do(shared.RouteParams(sort, sortOrder)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"service"}).
			Notes("List all services with specified tags."))

	ws.Route(
		ws.GET("/{id}").
			To(r.get).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(service, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError)).
			Do(shared.RouteWrites(service)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"service"}).
			Notes("Get a service with specified id"))

	ws.Route(
		ws.POST("").
			To(r.save).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(service, http.StatusOK, http.StatusCreated)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusInternalServerError)).
			Do(shared.RouteReads(service, "service"), shared.RouteWrites(service)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"service"}).
			Notes("Create a new service"))

	ws.Route(
		ws.PUT("/{id}").
			To(r.save).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(service, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteReads(service, "service"), shared.RouteWrites(service)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"service"}).
			Notes("Update existing service"))

	ws.Route(
		ws.PATCH("/{id}").
			To(r.patch).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(servicePatch, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteReads(servicePatch, "service"), shared.RouteWrites(service)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"service"}).
			Notes("Patch existing service"))

	ws.Route(
		ws.DELETE("/{id}").
			To(r.delete).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(nil, http.StatusNoContent)).
			Do(shared.RouteReturns(&se, http.StatusNotFound, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteWrites(nil)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"service"}).
			Notes("Delete existing service"))

	//
	// Non-CRUD routes
	//

	var (
		spec            models.Spec
		specs           []models.Spec
		specAnalysisReq models.SpecAnalysisRequest
		specAnalysisRes models.SpecAnalysisResponse
		specDiff        models.SpecDiff
		specDiffReq     models.SpecDiffRequest
		specID          = ws.PathParameter("specID", "unique identifier for service spec.").DataType("string")
		specTags        = ws.QueryParameter("tags", "tags for getting service specs").DataType("string")
		specQ           = ws.QueryParameter("q", "searching criteria for service specs").DataType("string")
	)

	ws.Route(
		ws.POST("/{id}/specs").
			To(r.saveSpec).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(spec, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteReads(spec, "spec"), shared.RouteWrites(spec)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"spec"}).
			Notes("Create a new service spec"))

	ws.Route(
		ws.GET("/{id}/specs").
			To(r.listSpecs).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteParams(offset)).
			Do(shared.RouteParams(limit, max)).
			Do(shared.RouteReturns(specs, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteWrites(specs)).
			Do(shared.RouteParams(specQ)).
			Do(shared.RouteParams(specTags)).
			Do(shared.RouteParams(id)).
			Do(shared.RouteParams(withDoc)).
			Do(shared.RouteParams(sort, sortOrder)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"spec"}).
			Notes("List all service specs with specified tags."))

	ws.Route(
		ws.GET("/{id}/specs/{specID}").
			To(r.getSpec).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(spec, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError)).
			Do(shared.RouteWrites(spec)).
			Do(shared.RouteParams(id)).
			Do(shared.RouteParams(specID)).
			Do(shared.RouteParams(withDoc)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"spec"}).
			Notes("Get a service spec with specified id"))

	ws.Route(
		ws.DELETE("/{id}/specs/{specID}").
			To(r.deleteSpec).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(nil, http.StatusNoContent)).
			Do(shared.RouteReturns(&se, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError)).
			Do(shared.RouteWrites(spec)).
			Do(shared.RouteParams(id)).
			Do(shared.RouteParams(specID)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"spec"}).
			Notes("Delete existing service spec"))

	ws.Route(
		ws.GET("/{id}/specs/{specID}/doc").
			To(r.getSpecDoc).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(specDoc, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError)).
			Do(shared.RouteWrites(specDoc)).
			Do(shared.RouteParams(id)).
			Do(shared.RouteParams(specID)).
			Do(shared.RouteParams(download)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"spec"}).
			Notes("Get a service spec doc with specified id").
			Produces(restful.MIME_JSON, mimeTextPlain))

	ws.Route(
		ws.POST("/{id}/specs/{specID}/analyses").
			To(r.createAnalysis).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(specAnalysisRes, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteReads(specAnalysisReq, "spec analysis request"), shared.RouteWrites(specAnalysisRes)).
			Do(shared.RouteParams(id)).
			Do(shared.RouteParams(specID)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"spec"}).
			Notes("Create a new service spec analysis"))

	ws.Route(
		ws.GET("/{id}/specs/{specID}/analyses").
			To(r.getAnalyses).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteParams(offset)).
			Do(shared.RouteParams(limit, max)).
			Do(shared.RouteReturns(specAnalyses, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError)).
			Do(shared.RouteWrites(specAnalyses)).
			Do(shared.RouteParams(id)).
			Do(shared.RouteParams(specID)).
			Do(shared.RouteParams(sort, sortOrder)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"spec"}).
			Notes("Get the service spec analyses (report)"))

	ws.Route(
		ws.GET("/{id}/specs/analyses").
			To(r.getServiceAnalyses).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteParams(offset)).
			Do(shared.RouteParams(limit, max)).
			Do(shared.RouteReturns(specAnalyses, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusUnauthorized, http.StatusForbidden, http.StatusNotFound, http.StatusInternalServerError)).
			Do(shared.RouteWrites(specAnalyses)).
			Do(shared.RouteParams(id)).
			Do(shared.RouteParams(specID)).
			Do(shared.RouteParams(withFindings)).
			Do(shared.RouteParams(sort, sortOrder)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"spec"}).
			Notes("List all the service analyses (reports)"))

	ws.Route(
		ws.POST("/{id}/specs/diff").
			To(r.createDiff).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(specDiff, http.StatusOK, http.StatusCreated)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteReads(specDiffReq, "spec diff request"), shared.RouteWrites(specDiff)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"spec"}).
			Notes("Create a service spec diff"))

	ws.Route(
		ws.GET("/{id}/specs/diff/{oldSpecID}/{newSpecID}").
			To(r.getOrCreateDiff).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns([]byte{}, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteWrites([]byte{})).
			Do(shared.RouteParams(id)).
			Do(shared.RouteParams(oldSpecID)).
			Do(shared.RouteParams(newSpecID)).
			Do(shared.RouteParams(specDiffFormat)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"spec"}).
			Notes("Create a service spec diff").
			Produces(restful.MIME_JSON, mimeTextMarkdown),
	)

	ws.Route(
		ws.POST("/{id}/specs/reconstruct").
			To(r.reconstructSpec).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(spec, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteWrites(spec)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"spec"}).
			Notes("Reconstruct a new service spec"))

	container.Add(ws)
}

func (r *serviceResource) list(req *restful.Request, res *restful.Response) {
	shared.LogDebugf("get request to list services.")

	var filter = &db.ListFilter{Model: &models.Service{}}
	err := filter.From(req)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	services, err := r.dao.List(req.Request.Context(), filter, orgServiceAccessDataFilterFromReq(req))
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	shared.LogDebugf("total %v service(s) returned", len(services))
	_ = res.WriteEntity(services)
}

func (r *serviceResource) get(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("id")
	shared.LogDebugf("get request to retrieve service: %v", id)

	service, err := r.dao.Get(req.Request.Context(), id)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	_ = res.WriteEntity(service)
}

func (r *serviceResource) save(req *restful.Request, res *restful.Response) {
	service := &models.Service{}
	id := req.PathParameter("id")
	err := req.ReadEntity(service)

	if err != nil {
		shared.LogErrorf("failed to get service from body: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	shared.LogDebugf("get request to save service with body: %#v", service)

	now := time.Now().UTC()

	if id == "" {
		service.CreatedAt = now
	} else {
		existing, err := r.dao.Get(req.Request.Context(), id)
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		service.ID = existing.ID
		service.CreatedAt = existing.CreatedAt
	}
	service.UpdatedAt = now

	err = r.validate.Struct(service)
	if err != nil {
		shared.LogErrorf("failed to validate Service %s - %v", service.ID, err.Error())
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = r.dao.Save(req.Request.Context(), service, orgServiceAccessDataFilterFromReq(req))
	if err != nil {
		handleError(res, err)
		return
	}
	if id == "" {
		// Auto-create an organization, if it doesn't exist, for unauthenticated version.
		if !r.info.Auth.Enabled {
			if _, err := r.organizationDAO.Get(req.Request.Context(), service.OrganizationID); err != nil {
				err = r.organizationDAO.Save(req.Request.Context(), &models.Organization{
					ID:          shared.TimeUUID(),
					NameID:      service.OrganizationID,
					Title:       service.OrganizationID,
					Description: "[auto-created]",
					CreatedAt:   now,
					UpdatedAt:   now,
				}, nil)
				if err != nil {
					shared.LogErrorf("failed to auto-create Organization(%s) for Service %s - %v", service.OrganizationID, service.ID, err.Error())
					handleError(res, err)
					return
				}
			}
		}

		res.Header().Add("Location", "/v1/apiregistry/services/"+service.ID)
		_ = res.WriteHeaderAndEntity(http.StatusCreated, service)
	} else {
		_ = res.WriteEntity(service)
	}
}

func (r *serviceResource) patch(req *restful.Request, res *restful.Response) {
	patch := &models.ServicePatch{}
	id := req.PathParameter("id")
	err := req.ReadEntity(patch)

	if err != nil {
		shared.LogErrorf("failed to get servicePatch from body: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	shared.LogDebugf("get request to patch service with body: %#v", patch)

	service, err := r.dao.Get(req.Request.Context(), id)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	service.UpdatedAt = time.Now().UTC()

	if patch.AdditionalInfo != nil && len(*patch.AdditionalInfo) > 0 {
		if service.AdditionalInfo == nil {
			service.AdditionalInfo = map[string]interface{}{}
		}
		for k, v := range *patch.AdditionalInfo {
			service.AdditionalInfo[k] = v
		}
	}
	if patch.Contact != nil {
		service.Contact = patch.Contact
	}
	if patch.Description != nil {
		service.Description = *patch.Description
	}
	if patch.NameID != nil {
		service.NameID = *patch.NameID
	}
	if patch.OrganizationID != nil {
		service.OrganizationID = *patch.OrganizationID
	}
	if patch.ProductTag != nil {
		service.ProductTag = *patch.ProductTag
	}
	if patch.Title != nil {
		service.Title = *patch.Title
	}
	if patch.AnalyzersConfigs != nil && len(*patch.AnalyzersConfigs) > 0 {
		if service.AnalyzersConfigs == nil {
			service.AnalyzersConfigs = models.AnalyzerConfigMap{}
		}
		for k, v := range *patch.AnalyzersConfigs {
			service.AnalyzersConfigs[k] = v
		}
	}

	err = r.validate.Struct(service)
	if err != nil {
		shared.LogErrorf("failed to validate Service %s - %v", service.ID, err.Error())
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	err = r.dao.Save(req.Request.Context(), service, orgServiceAccessDataFilterFromReq(req))
	if err != nil {
		handleError(res, err)
		return
	}
	_ = res.WriteEntity(service)
}

func (r *serviceResource) delete(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("id")
	shared.LogDebugf("get request to delete service: %v", id)

	err := r.dao.Delete(req.Request.Context(), id)
	if err == db.ErrNotFound {
		res.WriteHeader(http.StatusNotFound)
	} else if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	} else {
		res.WriteHeader(http.StatusNoContent)
	}
}

// POST /{id}/specs
func (r *serviceResource) saveSpec(req *restful.Request, res *restful.Response) {
	var (
		serviceID = req.PathParameter("id")
		spec      = &models.Spec{}
	)
	shared.LogDebugf("get request to save service (%v) spec", serviceID)

	if err := req.ReadEntity(spec); err != nil {
		shared.LogErrorf("failed to get spec from body: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	vr, err := models.ValidateSpecDoc(req.Request.Context(), spec.Doc)
	if err != nil {
		shared.LogErrorf("failed to validate service (%v) spec: %#v", serviceID, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !vr.Valid {
		shared.LogErrorf("invalid service (%v) spec: %#v", serviceID, vr.Findings)
		_ = res.WriteHeaderAndEntity(http.StatusBadRequest, vr)
		return
	}

	if _, err := spec.LoadDocAsOAS(req.Request.Context(), false, true); err != nil {
		shared.LogErrorf("failed to load Spec.Doc as OAS from body: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	service, err := r.dao.Get(req.Request.Context(), serviceID)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	serviceID = service.ID
	spec.ID = shared.TimeUUID()
	spec.ServiceID = serviceID
	now := time.Now().UTC()
	spec.CreatedAt = now
	spec.UpdatedAt = now

	if err := r.validate.Struct(spec); err != nil {
		shared.LogErrorf("failed to validate Spec %s - %v", spec.ID, err.Error())
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if err := r.specDAO.Save(req.Request.Context(), spec); err != nil {
		handleError(res, err)
		return
	}

	go func() {
		if _, err := r.runSpecAnalysisRequest(context.Background(), res, &models.SpecAnalysisRequest{
			Spec:    spec,
			Service: service,
		},
			true, true); err != nil {
			shared.LogErrorf("failed to analyze service (%v) spec (%v): %#v", serviceID, spec.ID, err)
		}
	}()

	res.Header().Add("Location", "/v1/apiregistry/services/"+serviceID+"/specs/"+spec.ID)
	_ = res.WriteHeaderAndEntity(http.StatusCreated, spec)
}

// runSpecAnalysisRequest is a utility method that runs a SpecAnalysisRequest.
// Important to note that runSpecAnalysisRequest does write to res, so handle accordingly.
func (r *serviceResource) runSpecAnalysisRequest(ctx context.Context, res *restful.Response, specAnalysisReq *models.SpecAnalysisRequest, updateSpec, updateService bool) (*models.SpecAnalysisResponse, error) {
	activeAnalyzers, err := r.analyzerDAO.List(context.Background(), &db.ListFilter{Indexes: map[string]string{"status": modelsanalyzer.AnalyzerStatusActive}}, true)
	if err != nil {
		shared.LogErrorf("failed to list active analyzers: %s", err.Error())
		handleError(res, err)
		return nil, err
	}

	var sas []modelsanalyzer.SpecAnalyzer
	for _, a := range activeAnalyzers {
		sas = append(sas, modelsanalyzer.SpecAnalyzer(a.NameID))
	}

	// Apply specAnalysisReq.Analyzers filter: if specified, filter out non-`specAnalysisReq.Analyzers` values from sas.
	if len(specAnalysisReq.Analyzers) != 0 {
		reqAnalyzerSet := make(map[modelsanalyzer.SpecAnalyzer]struct{})
		for _, a := range specAnalysisReq.Analyzers {
			reqAnalyzerSet[a] = struct{}{}
		}
		i := 0
		for _, activeAnalyzer := range activeAnalyzers {
			analyzerName := modelsanalyzer.SpecAnalyzer(activeAnalyzer.NameID)
			if _, ok := reqAnalyzerSet[analyzerName]; ok {
				sas[i] = analyzerName
				i++
			}
		}
		sas = sas[:i]
	}
	specAnalysisReq.Analyzers = sas
	specAnalysisReq.ActiveAnalyzers = modelsanalyzer.ListToMap(activeAnalyzers)

	service := specAnalysisReq.Service

	// Merge specAnalysisReq.AnalyzersConfigs, service.AnalyzersConfigs, and []activeAnalyzers.Config.
	analyzersConfigs := models.AnalyzerConfigMap{}
	for _, a := range activeAnalyzers {
		if len(a.Config) > 0 {
			analyzersConfigs[modelsanalyzer.SpecAnalyzer(a.NameID)] = a.Config
		}
	}
	if len(service.AnalyzersConfigs) > 0 {
		analyzersConfigs.Merge(service.AnalyzersConfigs)
	}
	if len(specAnalysisReq.AnalyzersConfigs) > 0 {
		analyzersConfigs.Merge(specAnalysisReq.AnalyzersConfigs)
	}
	specAnalysisReq.AnalyzersConfigs = analyzersConfigs

	specAnalysisRes, err := r.analyzerSvc.Analyze(specAnalysisReq)
	if err != nil {
		shared.LogErrorf("failed to analyze service (%v) spec (%v): %#v", service.ID, specAnalysisReq.Spec.ID, err)
		res.WriteHeader(http.StatusInternalServerError)
		return nil, err
	}

	for _, specAnalysis := range specAnalysisRes.Results {
		if err := r.specAnalysisDAO.Save(ctx, specAnalysis); err != nil {
			handleError(res, err)
			return nil, err
		}
	}
	if updateSpec {
		if err := r.updateSpecScore(ctx, res, specAnalysisRes.SpecScore, specAnalysisReq.Spec, updateService, service); err != nil {
			return nil, err
		}
	}
	return specAnalysisRes, nil
}

// runSpecAnalysisRequest is a utility method that updates the spec score (and optionally, the service score as well).
// Important to note that updateSpecScore does write to res, so handle accordingly.
func (r *serviceResource) updateSpecScore(ctx context.Context, res *restful.Response, score int, spec *models.Spec, updateService bool, service *models.Service) error {
	spec.Score = &score
	now := time.Now().UTC()
	spec.UpdatedAt = now
	err := r.specDAO.Save(ctx, spec)
	if err != nil {
		handleError(res, err)
		return err
	}
	if updateService {
		service.UpdatedAt = now
		service.SetSummary(score, spec.Version, spec.Revision, now)
		err = r.dao.Save(ctx, service, nil)
		if err != nil {
			handleError(res, err)
			return err
		}
	}
	return nil
}

// GET /{id}/specs
func (r *serviceResource) listSpecs(req *restful.Request, res *restful.Response) {
	var (
		serviceID = req.PathParameter("id")
		withDoc   = parseQueryBool(req, "withDoc")
	)
	shared.LogDebugf("get request to list service (%v) specs", serviceID)

	s, err := r.dao.Get(req.Request.Context(), serviceID)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	serviceID = s.ID
	var filter = &db.ListFilter{Model: &models.Spec{}}
	if err := filter.From(req); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	filter.Indexes["service_id"] = serviceID
	specs, err := r.specDAO.List(req.Request.Context(), filter, withDoc)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	shared.LogDebugf("total %v service (%v) spec(s) returned", len(specs), serviceID)
	_ = res.WriteEntity(specs)
}

// GET /{id}/specs/{specID}
func (r *serviceResource) getSpec(req *restful.Request, res *restful.Response) {
	var (
		serviceID = req.PathParameter("id")
		specID    = req.PathParameter("specID")
	)
	// Default withDoc as true.
	withDoc, err := strconv.ParseBool(req.QueryParameter("withDoc"))
	if err != nil {
		withDoc = true
	}

	s, err := r.dao.Get(req.Request.Context(), serviceID)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	serviceID = s.ID
	shared.LogDebugf("get request to get service (%v) spec: %v", serviceID, specID)

	spec, err := r.specDAO.Get(req.Request.Context(), specID, withDoc)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	_ = res.WriteEntity(spec)
}

// DELETE /{id}/specs/{specID}
func (r *serviceResource) deleteSpec(req *restful.Request, res *restful.Response) {
	var (
		serviceID = req.PathParameter("id")
		specID    = req.PathParameter("specID")
	)
	shared.LogDebugf("get request to delete service (%v) spec: %v", serviceID, specID)

	err := r.specDAO.Delete(req.Request.Context(), specID)
	if err == db.ErrNotFound {
		res.WriteHeader(http.StatusNotFound)
	} else if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}

	// Service.Summary contains info about the latest spec,
	// and if the spec we're deleting happens to the latest spec,
	// then we need to update Service.Summary to either:
	//	- the info about the second-to-latest spec, if more than one specs exist, or
	//	- nil, if only a single spec exists.
	if specs, err := r.specDAO.List(req.Request.Context(), &db.ListFilter{
		Model: &models.Spec{},
		Limit: 2,
		Sorters: []*db.Sorter{{
			Order: db.OrderDesc,
			Field: "created_at",
		}}}, false); err != nil {
		shared.LogErrorf("failed to update service after deleting service (%v) spec (%v): %v", serviceID, specID, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	} else if len(specs) > 0 {
		service, err := r.dao.Get(req.Request.Context(), serviceID)
		if err != nil {
			res.WriteHeader(http.StatusNotFound)
			return
		}
		if specs[0].ID == specID {
			now := time.Now().UTC()
			service.UpdatedAt = now
			if len(specs) == 2 {
				if specs[1].Score != nil {
					service.SetSummary(*specs[1].Score, specs[1].Version, specs[1].Revision, specs[1].UpdatedAt)
				}
			} else if len(specs) == 1 {
				service.Summary = nil
			}
			err = r.dao.Save(req.Request.Context(), service, orgServiceAccessDataFilterFromReq(req))
			if err != nil {
				shared.LogErrorf("failed to update service after deleting service (%v) spec (%v): %v", serviceID, specID, err)
				handleError(res, err)
				return
			}
		}
	}

	// Delete all spec-related entities (e.g. spec analyses, spec diffs).
	if err := r.specAnalysisDAO.BatchDeleteBySpecID(req.Request.Context(), specID); err != nil {
		shared.LogErrorf("failed to batch delete spec analyses after deleting service (%v) spec (%v): %v", serviceID, specID, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.specDiffDAO.BatchDeleteBySpecID(req.Request.Context(), specID); err != nil {
		shared.LogErrorf("failed to batch delete spec diffs after deleting service (%v) spec (%v): %v", serviceID, specID, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}

// GET /{id}/specs/{specID}/doc
func (r *serviceResource) getSpecDoc(req *restful.Request, res *restful.Response) {
	var (
		serviceID = req.PathParameter("id")
		specID    = req.PathParameter("specID")
	)
	// Default download as true.
	download, err := strconv.ParseBool(req.QueryParameter("download"))
	if err != nil {
		download = true
	}
	shared.LogDebugf("get request to get service (%v) spec doc: %v", serviceID, specID)

	s, err := r.dao.Get(req.Request.Context(), serviceID)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	serviceID = s.ID
	spec, err := r.specDAO.Get(req.Request.Context(), specID, true)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	_, isJSON, err := spec.GetDocAsMap()
	if err != nil {
		shared.LogErrorf("failed to get spec doc as map for service (%v) spec (%v): %v", serviceID, specID, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	var (
		filename    = "spec-" + spec.Version + "-" + spec.Revision
		contentType string
	)
	if isJSON {
		filename += ".json"
		contentType = restful.MIME_JSON
	} else {
		filename += ".yaml"
		contentType = mimeTextPlain
	}
	res.Header().Set(restful.HEADER_ContentType, contentType)
	if download {
		res.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%v", filename))
	}

	_, _ = res.Write([]byte(*spec.Doc))
}

// POST /{id}/specs/{specID}/analyses
func (r *serviceResource) createAnalysis(req *restful.Request, res *restful.Response) {
	var (
		serviceID       = req.PathParameter("id")
		specID          = req.PathParameter("specID")
		specAnalysisReq = &models.SpecAnalysisRequest{}
	)
	shared.LogDebugf("get request to analyze service (%v) spec (%v)", serviceID, specID)

	if err := req.ReadEntity(specAnalysisReq); err != nil {
		shared.LogErrorf("failed to get specAnalysisReq from body: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	service, err := r.dao.Get(req.Request.Context(), serviceID)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	serviceID = service.ID
	spec, err := r.specDAO.Get(req.Request.Context(), specID, true)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	specAnalysisReq.Spec = spec
	specAnalysisReq.Service = service

	specAnalysisRes, err := r.runSpecAnalysisRequest(req.Request.Context(), res, specAnalysisReq, true, true)
	if err != nil {
		shared.LogErrorf("failed to analyze service (%v) spec (%v): %#v", serviceID, spec.ID, err)
		return
	}

	_ = res.WriteHeaderAndEntity(http.StatusOK, specAnalysisRes)
}

// GET /{id}/specs/{specID}/analyses
func (r *serviceResource) getAnalyses(req *restful.Request, res *restful.Response) {
	var (
		serviceID = req.PathParameter("id")
		specID    = req.PathParameter("specID")
	)
	shared.LogDebugf("get request to get service (%v) spec (%v) analyses", serviceID, specID)

	service, err := r.dao.Get(req.Request.Context(), serviceID)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	serviceID = service.ID
	specAnalyses, err := r.specAnalysisDAO.List(req.Request.Context(), &db.ListFilter{
		Model: &models.SpecAnalysis{},
		Indexes: map[string]string{
			"service_id": serviceID,
			"spec_id":    specID,
		},
		Sorters: []*db.Sorter{{
			Order: db.OrderDesc,
			Field: "created_at",
		}},
	})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	unfilteredCount := len(specAnalyses)
	specAnalyses = models.DistinctSpecAnalyses(specAnalyses)

	shared.LogDebugf("total %v specAnalysis(s) (unfiltered=%s) returned", len(specAnalyses), unfilteredCount)

	_ = res.WriteHeaderAndEntity(http.StatusOK, specAnalyses)
}

// GET /{id}/specs/analyses
func (r *serviceResource) getServiceAnalyses(req *restful.Request, res *restful.Response) {
	var (
		serviceID    = req.PathParameter("id")
		withFindings = parseQueryBool(req, "withFindings")
	)
	shared.LogDebugf("get request to get service (%v) analyses", serviceID)

	service, err := r.dao.Get(req.Request.Context(), serviceID)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	serviceID = service.ID
	specAnalyses, err := r.specAnalysisDAO.List(req.Request.Context(), &db.ListFilter{
		Model: &models.SpecAnalysis{},
		Indexes: map[string]string{
			"service_id": serviceID,
		},
		Sorters: []*db.Sorter{{
			Order: db.OrderDesc,
			Field: "created_at",
		}},
	})
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	specAnalyses = models.DistinctSpecAnalyses(specAnalyses)

	if !withFindings {
		for _, specAnalysis := range specAnalyses {
			specAnalysis.OmitResultFindings()
		}
	}

	_ = res.WriteHeaderAndEntity(http.StatusOK, specAnalyses)
}

// POST /{id}/specs/diff
func (r *serviceResource) createDiff(req *restful.Request, res *restful.Response) {
	var (
		serviceID   = req.PathParameter("id")
		specDiffReq = &models.SpecDiffRequest{}
	)
	if err := req.ReadEntity(specDiffReq); err != nil {
		shared.LogErrorf("failed to get specDiffReq from body: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}
	shared.LogDebugf("get request to diff service (%v) specs (old=%v, new=%v)", serviceID, specDiffReq.OldSpecID, specDiffReq.NewSpecID)

	s, err := r.dao.Get(req.Request.Context(), serviceID)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	serviceID = s.ID

	// Return the spec diff if it already exists.
	if specDiffList, err := r.specDiffDAO.List(req.Request.Context(), &db.ListFilter{
		Model: new(models.Spec),
		Indexes: map[string]string{
			"old_spec_id": specDiffReq.OldSpecID,
			"new_spec_id": specDiffReq.NewSpecID,
		},
	}); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	} else if len(specDiffList) > 0 {
		var matchingStoredSpecDiff *models.SpecDiff
		for _, storedSpecDiff := range specDiffList {
			if specDiffReq.Compare(storedSpecDiff.SpecDiffRequest) {
				matchingStoredSpecDiff = storedSpecDiff
				break
			}
		}
		if matchingStoredSpecDiff != nil {
			res.Header().Add("Location", "/v1/apiregistry/services/"+serviceID+"/specs/diff/"+matchingStoredSpecDiff.ID)
			_ = res.WriteHeaderAndEntity(http.StatusCreated, matchingStoredSpecDiff)
			return
		}
	}

	oldSpec, err := r.specDAO.Get(req.Request.Context(), specDiffReq.OldSpecID, true)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	specDiffReq.OldSpecDoc = oldSpec.Doc

	newSpec, err := r.specDAO.Get(req.Request.Context(), specDiffReq.NewSpecID, true)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	specDiffReq.NewSpecDoc = newSpec.Doc

	now := time.Now().UTC()
	specDiff := &models.SpecDiff{
		ID: shared.TimeUUID(),
		SpecDiffRequest: &models.SpecDiffRequest{
			NewSpecID: specDiffReq.NewSpecID,
			OldSpecID: specDiffReq.OldSpecID,
		},
		ServiceID: serviceID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	specDiff.Config = specDiffReq.Config

	if result, err := r.differSvc.Diff(specDiffReq); err != nil {
		shared.LogErrorf("failed to diff service (%v) specs (old=%v, new=%v): %#v", serviceID, oldSpec.ID, newSpec.ID, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	} else if err := specDiff.SetResult(result, "Diffed"); err != nil {
		shared.LogErrorf("failed to diff service (%v) specs (old=%v, new=%v): %#v", serviceID, oldSpec.ID, newSpec.ID, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.specDiffDAO.Save(req.Request.Context(), specDiff); err != nil {
		handleError(res, err)
		return
	}

	res.Header().Add("Location", "/v1/apiregistry/services/"+serviceID+"/specs/diff/"+specDiff.ID)
	_ = res.WriteHeaderAndEntity(http.StatusCreated, specDiff)
}

// GET /{id}/specs/diff/{oldSpecID}/{newSpecID}
func (r *serviceResource) getOrCreateDiff(req *restful.Request, res *restful.Response) {
	var (
		serviceID      = req.PathParameter("id")
		oldSpecID      = req.PathParameter("oldSpecID")
		newSpecID      = req.PathParameter("newSpecID")
		specDiffFormat = req.QueryParameter("format")
	)
	if specDiffFormat == "" {
		specDiffFormat = "json"
	}
	specDiffReq := &models.SpecDiffRequest{
		OldSpecID: oldSpecID,
		NewSpecID: newSpecID,
		SpecDiffConfig: models.SpecDiffConfig{
			Config: &diff.Config{
				OutputFormat: specDiffFormat,
			},
		},
	}

	shared.LogDebugf("get request to diff service (%v) specs (old=%v, new=%v)", serviceID, specDiffReq.OldSpecID, specDiffReq.NewSpecID)

	s, err := r.dao.Get(req.Request.Context(), serviceID)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	serviceID = s.ID
	var writeSpecDiffResult = func(res *restful.Response, specDiff *models.SpecDiff, format string) error {
		switch format {
		case "markdown":
			res.Header().Set(restful.HEADER_ContentType, mimeTextMarkdown)
			res.Header().Set("Content-Disposition", "attachment; filename=changelog.md")
			res.WriteHeader(http.StatusOK)
			_, _ = res.Write([]byte(specDiff.Result.Markdown))
		case "json":
			res.Header().Set(restful.HEADER_ContentType, restful.MIME_JSON)
			return res.WriteHeaderAndEntity(http.StatusOK, specDiff.Result.JSON)
		}
		return nil
	}

	// Return the spec diff if it already exists.
	if specDiffList, err := r.specDiffDAO.List(req.Request.Context(), &db.ListFilter{
		Model: new(models.Spec),
		Indexes: map[string]string{
			"old_spec_id": specDiffReq.OldSpecID,
			"new_spec_id": specDiffReq.NewSpecID,
		},
	}); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	} else if len(specDiffList) > 0 {
		var matchingStoredSpecDiff *models.SpecDiff
		for _, storedSpecDiff := range specDiffList {
			if specDiffReq.Compare(storedSpecDiff.SpecDiffRequest) {
				matchingStoredSpecDiff = storedSpecDiff
				break
			}
		}
		if matchingStoredSpecDiff != nil {
			res.Header().Add("Location", "/v1/apiregistry/services/"+serviceID+"/specs/diff/"+matchingStoredSpecDiff.ID)
			_ = writeSpecDiffResult(res, matchingStoredSpecDiff, specDiffFormat)
			return
		}
	}

	oldSpec, err := r.specDAO.Get(req.Request.Context(), specDiffReq.OldSpecID, true)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	specDiffReq.OldSpecDoc = oldSpec.Doc

	newSpec, err := r.specDAO.Get(req.Request.Context(), specDiffReq.NewSpecID, true)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}
	specDiffReq.NewSpecDoc = newSpec.Doc

	now := time.Now().UTC()
	specDiff := &models.SpecDiff{
		ID: shared.TimeUUID(),
		SpecDiffRequest: &models.SpecDiffRequest{
			NewSpecID: specDiffReq.NewSpecID,
			OldSpecID: specDiffReq.OldSpecID,
		},
		ServiceID: serviceID,
		CreatedAt: now,
		UpdatedAt: now,
	}
	specDiff.Config = specDiffReq.Config

	if result, err := r.differSvc.Diff(specDiffReq); err != nil {
		shared.LogErrorf("failed to diff service (%v) specs (old=%v, new=%v): %#v", serviceID, oldSpec.ID, newSpec.ID, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	} else if err := specDiff.SetResult(result, "Diffed"); err != nil {
		shared.LogErrorf("failed to diff service (%v) specs (old=%v, new=%v): %#v", serviceID, oldSpec.ID, newSpec.ID, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.specDiffDAO.Save(req.Request.Context(), specDiff); err != nil {
		handleError(res, err)
		return
	}

	res.Header().Add("Location", "/v1/apiregistry/services/"+serviceID+"/specs/diff/"+specDiff.ID)
	_ = writeSpecDiffResult(res, specDiff, specDiffFormat)
}

// POST /{id}/specs/reconstruct
func (r *serviceResource) reconstructSpec(req *restful.Request, res *restful.Response) {
	var (
		serviceID = req.PathParameter("id")
	)
	shared.LogDebugf("get request to reconstruct service (%v) spec", serviceID)

	s, err := r.dao.Get(req.Request.Context(), serviceID)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	serviceID = s.ID
	apiName := s.GetNameID(modelsanalyzer.Drift, nil)

	reconstructedSpecDoc, err := apiclarity.ReconstructSpec(req.Request.Context(), r.apiclarityClient, apiName)
	if err != nil {
		shared.LogErrorf("failed to reconstruct service (%v) spec: %#v", serviceID, err)
		if errors.Is(err, apiclarity.ErrNoAPITrafficFound) {
			_ = res.WriteHeaderAndEntity(http.StatusBadRequest, err.Error())
			return
		}
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now().UTC()
	spec := &models.Spec{
		ID:        shared.TimeUUID(),
		Doc:       reconstructedSpecDoc,
		Revision:  "1",
		ServiceID: serviceID,
		State:     "Reconstructed",
		CreatedAt: now,
		UpdatedAt: now,
	}

	if _, err := spec.LoadDocAsOAS(req.Request.Context(), false, true); err != nil {
		shared.LogErrorf("failed to load Spec.Doc as OAS from body: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if err := r.validate.Struct(spec); err != nil {
		shared.LogErrorf("failed to validate Spec %s - %v", spec.ID, err.Error())
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if err := r.specDAO.Save(req.Request.Context(), spec); err != nil {
		handleError(res, err)
		return
	}

	res.Header().Add("Location", "/v1/apiregistry/services/"+serviceID+"/specs/"+spec.ID)
	_ = res.WriteHeaderAndEntity(http.StatusCreated, spec)
}
