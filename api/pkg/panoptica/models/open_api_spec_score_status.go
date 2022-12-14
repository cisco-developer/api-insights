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

// OpenAPISpecScoreStatus OpenApiSpecScoreStatus
//
// swagger:model OpenApiSpecScoreStatus
type OpenAPISpecScoreStatus string

func NewOpenAPISpecScoreStatus(value OpenAPISpecScoreStatus) *OpenAPISpecScoreStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated OpenAPISpecScoreStatus.
func (m OpenAPISpecScoreStatus) Pointer() *OpenAPISpecScoreStatus {
	return &m
}

const (

	// OpenAPISpecScoreStatusSCORED captures enum value "SCORED"
	OpenAPISpecScoreStatusSCORED OpenAPISpecScoreStatus = "SCORED"

	// OpenAPISpecScoreStatusNOTSCORED captures enum value "NOT_SCORED"
	OpenAPISpecScoreStatusNOTSCORED OpenAPISpecScoreStatus = "NOT_SCORED"
)

// for schema
var openApiSpecScoreStatusEnum []interface{}

func init() {
	var res []OpenAPISpecScoreStatus
	if err := json.Unmarshal([]byte(`["SCORED","NOT_SCORED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openApiSpecScoreStatusEnum = append(openApiSpecScoreStatusEnum, v)
	}
}

func (m OpenAPISpecScoreStatus) validateOpenAPISpecScoreStatusEnum(path, location string, value OpenAPISpecScoreStatus) error {
	if err := validate.EnumCase(path, location, value, openApiSpecScoreStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this open Api spec score status
func (m OpenAPISpecScoreStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOpenAPISpecScoreStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this open Api spec score status based on context it is used
func (m OpenAPISpecScoreStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
