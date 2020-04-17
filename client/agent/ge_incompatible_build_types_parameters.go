// Code generated by go-swagger; DO NOT EDIT.

package agent

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

// NewGeIncompatibleBuildTypesParams creates a new GeIncompatibleBuildTypesParams object
// with the default values initialized.
func NewGeIncompatibleBuildTypesParams() *GeIncompatibleBuildTypesParams {
	var ()
	return &GeIncompatibleBuildTypesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGeIncompatibleBuildTypesParamsWithTimeout creates a new GeIncompatibleBuildTypesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGeIncompatibleBuildTypesParamsWithTimeout(timeout time.Duration) *GeIncompatibleBuildTypesParams {
	var ()
	return &GeIncompatibleBuildTypesParams{

		timeout: timeout,
	}
}

// NewGeIncompatibleBuildTypesParamsWithContext creates a new GeIncompatibleBuildTypesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGeIncompatibleBuildTypesParamsWithContext(ctx context.Context) *GeIncompatibleBuildTypesParams {
	var ()
	return &GeIncompatibleBuildTypesParams{

		Context: ctx,
	}
}

// NewGeIncompatibleBuildTypesParamsWithHTTPClient creates a new GeIncompatibleBuildTypesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGeIncompatibleBuildTypesParamsWithHTTPClient(client *http.Client) *GeIncompatibleBuildTypesParams {
	var ()
	return &GeIncompatibleBuildTypesParams{
		HTTPClient: client,
	}
}

/*GeIncompatibleBuildTypesParams contains all the parameters to send to the API endpoint
for the ge incompatible build types operation typically these are written to a http.Request
*/
type GeIncompatibleBuildTypesParams struct {

	/*AgentLocator*/
	AgentLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ge incompatible build types params
func (o *GeIncompatibleBuildTypesParams) WithTimeout(timeout time.Duration) *GeIncompatibleBuildTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ge incompatible build types params
func (o *GeIncompatibleBuildTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ge incompatible build types params
func (o *GeIncompatibleBuildTypesParams) WithContext(ctx context.Context) *GeIncompatibleBuildTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ge incompatible build types params
func (o *GeIncompatibleBuildTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ge incompatible build types params
func (o *GeIncompatibleBuildTypesParams) WithHTTPClient(client *http.Client) *GeIncompatibleBuildTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ge incompatible build types params
func (o *GeIncompatibleBuildTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentLocator adds the agentLocator to the ge incompatible build types params
func (o *GeIncompatibleBuildTypesParams) WithAgentLocator(agentLocator string) *GeIncompatibleBuildTypesParams {
	o.SetAgentLocator(agentLocator)
	return o
}

// SetAgentLocator adds the agentLocator to the ge incompatible build types params
func (o *GeIncompatibleBuildTypesParams) SetAgentLocator(agentLocator string) {
	o.AgentLocator = agentLocator
}

// WithFields adds the fields to the ge incompatible build types params
func (o *GeIncompatibleBuildTypesParams) WithFields(fields *string) *GeIncompatibleBuildTypesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the ge incompatible build types params
func (o *GeIncompatibleBuildTypesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GeIncompatibleBuildTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentLocator
	if err := r.SetPathParam("agentLocator", o.AgentLocator); err != nil {
		return err
	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}