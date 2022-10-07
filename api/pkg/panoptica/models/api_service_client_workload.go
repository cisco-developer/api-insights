// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIServiceClientWorkload Api service client workload
//
// swagger:model ApiServiceClientWorkload
type APIServiceClientWorkload struct {

	// cluster
	Cluster string `json:"cluster,omitempty"`

	// namespace
	Namespace string `json:"namespace,omitempty"`

	// workload name
	WorkloadName string `json:"workloadName,omitempty"`
}

// Validate validates this Api service client workload
func (m *APIServiceClientWorkload) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Api service client workload based on context it is used
func (m *APIServiceClientWorkload) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIServiceClientWorkload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIServiceClientWorkload) UnmarshalBinary(b []byte) error {
	var res APIServiceClientWorkload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}