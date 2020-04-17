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

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// NewSetAllowedRunConfigurationsParams creates a new SetAllowedRunConfigurationsParams object
// with the default values initialized.
func NewSetAllowedRunConfigurationsParams() *SetAllowedRunConfigurationsParams {
	var ()
	return &SetAllowedRunConfigurationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetAllowedRunConfigurationsParamsWithTimeout creates a new SetAllowedRunConfigurationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetAllowedRunConfigurationsParamsWithTimeout(timeout time.Duration) *SetAllowedRunConfigurationsParams {
	var ()
	return &SetAllowedRunConfigurationsParams{

		timeout: timeout,
	}
}

// NewSetAllowedRunConfigurationsParamsWithContext creates a new SetAllowedRunConfigurationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetAllowedRunConfigurationsParamsWithContext(ctx context.Context) *SetAllowedRunConfigurationsParams {
	var ()
	return &SetAllowedRunConfigurationsParams{

		Context: ctx,
	}
}

// NewSetAllowedRunConfigurationsParamsWithHTTPClient creates a new SetAllowedRunConfigurationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetAllowedRunConfigurationsParamsWithHTTPClient(client *http.Client) *SetAllowedRunConfigurationsParams {
	var ()
	return &SetAllowedRunConfigurationsParams{
		HTTPClient: client,
	}
}

/*SetAllowedRunConfigurationsParams contains all the parameters to send to the API endpoint
for the set allowed run configurations operation typically these are written to a http.Request
*/
type SetAllowedRunConfigurationsParams struct {

	/*AgentLocator*/
	AgentLocator string
	/*Body*/
	Body *models.CompatibilityPolicy
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set allowed run configurations params
func (o *SetAllowedRunConfigurationsParams) WithTimeout(timeout time.Duration) *SetAllowedRunConfigurationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set allowed run configurations params
func (o *SetAllowedRunConfigurationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set allowed run configurations params
func (o *SetAllowedRunConfigurationsParams) WithContext(ctx context.Context) *SetAllowedRunConfigurationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set allowed run configurations params
func (o *SetAllowedRunConfigurationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set allowed run configurations params
func (o *SetAllowedRunConfigurationsParams) WithHTTPClient(client *http.Client) *SetAllowedRunConfigurationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set allowed run configurations params
func (o *SetAllowedRunConfigurationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentLocator adds the agentLocator to the set allowed run configurations params
func (o *SetAllowedRunConfigurationsParams) WithAgentLocator(agentLocator string) *SetAllowedRunConfigurationsParams {
	o.SetAgentLocator(agentLocator)
	return o
}

// SetAgentLocator adds the agentLocator to the set allowed run configurations params
func (o *SetAllowedRunConfigurationsParams) SetAgentLocator(agentLocator string) {
	o.AgentLocator = agentLocator
}

// WithBody adds the body to the set allowed run configurations params
func (o *SetAllowedRunConfigurationsParams) WithBody(body *models.CompatibilityPolicy) *SetAllowedRunConfigurationsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set allowed run configurations params
func (o *SetAllowedRunConfigurationsParams) SetBody(body *models.CompatibilityPolicy) {
	o.Body = body
}

// WithFields adds the fields to the set allowed run configurations params
func (o *SetAllowedRunConfigurationsParams) WithFields(fields *string) *SetAllowedRunConfigurationsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the set allowed run configurations params
func (o *SetAllowedRunConfigurationsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *SetAllowedRunConfigurationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param agentLocator
	if err := r.SetPathParam("agentLocator", o.AgentLocator); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
