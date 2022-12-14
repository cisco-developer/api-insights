// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// HTTPMethod Http method
//
// swagger:model HttpMethod
type HTTPMethod string

func NewHTTPMethod(value HTTPMethod) *HTTPMethod {
	return &value
}

// Pointer returns a pointer to a freshly-allocated HTTPMethod.
func (m HTTPMethod) Pointer() *HTTPMethod {
	return &m
}

const (

	// HTTPMethodGET captures enum value "GET"
	HTTPMethodGET HTTPMethod = "GET"

	// HTTPMethodPOST captures enum value "POST"
	HTTPMethodPOST HTTPMethod = "POST"

	// HTTPMethodPUT captures enum value "PUT"
	HTTPMethodPUT HTTPMethod = "PUT"

	// HTTPMethodDELETE captures enum value "DELETE"
	HTTPMethodDELETE HTTPMethod = "DELETE"

	// HTTPMethodHEAD captures enum value "HEAD"
	HTTPMethodHEAD HTTPMethod = "HEAD"

	// HTTPMethodCONNECT captures enum value "CONNECT"
	HTTPMethodCONNECT HTTPMethod = "CONNECT"

	// HTTPMethodOPTIONS captures enum value "OPTIONS"
	HTTPMethodOPTIONS HTTPMethod = "OPTIONS"

	// HTTPMethodTRACE captures enum value "TRACE"
	HTTPMethodTRACE HTTPMethod = "TRACE"

	// HTTPMethodPATCH captures enum value "PATCH"
	HTTPMethodPATCH HTTPMethod = "PATCH"
)

// for schema
var httpMethodEnum []interface{}

func init() {
	var res []HTTPMethod
	if err := json.Unmarshal([]byte(`["GET","POST","PUT","DELETE","HEAD","CONNECT","OPTIONS","TRACE","PATCH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		httpMethodEnum = append(httpMethodEnum, v)
	}
}

func (m HTTPMethod) validateHTTPMethodEnum(path, location string, value HTTPMethod) error {
	if err := validate.EnumCase(path, location, value, httpMethodEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this Http method
func (m HTTPMethod) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateHTTPMethodEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this Http method based on context it is used
func (m HTTPMethod) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
