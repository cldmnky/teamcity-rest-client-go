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

// MultipleOperationResult multiple operation result
//
// swagger:model multipleOperationResult
type MultipleOperationResult struct {

	// count
	Count int32 `json:"count,omitempty" xml:"count"`

	// error count
	ErrorCount int32 `json:"errorCount,omitempty" xml:"errorCount"`

	// operation result
	OperationResult []*OperationResult `json:"operationResult"`
}

// Validate validates this multiple operation result
func (m *MultipleOperationResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperationResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultipleOperationResult) validateOperationResult(formats strfmt.Registry) error {

	if swag.IsZero(m.OperationResult) { // not required
		return nil
	}

	for i := 0; i < len(m.OperationResult); i++ {
		if swag.IsZero(m.OperationResult[i]) { // not required
			continue
		}

		if m.OperationResult[i] != nil {
			if err := m.OperationResult[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("operationResult" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MultipleOperationResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultipleOperationResult) UnmarshalBinary(b []byte) error {
	var res MultipleOperationResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
