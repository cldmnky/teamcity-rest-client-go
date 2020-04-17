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

// CloudInstances cloud instances
//
// swagger:model cloudInstances
type CloudInstances struct {

	// cloud instance
	CloudInstance []*CloudInstance `json:"cloudInstance"`

	// count
	Count int32 `json:"count,omitempty" xml:"count"`

	// href
	Href string `json:"href,omitempty" xml:"href"`

	// next href
	NextHref string `json:"nextHref,omitempty" xml:"nextHref"`

	// prev href
	PrevHref string `json:"prevHref,omitempty" xml:"prevHref"`
}

// Validate validates this cloud instances
func (m *CloudInstances) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudInstance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudInstances) validateCloudInstance(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudInstance) { // not required
		return nil
	}

	for i := 0; i < len(m.CloudInstance); i++ {
		if swag.IsZero(m.CloudInstance[i]) { // not required
			continue
		}

		if m.CloudInstance[i] != nil {
			if err := m.CloudInstance[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloudInstance" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudInstances) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudInstances) UnmarshalBinary(b []byte) error {
	var res CloudInstances
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
