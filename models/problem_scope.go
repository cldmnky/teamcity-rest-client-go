// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProblemScope problem scope
//
// swagger:model ProblemScope
type ProblemScope struct {

	// build type
	BuildType *BuildType `json:"buildType,omitempty"`

	// build types
	BuildTypes *BuildTypes `json:"buildTypes,omitempty"`

	// project
	Project *Project `json:"project,omitempty"`
}

// Validate validates this problem scope
func (m *ProblemScope) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBuildTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProblemScope) validateBuildType(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildType) { // not required
		return nil
	}

	if m.BuildType != nil {
		if err := m.BuildType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buildType")
			}
			return err
		}
	}

	return nil
}

func (m *ProblemScope) validateBuildTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.BuildTypes) { // not required
		return nil
	}

	if m.BuildTypes != nil {
		if err := m.BuildTypes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("buildTypes")
			}
			return err
		}
	}

	return nil
}

func (m *ProblemScope) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProblemScope) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProblemScope) UnmarshalBinary(b []byte) error {
	var res ProblemScope
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
