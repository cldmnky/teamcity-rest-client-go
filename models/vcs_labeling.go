// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VcsLabeling vcs labeling
//
// swagger:model vcsLabeling
type VcsLabeling struct {

	// branch filter
	BranchFilter string `json:"branchFilter,omitempty"`

	// label name
	LabelName string `json:"labelName,omitempty" xml:"labelName"`

	// type
	Type string `json:"type,omitempty" xml:"type"`

	// vcs roots
	VcsRoots *VcsRoots `json:"vcsRoots,omitempty"`
}

// Validate validates this vcs labeling
func (m *VcsLabeling) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVcsRoots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VcsLabeling) validateVcsRoots(formats strfmt.Registry) error {

	if swag.IsZero(m.VcsRoots) { // not required
		return nil
	}

	if m.VcsRoots != nil {
		if err := m.VcsRoots.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vcsRoots")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VcsLabeling) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcsLabeling) UnmarshalBinary(b []byte) error {
	var res VcsLabeling
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}