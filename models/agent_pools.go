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

// AgentPools agent pools
//
// swagger:model agentPools
type AgentPools struct {

	// agent pool
	AgentPool []*AgentPool `json:"agentPool"`

	// count
	Count int32 `json:"count,omitempty" xml:"count"`

	// href
	Href string `json:"href,omitempty" xml:"href"`

	// next href
	NextHref string `json:"nextHref,omitempty" xml:"nextHref"`

	// prev href
	PrevHref string `json:"prevHref,omitempty" xml:"prevHref"`
}

// Validate validates this agent pools
func (m *AgentPools) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentPool(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AgentPools) validateAgentPool(formats strfmt.Registry) error {

	if swag.IsZero(m.AgentPool) { // not required
		return nil
	}

	for i := 0; i < len(m.AgentPool); i++ {
		if swag.IsZero(m.AgentPool[i]) { // not required
			continue
		}

		if m.AgentPool[i] != nil {
			if err := m.AgentPool[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agentPool" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AgentPools) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AgentPools) UnmarshalBinary(b []byte) error {
	var res AgentPools
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
