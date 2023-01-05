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
	"github.com/cisco-developer/api-insights/api/pkg/differ"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type specDiffResource struct {
	config    *shared.AppConfig
	dao       db.SpecDiffDAO
	validate  *validator.Validate
	differSvc differ.Service
	specDAO   db.SpecDAO
}

// Register the API
// prefix: /v1/apiregistry/specs/diffs
func (r *specDiffResource) Register(config *shared.AppConfig, container *restful.Container, prefix string) {
	ws := &restful.WebService{}
	ws.Path(prefix).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON).ApiVersion(config.AppVersion).Doc("APIs for SpecDiff.")

	var specDiff models.SpecDiff
	var specDiffReq models.SpecDiffRequest
	var id = ws.PathParameter("id", "unique identifier for specDiff.").DataType("string")

	ws.Route(
		ws.POST("/diff").
			To(r.performDiff).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(specDiff, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteReads(specDiffReq, "spec diff request"), shared.RouteWrites(specDiff)).
			Do(shared.RouteParams(id)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"stateless"}).
			Notes("Perform a stateless spec diff").
			Consumes(restful.MIME_JSON, "multipart/form-data"))

	container.Add(ws)
}

func (r *specDiffResource) performDiff(req *restful.Request, res *restful.Response) {
	var (
		specDiffReq = &models.SpecDiffRequest{}
	)
	shared.LogDebugf("get request to diff specs")

	var specsGetter = func(oldSpecID, newSpecID string) (oldSpec *models.Spec, newSpec *models.Spec, err error) {
		oldSpec, err = r.specDAO.Get(req.Request.Context(), oldSpecID, true)
		if err == nil {
			newSpec, err = r.specDAO.Get(req.Request.Context(), newSpecID, true)
		}
		return oldSpec, newSpec, err
	}
	if err := specDiffReq.From(req, specsGetter); err != nil {
		shared.LogErrorf("failed to get specDiffReq from request: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	// validate old spec
	vrOld, err := models.ValidateSpecDoc(req.Request.Context(), specDiffReq.OldSpecDoc)
	if err != nil {
		shared.LogErrorf("failed to validate spec: %#v", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !vrOld.Valid {
		shared.LogErrorf("invalid spec: %#v", vrOld.Findings)
		_ = res.WriteHeaderAndEntity(http.StatusBadRequest, vrOld)
		return
	}

	// validate new spec
	vrNew, err := models.ValidateSpecDoc(req.Request.Context(), specDiffReq.NewSpecDoc)
	if err != nil {
		shared.LogErrorf("failed to validate spec: %#v", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !vrNew.Valid {
		shared.LogErrorf("invalid spec: %#v", vrNew.Findings)
		_ = res.WriteHeaderAndEntity(http.StatusBadRequest, vrNew)
		return
	}

	// diff old and new specs
	result, err := r.differSvc.Diff(specDiffReq)
	if err != nil {
		shared.LogErrorf("failed to diff specs: %#v", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	specDiff := &models.SpecDiff{
		SpecDiffResult: models.SpecDiffResult{
			Result: result,
		},
	}

	_ = res.WriteHeaderAndEntity(http.StatusOK, specDiff)
}
