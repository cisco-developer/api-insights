// Code generated by go-swagger; DO NOT EDIT.

package api_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetAPISecurityExternalCatalogCatalogIDParams creates a new GetAPISecurityExternalCatalogCatalogIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPISecurityExternalCatalogCatalogIDParams() *GetAPISecurityExternalCatalogCatalogIDParams {
	return &GetAPISecurityExternalCatalogCatalogIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPISecurityExternalCatalogCatalogIDParamsWithTimeout creates a new GetAPISecurityExternalCatalogCatalogIDParams object
// with the ability to set a timeout on a request.
func NewGetAPISecurityExternalCatalogCatalogIDParamsWithTimeout(timeout time.Duration) *GetAPISecurityExternalCatalogCatalogIDParams {
	return &GetAPISecurityExternalCatalogCatalogIDParams{
		timeout: timeout,
	}
}

// NewGetAPISecurityExternalCatalogCatalogIDParamsWithContext creates a new GetAPISecurityExternalCatalogCatalogIDParams object
// with the ability to set a context for a request.
func NewGetAPISecurityExternalCatalogCatalogIDParamsWithContext(ctx context.Context) *GetAPISecurityExternalCatalogCatalogIDParams {
	return &GetAPISecurityExternalCatalogCatalogIDParams{
		Context: ctx,
	}
}

// NewGetAPISecurityExternalCatalogCatalogIDParamsWithHTTPClient creates a new GetAPISecurityExternalCatalogCatalogIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPISecurityExternalCatalogCatalogIDParamsWithHTTPClient(client *http.Client) *GetAPISecurityExternalCatalogCatalogIDParams {
	return &GetAPISecurityExternalCatalogCatalogIDParams{
		HTTPClient: client,
	}
}

/* GetAPISecurityExternalCatalogCatalogIDParams contains all the parameters to send to the API endpoint
   for the get API security external catalog catalog ID operation.

   Typically these are written to a http.Request.
*/
type GetAPISecurityExternalCatalogCatalogIDParams struct {

	/* APIPolicyProfiles.

	   Names of the Api Policy Profiles
	*/
	APIPolicyProfiles []string

	// CatalogID.
	//
	// Format: uuid
	CatalogID strfmt.UUID

	/* DownloadAsJSON.

	   When true, the API will return an json file, and pagination will be ignored
	*/
	DownloadAsJSON *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API security external catalog catalog ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPISecurityExternalCatalogCatalogIDParams) WithDefaults() *GetAPISecurityExternalCatalogCatalogIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API security external catalog catalog ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPISecurityExternalCatalogCatalogIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get API security external catalog catalog ID params
func (o *GetAPISecurityExternalCatalogCatalogIDParams) WithTimeout(timeout time.Duration) *GetAPISecurityExternalCatalogCatalogIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API security external catalog catalog ID params
func (o *GetAPISecurityExternalCatalogCatalogIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API security external catalog catalog ID params
func (o *GetAPISecurityExternalCatalogCatalogIDParams) WithContext(ctx context.Context) *GetAPISecurityExternalCatalogCatalogIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API security external catalog catalog ID params
func (o *GetAPISecurityExternalCatalogCatalogIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API security external catalog catalog ID params
func (o *GetAPISecurityExternalCatalogCatalogIDParams) WithHTTPClient(client *http.Client) *GetAPISecurityExternalCatalogCatalogIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API security external catalog catalog ID params
func (o *GetAPISecurityExternalCatalogCatalogIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIPolicyProfiles adds the aPIPolicyProfiles to the get API security external catalog catalog ID params
func (o *GetAPISecurityExternalCatalogCatalogIDParams) WithAPIPolicyProfiles(aPIPolicyProfiles []string) *GetAPISecurityExternalCatalogCatalogIDParams {
	o.SetAPIPolicyProfiles(aPIPolicyProfiles)
	return o
}

// SetAPIPolicyProfiles adds the apiPolicyProfiles to the get API security external catalog catalog ID params
func (o *GetAPISecurityExternalCatalogCatalogIDParams) SetAPIPolicyProfiles(aPIPolicyProfiles []string) {
	o.APIPolicyProfiles = aPIPolicyProfiles
}

// WithCatalogID adds the catalogID to the get API security external catalog catalog ID params
func (o *GetAPISecurityExternalCatalogCatalogIDParams) WithCatalogID(catalogID strfmt.UUID) *GetAPISecurityExternalCatalogCatalogIDParams {
	o.SetCatalogID(catalogID)
	return o
}

// SetCatalogID adds the catalogId to the get API security external catalog catalog ID params
func (o *GetAPISecurityExternalCatalogCatalogIDParams) SetCatalogID(catalogID strfmt.UUID) {
	o.CatalogID = catalogID
}

// WithDownloadAsJSON adds the downloadAsJSON to the get API security external catalog catalog ID params
func (o *GetAPISecurityExternalCatalogCatalogIDParams) WithDownloadAsJSON(downloadAsJSON *bool) *GetAPISecurityExternalCatalogCatalogIDParams {
	o.SetDownloadAsJSON(downloadAsJSON)
	return o
}

// SetDownloadAsJSON adds the downloadAsJson to the get API security external catalog catalog ID params
func (o *GetAPISecurityExternalCatalogCatalogIDParams) SetDownloadAsJSON(downloadAsJSON *bool) {
	o.DownloadAsJSON = downloadAsJSON
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPISecurityExternalCatalogCatalogIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIPolicyProfiles != nil {

		// binding items for apiPolicyProfiles
		joinedAPIPolicyProfiles := o.bindParamAPIPolicyProfiles(reg)

		// query array param apiPolicyProfiles
		if err := r.SetQueryParam("apiPolicyProfiles", joinedAPIPolicyProfiles...); err != nil {
			return err
		}
	}

	// path param catalogId
	if err := r.SetPathParam("catalogId", o.CatalogID.String()); err != nil {
		return err
	}

	if o.DownloadAsJSON != nil {

		// query param downloadAsJson
		var qrDownloadAsJSON bool

		if o.DownloadAsJSON != nil {
			qrDownloadAsJSON = *o.DownloadAsJSON
		}
		qDownloadAsJSON := swag.FormatBool(qrDownloadAsJSON)
		if qDownloadAsJSON != "" {

			if err := r.SetQueryParam("downloadAsJson", qDownloadAsJSON); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetAPISecurityExternalCatalogCatalogID binds the parameter apiPolicyProfiles
func (o *GetAPISecurityExternalCatalogCatalogIDParams) bindParamAPIPolicyProfiles(formats strfmt.Registry) []string {
	aPIPolicyProfilesIR := o.APIPolicyProfiles

	var aPIPolicyProfilesIC []string
	for _, aPIPolicyProfilesIIR := range aPIPolicyProfilesIR { // explode []string

		aPIPolicyProfilesIIV := aPIPolicyProfilesIIR // string as string
		aPIPolicyProfilesIC = append(aPIPolicyProfilesIC, aPIPolicyProfilesIIV)
	}

	// items.CollectionFormat: ""
	aPIPolicyProfilesIS := swag.JoinByFormat(aPIPolicyProfilesIC, "")

	return aPIPolicyProfilesIS
}