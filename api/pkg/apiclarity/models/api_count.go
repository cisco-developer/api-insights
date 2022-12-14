// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APICount Api count
//
// swagger:model ApiCount
type APICount struct {

	// api host name
	APIHostName string `json:"apiHostName,omitempty"`

	// hold the relevant api info id
	APIInfoID uint32 `json:"apiInfoId,omitempty"`

	// api port
	APIPort int64 `json:"apiPort,omitempty"`

	// api type
	APIType APIType `json:"apiType,omitempty"`

	// num calls
	NumCalls int64 `json:"numCalls,omitempty"`
}

// Validate validates this Api count
func (m *APICount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APICount) validateAPIType(formats strfmt.Registry) error {
	if swag.IsZero(m.APIType) { // not required
		return nil
	}

	if err := m.APIType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("apiType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("apiType")
		}
		return err
	}

	return nil
}

// ContextValidate validate this Api count based on the context it is used
func (m *APICount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPIType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APICount) contextValidateAPIType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.APIType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("apiType")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("apiType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APICount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APICount) UnmarshalBinary(b []byte) error {
	var res APICount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
