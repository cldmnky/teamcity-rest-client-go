// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Token token
//
// swagger:model token
type Token struct {

	// creation time
	CreationTime string `json:"creationTime,omitempty" xml:"creationTime"`

	// name
	Name string `json:"name,omitempty" xml:"name"`

	// value
	Value string `json:"value,omitempty" xml:"value"`
}

// Validate validates this token
func (m *Token) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Token) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Token) UnmarshalBinary(b []byte) error {
	var res Token
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}