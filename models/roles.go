// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Roles roles
//
// swagger:model roles
type Roles struct {

	// role
	Role []*Role `json:"role"`
}

// Validate validates this roles
func (m *Roles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Roles) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	for i := 0; i < len(m.Role); i++ {
		if swag.IsZero(m.Role[i]) { // not required
			continue
		}

		if m.Role[i] != nil {
			if err := m.Role[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("role" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Roles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Roles) UnmarshalBinary(b []byte) error {
	var res Roles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
