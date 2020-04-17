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

// RepositoryState repository state
//
// swagger:model repositoryState
type RepositoryState struct {

	// branch
	Branch []*BranchVersion `json:"branch"`

	// count
	Count int32 `json:"count,omitempty" xml:"count"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty" xml:"timestamp"`
}

// Validate validates this repository state
func (m *RepositoryState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBranch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RepositoryState) validateBranch(formats strfmt.Registry) error {

	if swag.IsZero(m.Branch) { // not required
		return nil
	}

	for i := 0; i < len(m.Branch); i++ {
		if swag.IsZero(m.Branch[i]) { // not required
			continue
		}

		if m.Branch[i] != nil {
			if err := m.Branch[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("branch" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RepositoryState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RepositoryState) UnmarshalBinary(b []byte) error {
	var res RepositoryState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}