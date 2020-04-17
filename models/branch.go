// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Branch branch
//
// swagger:model branch
type Branch struct {

	// active
	Active *bool `json:"active,omitempty" xml:"active"`

	// builds
	Builds *Builds `json:"builds,omitempty"`

	// default
	Default *bool `json:"default,omitempty" xml:"default"`

	// group flag
	GroupFlag *bool `json:"groupFlag,omitempty" xml:"groupFlag"`

	// internal name
	InternalName string `json:"internalName,omitempty" xml:"internalName"`

	// last activity
	LastActivity string `json:"lastActivity,omitempty" xml:"lastActivity"`

	// name
	Name string `json:"name,omitempty" xml:"name"`

	// unspecified
	Unspecified *bool `json:"unspecified,omitempty" xml:"unspecified"`
}

// Validate validates this branch
func (m *Branch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuilds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Branch) validateBuilds(formats strfmt.Registry) error {

	if swag.IsZero(m.Builds) { // not required
		return nil
	}

	if m.Builds != nil {
		if err := m.Builds.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("builds")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Branch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Branch) UnmarshalBinary(b []byte) error {
	var res Branch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
