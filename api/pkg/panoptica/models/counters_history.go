// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CountersHistory History Counters
//
// swagger:model CountersHistory
type CountersHistory struct {

	// Critical
	Critical []int64 `json:"critical"`

	// High
	High []int64 `json:"high"`

	// Low
	Low []int64 `json:"low"`

	// Medium
	Medium []int64 `json:"medium"`

	// Timestamp
	Timestamp []string `json:"timestamp"`

	// Unclassified
	Unclassified []int64 `json:"unclassified"`
}

// Validate validates this counters history
func (m *CountersHistory) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this counters history based on context it is used
func (m *CountersHistory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CountersHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CountersHistory) UnmarshalBinary(b []byte) error {
	var res CountersHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
