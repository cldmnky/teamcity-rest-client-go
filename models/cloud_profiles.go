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

// CloudProfiles cloud profiles
//
// swagger:model cloudProfiles
type CloudProfiles struct {

	// cloud profile
	CloudProfile []*CloudProfile `json:"cloudProfile"`

	// count
	Count int32 `json:"count,omitempty" xml:"count"`

	// href
	Href string `json:"href,omitempty" xml:"href"`

	// next href
	NextHref string `json:"nextHref,omitempty" xml:"nextHref"`

	// prev href
	PrevHref string `json:"prevHref,omitempty" xml:"prevHref"`
}

// Validate validates this cloud profiles
func (m *CloudProfiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudProfiles) validateCloudProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudProfile) { // not required
		return nil
	}

	for i := 0; i < len(m.CloudProfile); i++ {
		if swag.IsZero(m.CloudProfile[i]) { // not required
			continue
		}

		if m.CloudProfile[i] != nil {
			if err := m.CloudProfile[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloudProfile" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudProfiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudProfiles) UnmarshalBinary(b []byte) error {
	var res CloudProfiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}