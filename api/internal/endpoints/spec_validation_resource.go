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
	"github.com/cisco-developer/api-insights/api/internal/models"
	"github.com/cisco-developer/api-insights/api/pkg/utils/shared"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"net/http"
)

type specValidationResource struct {
	config *shared.AppConfig
}

// Register the API
// prefix: /v1/apiregistry/specs/validations
func (r *specValidationResource) Register(config *shared.AppConfig, container *restful.Container, prefix string) {
	ws := &restful.WebService{}
	ws.Path(prefix).Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON).ApiVersion(config.AppVersion).Doc("APIs for SpecAnalysis.")

	var specValidationReq models.SpecValidationRequest
	var specValidationRes models.SpecValidationResult

	ws.Route(
		ws.POST("/validate").
			To(r.validate).
			Do(shared.RouteAuthHeader(ws)).
			Do(shared.RouteReturns(specValidationReq, http.StatusOK)).
			Do(shared.RouteReturns(&se, http.StatusBadRequest, http.StatusUnauthorized, http.StatusForbidden, http.StatusInternalServerError)).
			Do(shared.RouteReads(specValidationRes, "spec validation request"), shared.RouteWrites(specValidationRes)).
			Metadata(restfulspec.KeyOpenAPITags, []string{"stateless"}).
			Notes("Perform a stateless spec validation"))

	container.Add(ws)
}

func (r *specValidationResource) validate(req *restful.Request, res *restful.Response) {
	var (
		specValidationReq models.SpecValidationRequest
	)

	if err := req.ReadEntity(specValidationReq); err != nil {
		shared.LogErrorf("failed to get specValidationReq from body: %#v", err)
		_ = res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	vr, err := models.ValidateSpecDoc(req.Request.Context(), specValidationReq.Doc)
	if err != nil {
		shared.LogErrorf("failed to validate spec: %#v", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	_ = res.WriteHeaderAndEntity(http.StatusOK, vr)
}
