// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewChangeRequirementSettingParams creates a new ChangeRequirementSettingParams object
// with the default values initialized.
func NewChangeRequirementSettingParams() *ChangeRequirementSettingParams {
	var ()
	return &ChangeRequirementSettingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeRequirementSettingParamsWithTimeout creates a new ChangeRequirementSettingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeRequirementSettingParamsWithTimeout(timeout time.Duration) *ChangeRequirementSettingParams {
	var ()
	return &ChangeRequirementSettingParams{

		timeout: timeout,
	}
}

// NewChangeRequirementSettingParamsWithContext creates a new ChangeRequirementSettingParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeRequirementSettingParamsWithContext(ctx context.Context) *ChangeRequirementSettingParams {
	var ()
	return &ChangeRequirementSettingParams{

		Context: ctx,
	}
}

// NewChangeRequirementSettingParamsWithHTTPClient creates a new ChangeRequirementSettingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeRequirementSettingParamsWithHTTPClient(client *http.Client) *ChangeRequirementSettingParams {
	var ()
	return &ChangeRequirementSettingParams{
		HTTPClient: client,
	}
}

/*ChangeRequirementSettingParams contains all the parameters to send to the API endpoint
for the change requirement setting operation typically these are written to a http.Request
*/
type ChangeRequirementSettingParams struct {

	/*AgentRequirementLocator*/
	AgentRequirementLocator string
	/*Body*/
	Body string
	/*BtLocator*/
	BtLocator string
	/*FieldName*/
	FieldName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change requirement setting params
func (o *ChangeRequirementSettingParams) WithTimeout(timeout time.Duration) *ChangeRequirementSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change requirement setting params
func (o *ChangeRequirementSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change requirement setting params
func (o *ChangeRequirementSettingParams) WithContext(ctx context.Context) *ChangeRequirementSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change requirement setting params
func (o *ChangeRequirementSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change requirement setting params
func (o *ChangeRequirementSettingParams) WithHTTPClient(client *http.Client) *ChangeRequirementSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change requirement setting params
func (o *ChangeRequirementSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentRequirementLocator adds the agentRequirementLocator to the change requirement setting params
func (o *ChangeRequirementSettingParams) WithAgentRequirementLocator(agentRequirementLocator string) *ChangeRequirementSettingParams {
	o.SetAgentRequirementLocator(agentRequirementLocator)
	return o
}

// SetAgentRequirementLocator adds the agentRequirementLocator to the change requirement setting params
func (o *ChangeRequirementSettingParams) SetAgentRequirementLocator(agentRequirementLocator string) {
	o.AgentRequirementLocator = agentRequirementLocator
}

// WithBody adds the body to the change requirement setting params
func (o *ChangeRequirementSettingParams) WithBody(body string) *ChangeRequirementSettingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change requirement setting params
func (o *ChangeRequirementSettingParams) SetBody(body string) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the change requirement setting params
func (o *ChangeRequirementSettingParams) WithBtLocator(btLocator string) *ChangeRequirementSettingParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the change requirement setting params
func (o *ChangeRequirementSettingParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFieldName adds the fieldName to the change requirement setting params
func (o *ChangeRequirementSettingParams) WithFieldName(fieldName string) *ChangeRequirementSettingParams {
	o.SetFieldName(fieldName)
	return o
}

// SetFieldName adds the fieldName to the change requirement setting params
func (o *ChangeRequirementSettingParams) SetFieldName(fieldName string) {
	o.FieldName = fieldName
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeRequirementSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentRequirementLocator
	if err := r.SetPathParam("agentRequirementLocator", o.AgentRequirementLocator); err != nil {
		return err
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param fieldName
	if err := r.SetPathParam("fieldName", o.FieldName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
