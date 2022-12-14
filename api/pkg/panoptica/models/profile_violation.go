// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProfileViolation ProfileViolation
//
// swagger:model ProfileViolation
type ProfileViolation struct {

	// Condition Index
	//
	// Index of the condition violated in the API Policy Profile
	// Required: true
	ConditionIndex *int64 `json:"condition_index"`

	// Description
	//
	// Human readable description of the violation
	// Required: true
	Description *string `json:"description"`
}

// Validate validates this profile violation
func (m *ProfileViolation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConditionIndex(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProfileViolation) validateConditionIndex(formats strfmt.Registry) error {

	if err := validate.Required("condition_index", "body", m.ConditionIndex); err != nil {
		return err
	}

	return nil
}

func (m *ProfileViolation) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this profile violation based on context it is used
func (m *ProfileViolation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProfileViolation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProfileViolation) UnmarshalBinary(b []byte) error {
	var res ProfileViolation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
