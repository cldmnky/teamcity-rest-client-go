// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProgressInfo progress info
//
// swagger:model progress-info
type ProgressInfo struct {

	// current stage text
	CurrentStageText string `json:"currentStageText,omitempty" xml:"currentStageText"`

	// elapsed seconds
	ElapsedSeconds int64 `json:"elapsedSeconds,omitempty" xml:"elapsedSeconds"`

	// estimated total seconds
	EstimatedTotalSeconds int64 `json:"estimatedTotalSeconds,omitempty" xml:"estimatedTotalSeconds"`

	// last activity time
	LastActivityTime string `json:"lastActivityTime,omitempty" xml:"lastActivityTime"`

	// left seconds
	LeftSeconds int64 `json:"leftSeconds,omitempty" xml:"leftSeconds"`

	// outdated
	Outdated *bool `json:"outdated,omitempty" xml:"outdated"`

	// outdated reason build
	OutdatedReasonBuild *Build `json:"outdatedReasonBuild,omitempty"`

	// percentage complete
	PercentageComplete int32 `json:"percentageComplete,omitempty" xml:"percentageComplete"`

	// probably hanging
	ProbablyHanging *bool `json:"probablyHanging,omitempty" xml:"probablyHanging"`
}

// Validate validates this progress info
func (m *ProgressInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOutdatedReasonBuild(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProgressInfo) validateOutdatedReasonBuild(formats strfmt.Registry) error {

	if swag.IsZero(m.OutdatedReasonBuild) { // not required
		return nil
	}

	if m.OutdatedReasonBuild != nil {
		if err := m.OutdatedReasonBuild.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outdatedReasonBuild")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProgressInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProgressInfo) UnmarshalBinary(b []byte) error {
	var res ProgressInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
