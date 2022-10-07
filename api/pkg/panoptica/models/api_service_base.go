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

// APIServiceBase ApiService_Base
//
// swagger:model ApiService_Base
type APIServiceBase struct {

	// Classification
	//
	// API classification label as determined by Crankshaft, e.g. ['meetings', 'messaging']
	Classification []string `json:"classification"`

	// Timestamp
	// Format: date-time
	CreationTimestamp strfmt.DateTime `json:"creation_timestamp,omitempty"`

	// Description
	//
	// Textual description of the Service
	Description string `json:"description,omitempty"`

	// id
	//
	// Unique id of the subject API as assigned by Crankshaft
	// Required: true
	// Format: uuid
	Identifier *strfmt.UUID `json:"identifier"`

	// Name
	//
	// API name, usually an FQDN as determined by crankshaft, it can be logical or can correspond to one of the endpoints where the API is reachable, i.e. api.webex.com
	// Required: true
	Name *string `json:"name"`

	// Provider Id
	//
	// API provider id
	// Format: uuid
	ProviderID strfmt.UUID `json:"provider_id,omitempty"`

	// risk
	Risk APISecurityRiskSeverity `json:"risk,omitempty"`

	// status
	Status APISecurityAPIStatus `json:"status,omitempty"`

	// status description
	StatusDescription string `json:"status_description,omitempty"`
}

// Validate validates this Api service base
func (m *APIServiceBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProviderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIServiceBase) validateCreationTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("creation_timestamp", "body", "date-time", m.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIServiceBase) validateIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("identifier", "body", m.Identifier); err != nil {
		return err
	}

	if err := validate.FormatOf("identifier", "body", "uuid", m.Identifier.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIServiceBase) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *APIServiceBase) validateProviderID(formats strfmt.Registry) error {
	if swag.IsZero(m.ProviderID) { // not required
		return nil
	}

	if err := validate.FormatOf("provider_id", "body", "uuid", m.ProviderID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *APIServiceBase) validateRisk(formats strfmt.Registry) error {
	if swag.IsZero(m.Risk) { // not required
		return nil
	}

	if err := m.Risk.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("risk")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("risk")
		}
		return err
	}

	return nil
}

func (m *APIServiceBase) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// ContextValidate validate this Api service base based on the context it is used
func (m *APIServiceBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIServiceBase) contextValidateRisk(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Risk.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("risk")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("risk")
		}
		return err
	}

	return nil
}

func (m *APIServiceBase) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIServiceBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIServiceBase) UnmarshalBinary(b []byte) error {
	var res APIServiceBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}