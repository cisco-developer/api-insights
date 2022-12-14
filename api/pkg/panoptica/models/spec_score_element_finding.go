// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SpecScoreElementFinding spec score element finding
//
// swagger:model SpecScoreElementFinding
type SpecScoreElementFinding struct {

	// description
	Description string `json:"description,omitempty"`

	// mitigation
	Mitigation string `json:"mitigation,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// spec path
	SpecPath string `json:"specPath,omitempty"`
}

// Validate validates this spec score element finding
func (m *SpecScoreElementFinding) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this spec score element finding based on context it is used
func (m *SpecScoreElementFinding) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SpecScoreElementFinding) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpecScoreElementFinding) UnmarshalBinary(b []byte) error {
	var res SpecScoreElementFinding
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
