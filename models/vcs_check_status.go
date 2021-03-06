// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VcsCheckStatus vcs check status
//
// swagger:model VcsCheckStatus
type VcsCheckStatus struct {

	// requestor type
	RequestorType string `json:"requestorType,omitempty" xml:"requestorType"`

	// status
	Status string `json:"status,omitempty" xml:"status"`

	// timestamp
	Timestamp string `json:"timestamp,omitempty" xml:"timestamp"`
}

// Validate validates this vcs check status
func (m *VcsCheckStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VcsCheckStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VcsCheckStatus) UnmarshalBinary(b []byte) error {
	var res VcsCheckStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
