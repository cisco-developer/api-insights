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

// SpecScoreFindings spec score findings
//
// swagger:model SpecScoreFindings
type SpecScoreFindings struct {

	// critical
	Critical *SpecScoreFindingsList `json:"critical,omitempty"`

	// high
	High *SpecScoreFindingsList `json:"high,omitempty"`

	// low
	Low *SpecScoreFindingsList `json:"low,omitempty"`

	// medium
	Medium *SpecScoreFindingsList `json:"medium,omitempty"`

	// unclassified
	Unclassified *SpecScoreFindingsList `json:"unclassified,omitempty"`
}

// Validate validates this spec score findings
func (m *SpecScoreFindings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCritical(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHigh(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMedium(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnclassified(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpecScoreFindings) validateCritical(formats strfmt.Registry) error {
	if swag.IsZero(m.Critical) { // not required
		return nil
	}

	if m.Critical != nil {
		if err := m.Critical.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("critical")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("critical")
			}
			return err
		}
	}

	return nil
}

func (m *SpecScoreFindings) validateHigh(formats strfmt.Registry) error {
	if swag.IsZero(m.High) { // not required
		return nil
	}

	if m.High != nil {
		if err := m.High.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("high")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("high")
			}
			return err
		}
	}

	return nil
}

func (m *SpecScoreFindings) validateLow(formats strfmt.Registry) error {
	if swag.IsZero(m.Low) { // not required
		return nil
	}

	if m.Low != nil {
		if err := m.Low.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("low")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("low")
			}
			return err
		}
	}

	return nil
}

func (m *SpecScoreFindings) validateMedium(formats strfmt.Registry) error {
	if swag.IsZero(m.Medium) { // not required
		return nil
	}

	if m.Medium != nil {
		if err := m.Medium.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("medium")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("medium")
			}
			return err
		}
	}

	return nil
}

func (m *SpecScoreFindings) validateUnclassified(formats strfmt.Registry) error {
	if swag.IsZero(m.Unclassified) { // not required
		return nil
	}

	if m.Unclassified != nil {
		if err := m.Unclassified.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unclassified")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unclassified")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this spec score findings based on the context it is used
func (m *SpecScoreFindings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCritical(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHigh(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMedium(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUnclassified(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SpecScoreFindings) contextValidateCritical(ctx context.Context, formats strfmt.Registry) error {

	if m.Critical != nil {
		if err := m.Critical.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("critical")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("critical")
			}
			return err
		}
	}

	return nil
}

func (m *SpecScoreFindings) contextValidateHigh(ctx context.Context, formats strfmt.Registry) error {

	if m.High != nil {
		if err := m.High.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("high")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("high")
			}
			return err
		}
	}

	return nil
}

func (m *SpecScoreFindings) contextValidateLow(ctx context.Context, formats strfmt.Registry) error {

	if m.Low != nil {
		if err := m.Low.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("low")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("low")
			}
			return err
		}
	}

	return nil
}

func (m *SpecScoreFindings) contextValidateMedium(ctx context.Context, formats strfmt.Registry) error {

	if m.Medium != nil {
		if err := m.Medium.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("medium")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("medium")
			}
			return err
		}
	}

	return nil
}

func (m *SpecScoreFindings) contextValidateUnclassified(ctx context.Context, formats strfmt.Registry) error {

	if m.Unclassified != nil {
		if err := m.Unclassified.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("unclassified")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("unclassified")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SpecScoreFindings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SpecScoreFindings) UnmarshalBinary(b []byte) error {
	var res SpecScoreFindings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
