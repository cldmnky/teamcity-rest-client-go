// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BackupProcessManager backup process manager
//
// swagger:model BackupProcessManager
type BackupProcessManager struct {

	// current backup process
	CurrentBackupProcess *BackupProcess `json:"currentBackupProcess,omitempty"`
}

// Validate validates this backup process manager
func (m *BackupProcessManager) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentBackupProcess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupProcessManager) validateCurrentBackupProcess(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentBackupProcess) { // not required
		return nil
	}

	if m.CurrentBackupProcess != nil {
		if err := m.CurrentBackupProcess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentBackupProcess")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupProcessManager) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupProcessManager) UnmarshalBinary(b []byte) error {
	var res BackupProcessManager
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
