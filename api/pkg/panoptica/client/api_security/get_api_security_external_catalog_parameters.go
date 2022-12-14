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

// NewGetAPISecurityExternalCatalogParams creates a new GetAPISecurityExternalCatalogParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAPISecurityExternalCatalogParams() *GetAPISecurityExternalCatalogParams {
	return &GetAPISecurityExternalCatalogParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPISecurityExternalCatalogParamsWithTimeout creates a new GetAPISecurityExternalCatalogParams object
// with the ability to set a timeout on a request.
func NewGetAPISecurityExternalCatalogParamsWithTimeout(timeout time.Duration) *GetAPISecurityExternalCatalogParams {
	return &GetAPISecurityExternalCatalogParams{
		timeout: timeout,
	}
}

// NewGetAPISecurityExternalCatalogParamsWithContext creates a new GetAPISecurityExternalCatalogParams object
// with the ability to set a context for a request.
func NewGetAPISecurityExternalCatalogParamsWithContext(ctx context.Context) *GetAPISecurityExternalCatalogParams {
	return &GetAPISecurityExternalCatalogParams{
		Context: ctx,
	}
}

// NewGetAPISecurityExternalCatalogParamsWithHTTPClient creates a new GetAPISecurityExternalCatalogParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAPISecurityExternalCatalogParamsWithHTTPClient(client *http.Client) *GetAPISecurityExternalCatalogParams {
	return &GetAPISecurityExternalCatalogParams{
		HTTPClient: client,
	}
}

/* GetAPISecurityExternalCatalogParams contains all the parameters to send to the API endpoint
   for the get API security external catalog operation.

   Typically these are written to a http.Request.
*/
type GetAPISecurityExternalCatalogParams struct {

	/* APIPolicyProfiles.

	   Names of the Api Policy Profiles
	*/
	APIPolicyProfiles []string

	/* IncludeServiceWithNoSpec.

	   When false, only services with specs wikk be returned

	   Default: true
	*/
	IncludeServiceWithNoSpec *bool

	/* MaxResults.

	   The number of entries to return (pagination)

	   Default: 100
	*/
	MaxResults *float64

	/* Name.

	   the Api Catalog name filter
	*/
	Name *string

	/* NoPagination.

	   When true, the pagination params will be ignored
	*/
	NoPagination *bool

	/* Offset.

	   Return entries from this offset (pagination)
	*/
	Offset *float64

	/* SortDir.

	   sorting direction

	   Default: "ASC"
	*/
	SortDir *string

	/* SortKey.

	   the Api Catalog sort key
	*/
	SortKey *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get API security external catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPISecurityExternalCatalogParams) WithDefaults() *GetAPISecurityExternalCatalogParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get API security external catalog params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAPISecurityExternalCatalogParams) SetDefaults() {
	var (
		includeServiceWithNoSpecDefault = bool(true)

		maxResultsDefault = float64(100)

		offsetDefault = float64(0)

		sortDirDefault = string("ASC")
	)

	val := GetAPISecurityExternalCatalogParams{
		IncludeServiceWithNoSpec: &includeServiceWithNoSpecDefault,
		MaxResults:               &maxResultsDefault,
		Offset:                   &offsetDefault,
		SortDir:                  &sortDirDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) WithTimeout(timeout time.Duration) *GetAPISecurityExternalCatalogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) WithContext(ctx context.Context) *GetAPISecurityExternalCatalogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) WithHTTPClient(client *http.Client) *GetAPISecurityExternalCatalogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIPolicyProfiles adds the aPIPolicyProfiles to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) WithAPIPolicyProfiles(aPIPolicyProfiles []string) *GetAPISecurityExternalCatalogParams {
	o.SetAPIPolicyProfiles(aPIPolicyProfiles)
	return o
}

// SetAPIPolicyProfiles adds the apiPolicyProfiles to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) SetAPIPolicyProfiles(aPIPolicyProfiles []string) {
	o.APIPolicyProfiles = aPIPolicyProfiles
}

// WithIncludeServiceWithNoSpec adds the includeServiceWithNoSpec to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) WithIncludeServiceWithNoSpec(includeServiceWithNoSpec *bool) *GetAPISecurityExternalCatalogParams {
	o.SetIncludeServiceWithNoSpec(includeServiceWithNoSpec)
	return o
}

// SetIncludeServiceWithNoSpec adds the includeServiceWithNoSpec to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) SetIncludeServiceWithNoSpec(includeServiceWithNoSpec *bool) {
	o.IncludeServiceWithNoSpec = includeServiceWithNoSpec
}

// WithMaxResults adds the maxResults to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) WithMaxResults(maxResults *float64) *GetAPISecurityExternalCatalogParams {
	o.SetMaxResults(maxResults)
	return o
}

// SetMaxResults adds the maxResults to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) SetMaxResults(maxResults *float64) {
	o.MaxResults = maxResults
}

// WithName adds the name to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) WithName(name *string) *GetAPISecurityExternalCatalogParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) SetName(name *string) {
	o.Name = name
}

// WithNoPagination adds the noPagination to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) WithNoPagination(noPagination *bool) *GetAPISecurityExternalCatalogParams {
	o.SetNoPagination(noPagination)
	return o
}

// SetNoPagination adds the noPagination to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) SetNoPagination(noPagination *bool) {
	o.NoPagination = noPagination
}

// WithOffset adds the offset to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) WithOffset(offset *float64) *GetAPISecurityExternalCatalogParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) SetOffset(offset *float64) {
	o.Offset = offset
}

// WithSortDir adds the sortDir to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) WithSortDir(sortDir *string) *GetAPISecurityExternalCatalogParams {
	o.SetSortDir(sortDir)
	return o
}

// SetSortDir adds the sortDir to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) SetSortDir(sortDir *string) {
	o.SortDir = sortDir
}

// WithSortKey adds the sortKey to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) WithSortKey(sortKey *string) *GetAPISecurityExternalCatalogParams {
	o.SetSortKey(sortKey)
	return o
}

// SetSortKey adds the sortKey to the get API security external catalog params
func (o *GetAPISecurityExternalCatalogParams) SetSortKey(sortKey *string) {
	o.SortKey = sortKey
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPISecurityExternalCatalogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IncludeServiceWithNoSpec != nil {

		// query param includeServiceWithNoSpec
		var qrIncludeServiceWithNoSpec bool

		if o.IncludeServiceWithNoSpec != nil {
			qrIncludeServiceWithNoSpec = *o.IncludeServiceWithNoSpec
		}
		qIncludeServiceWithNoSpec := swag.FormatBool(qrIncludeServiceWithNoSpec)
		if qIncludeServiceWithNoSpec != "" {

			if err := r.SetQueryParam("includeServiceWithNoSpec", qIncludeServiceWithNoSpec); err != nil {
				return err
			}
		}
	}

	if o.MaxResults != nil {

		// query param maxResults
		var qrMaxResults float64

		if o.MaxResults != nil {
			qrMaxResults = *o.MaxResults
		}
		qMaxResults := swag.FormatFloat64(qrMaxResults)
		if qMaxResults != "" {

			if err := r.SetQueryParam("maxResults", qMaxResults); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NoPagination != nil {

		// query param noPagination
		var qrNoPagination bool

		if o.NoPagination != nil {
			qrNoPagination = *o.NoPagination
		}
		qNoPagination := swag.FormatBool(qrNoPagination)
		if qNoPagination != "" {

			if err := r.SetQueryParam("noPagination", qNoPagination); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset float64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatFloat64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.SortDir != nil {

		// query param sortDir
		var qrSortDir string

		if o.SortDir != nil {
			qrSortDir = *o.SortDir
		}
		qSortDir := qrSortDir
		if qSortDir != "" {

			if err := r.SetQueryParam("sortDir", qSortDir); err != nil {
				return err
			}
		}
	}

	if o.SortKey != nil {

		// query param sortKey
		var qrSortKey string

		if o.SortKey != nil {
			qrSortKey = *o.SortKey
		}
		qSortKey := qrSortKey
		if qSortKey != "" {

			if err := r.SetQueryParam("sortKey", qSortKey); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetAPISecurityExternalCatalog binds the parameter apiPolicyProfiles
func (o *GetAPISecurityExternalCatalogParams) bindParamAPIPolicyProfiles(formats strfmt.Registry) []string {
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
