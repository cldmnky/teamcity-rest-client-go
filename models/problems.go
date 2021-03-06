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

// Problems problems
//
// swagger:model problems
type Problems struct {

	// count
	Count int32 `json:"count,omitempty" xml:"count"`

	// next href
	NextHref string `json:"nextHref,omitempty" xml:"nextHref"`

	// prev href
	PrevHref string `json:"prevHref,omitempty" xml:"prevHref"`

	// problem
	Problem []*Problem `json:"problem"`
}

// Validate validates this problems
func (m *Problems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProblem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Problems) validateProblem(formats strfmt.Registry) error {

	if swag.IsZero(m.Problem) { // not required
		return nil
	}

	for i := 0; i < len(m.Problem); i++ {
		if swag.IsZero(m.Problem[i]) { // not required
			continue
		}

		if m.Problem[i] != nil {
			if err := m.Problem[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("problem" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Problems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Problems) UnmarshalBinary(b []byte) error {
	var res Problems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
