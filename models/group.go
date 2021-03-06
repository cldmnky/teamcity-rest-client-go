// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Group group
//
// swagger:model group
type Group struct {

	// child groups
	ChildGroups *Groups `json:"child-groups,omitempty"`

	// description
	Description string `json:"description,omitempty" xml:"description"`

	// href
	Href string `json:"href,omitempty" xml:"href"`

	// key
	Key string `json:"key,omitempty" xml:"key"`

	// name
	Name string `json:"name,omitempty" xml:"name"`

	// parent groups
	ParentGroups *Groups `json:"parent-groups,omitempty"`

	// properties
	Properties *Properties `json:"properties,omitempty"`

	// roles
	Roles *Roles `json:"roles,omitempty"`

	// users
	Users *Users `json:"users,omitempty"`
}

// Validate validates this group
func (m *Group) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChildGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Group) validateChildGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.ChildGroups) { // not required
		return nil
	}

	if m.ChildGroups != nil {
		if err := m.ChildGroups.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("child-groups")
			}
			return err
		}
	}

	return nil
}

func (m *Group) validateParentGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.ParentGroups) { // not required
		return nil
	}

	if m.ParentGroups != nil {
		if err := m.ParentGroups.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent-groups")
			}
			return err
		}
	}

	return nil
}

func (m *Group) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if m.Properties != nil {
		if err := m.Properties.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("properties")
			}
			return err
		}
	}

	return nil
}

func (m *Group) validateRoles(formats strfmt.Registry) error {

	if swag.IsZero(m.Roles) { // not required
		return nil
	}

	if m.Roles != nil {
		if err := m.Roles.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("roles")
			}
			return err
		}
	}

	return nil
}

func (m *Group) validateUsers(formats strfmt.Registry) error {

	if swag.IsZero(m.Users) { // not required
		return nil
	}

	if m.Users != nil {
		if err := m.Users.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("users")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Group) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Group) UnmarshalBinary(b []byte) error {
	var res Group
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
