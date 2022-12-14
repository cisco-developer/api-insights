// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// APIServiceComplianceSimple ApiServiceCompliance_Simple
//
// swagger:model ApiServiceCompliance_Simple
type APIServiceComplianceSimple struct {

	// Compliant
	// Required: true
	Compliant *bool `json:"compliant"`

	// Profilescompliance
	// Required: true
	Profilescompliance []*APIServiceProfileComplianceSimple `json:"profilescompliance"`
}

// Validate validates this Api service compliance simple
func (m *APIServiceComplianceSimple) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompliant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfilescompliance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIServiceComplianceSimple) validateCompliant(formats strfmt.Registry) error {

	if err := validate.Required("compliant", "body", m.Compliant); err != nil {
		return err
	}

	return nil
}

func (m *APIServiceComplianceSimple) validateProfilescompliance(formats strfmt.Registry) error {

	if err := validate.Required("profilescompliance", "body", m.Profilescompliance); err != nil {
		return err
	}

	for i := 0; i < len(m.Profilescompliance); i++ {
		if swag.IsZero(m.Profilescompliance[i]) { // not required
			continue
		}

		if m.Profilescompliance[i] != nil {
			if err := m.Profilescompliance[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("profilescompliance" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("profilescompliance" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this Api service compliance simple based on the context it is used
func (m *APIServiceComplianceSimple) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProfilescompliance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIServiceComplianceSimple) contextValidateProfilescompliance(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Profilescompliance); i++ {

		if m.Profilescompliance[i] != nil {
			if err := m.Profilescompliance[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("profilescompliance" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("profilescompliance" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIServiceComplianceSimple) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIServiceComplianceSimple) UnmarshalBinary(b []byte) error {
	var res APIServiceComplianceSimple
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
