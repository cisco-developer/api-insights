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
	"github.com/cisco-developer/api-insights/api/internal/db"
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/pkg/analyzer"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type specAnalysisResource struct {
	config      *shared.AppConfig
	dao         db.SpecAnalysisDAO
	validate    *validator.Validate
	analyzerSvc analyzer.Service
}

// Register the API
// prefix: /v1/apiregistry/specs/analyses
func (r *specAnalysisResource) Register(config *shared.AppConfig, container *restful.Container, prefix string) {
	ws := &restful.WebService{}
	ws.Path(prefix).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON).ApiVersion(config.AppVersion).Doc("APIs for SpecAnalysis.")

	var specAnalysisReq models.SpecAnalysisRequest
	var specAnalysisRes models.SpecAnalysisResponse

	ws.Route(
		ws.POST("/analyze").
			To(r.analyze).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(specAnalysisRes, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteReads(specAnalysisReq, "spec analysis request"), shared.RouteWrites(specAnalysisRes)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"stateless"}).
			Notes("Create a new spec analysis"))

	container.Add(ws)
}

func (r *specAnalysisResource) analyze(req *restful.Request, res *restful.Response) {
	var (
		specAnalysisReq = &models.SpecAnalysisRequest{}
	)

	if err := req.ReadEntity(specAnalysisReq); err != nil {
		shared.LogErrorf("failed to get specAnalysisReq from body: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if !specAnalysisReq.HasSpec() {
		shared.LogErrorf("failed to get spec from body")
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	vr, err := models.ValidateSpecDoc(req.Request.Context(), specAnalysisReq.Spec.Doc)
	if err != nil {
		shared.LogErrorf("failed to validate service (%v) spec: %#v", specAnalysisReq.Service.NameID, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !vr.Valid {
		shared.LogErrorf("invalid service (%v) spec: %#v", specAnalysisReq.Service.NameID, vr.Findings)
		_ = res.WriteHeaderAndEntity(http.StatusBadRequest, vr)
		return
	}

	specAnalysisRes, err := r.analyzerSvc.Analyze(specAnalysisReq)
	if err != nil {
		shared.LogErrorf("failed to analyze service (%v) spec (%v): %#v", specAnalysisReq.Spec.ServiceID, specAnalysisReq.Spec.ID, err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = res.WriteHeaderAndEntity(http.StatusOK, specAnalysisRes)
}
