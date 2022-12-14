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

// CategoryScoreGrade CategoryScoreGrade
//
// swagger:model CategoryScoreGrade
type CategoryScoreGrade struct {

	// Additional Info
	AdditionalInfo []*AdditionalInfo `json:"additional_info"`

	// confidence
	Confidence RiskConfidenceEnum `json:"confidence,omitempty"`

	// counters history
	CountersHistory *CountersHistory `json:"counters_history,omitempty"`

	// critical
	// Required: true
	Critical *ScoreFindingGroup `json:"critical"`

	// high
	// Required: true
	High *ScoreFindingGroup `json:"high"`

	// low
	// Required: true
	Low *ScoreFindingGroup `json:"low"`

	// medium
	// Required: true
	Medium *ScoreFindingGroup `json:"medium"`

	// Name
	// Required: true
	Name *string `json:"name"`

	// risk
	// Required: true
	Risk *APISecurityRiskSeverity `json:"risk"`

	// Scorer Version
	// Required: true
	ScorerVersion *int64 `json:"scorer_version"`

	// trend
	Trend RiskTrendEnum `json:"trend,omitempty"`

	// unclassified
	// Required: true
	Unclassified *ScoreFindingGroup `json:"unclassified"`
}

// Validate validates this category score grade
func (m *CategoryScoreGrade) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdditionalInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfidence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCountersHistory(formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScorerVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrend(formats); err != nil {
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

func (m *CategoryScoreGrade) validateAdditionalInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.AdditionalInfo) { // not required
		return nil
	}

	for i := 0; i < len(m.AdditionalInfo); i++ {
		if swag.IsZero(m.AdditionalInfo[i]) { // not required
			continue
		}

		if m.AdditionalInfo[i] != nil {
			if err := m.AdditionalInfo[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additional_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additional_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CategoryScoreGrade) validateConfidence(formats strfmt.Registry) error {
	if swag.IsZero(m.Confidence) { // not required
		return nil
	}

	if err := m.Confidence.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("confidence")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("confidence")
		}
		return err
	}

	return nil
}

func (m *CategoryScoreGrade) validateCountersHistory(formats strfmt.Registry) error {
	if swag.IsZero(m.CountersHistory) { // not required
		return nil
	}

	if m.CountersHistory != nil {
		if err := m.CountersHistory.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counters_history")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counters_history")
			}
			return err
		}
	}

	return nil
}

func (m *CategoryScoreGrade) validateCritical(formats strfmt.Registry) error {

	if err := validate.Required("critical", "body", m.Critical); err != nil {
		return err
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

func (m *CategoryScoreGrade) validateHigh(formats strfmt.Registry) error {

	if err := validate.Required("high", "body", m.High); err != nil {
		return err
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

func (m *CategoryScoreGrade) validateLow(formats strfmt.Registry) error {

	if err := validate.Required("low", "body", m.Low); err != nil {
		return err
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

func (m *CategoryScoreGrade) validateMedium(formats strfmt.Registry) error {

	if err := validate.Required("medium", "body", m.Medium); err != nil {
		return err
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

func (m *CategoryScoreGrade) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CategoryScoreGrade) validateRisk(formats strfmt.Registry) error {

	if err := validate.Required("risk", "body", m.Risk); err != nil {
		return err
	}

	if err := validate.Required("risk", "body", m.Risk); err != nil {
		return err
	}

	if m.Risk != nil {
		if err := m.Risk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("risk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("risk")
			}
			return err
		}
	}

	return nil
}

func (m *CategoryScoreGrade) validateScorerVersion(formats strfmt.Registry) error {

	if err := validate.Required("scorer_version", "body", m.ScorerVersion); err != nil {
		return err
	}

	return nil
}

func (m *CategoryScoreGrade) validateTrend(formats strfmt.Registry) error {
	if swag.IsZero(m.Trend) { // not required
		return nil
	}

	if err := m.Trend.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trend")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("trend")
		}
		return err
	}

	return nil
}

func (m *CategoryScoreGrade) validateUnclassified(formats strfmt.Registry) error {

	if err := validate.Required("unclassified", "body", m.Unclassified); err != nil {
		return err
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

// ContextValidate validate this category score grade based on the context it is used
func (m *CategoryScoreGrade) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdditionalInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConfidence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountersHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.contextValidateRisk(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrend(ctx, formats); err != nil {
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

func (m *CategoryScoreGrade) contextValidateAdditionalInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AdditionalInfo); i++ {

		if m.AdditionalInfo[i] != nil {
			if err := m.AdditionalInfo[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("additional_info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("additional_info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CategoryScoreGrade) contextValidateConfidence(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Confidence.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("confidence")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("confidence")
		}
		return err
	}

	return nil
}

func (m *CategoryScoreGrade) contextValidateCountersHistory(ctx context.Context, formats strfmt.Registry) error {

	if m.CountersHistory != nil {
		if err := m.CountersHistory.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counters_history")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counters_history")
			}
			return err
		}
	}

	return nil
}

func (m *CategoryScoreGrade) contextValidateCritical(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CategoryScoreGrade) contextValidateHigh(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CategoryScoreGrade) contextValidateLow(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CategoryScoreGrade) contextValidateMedium(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CategoryScoreGrade) contextValidateRisk(ctx context.Context, formats strfmt.Registry) error {

	if m.Risk != nil {
		if err := m.Risk.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("risk")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("risk")
			}
			return err
		}
	}

	return nil
}

func (m *CategoryScoreGrade) contextValidateTrend(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Trend.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("trend")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("trend")
		}
		return err
	}

	return nil
}

func (m *CategoryScoreGrade) contextValidateUnclassified(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CategoryScoreGrade) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CategoryScoreGrade) UnmarshalBinary(b []byte) error {
	var res CategoryScoreGrade
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
