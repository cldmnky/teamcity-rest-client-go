// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewSetProjectAgentPoolsParams creates a new SetProjectAgentPoolsParams object
// with the default values initialized.
func NewSetProjectAgentPoolsParams() *SetProjectAgentPoolsParams {
	var ()
	return &SetProjectAgentPoolsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetProjectAgentPoolsParamsWithTimeout creates a new SetProjectAgentPoolsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetProjectAgentPoolsParamsWithTimeout(timeout time.Duration) *SetProjectAgentPoolsParams {
	var ()
	return &SetProjectAgentPoolsParams{

		timeout: timeout,
	}
}

// NewSetProjectAgentPoolsParamsWithContext creates a new SetProjectAgentPoolsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetProjectAgentPoolsParamsWithContext(ctx context.Context) *SetProjectAgentPoolsParams {
	var ()
	return &SetProjectAgentPoolsParams{

		Context: ctx,
	}
}

// NewSetProjectAgentPoolsParamsWithHTTPClient creates a new SetProjectAgentPoolsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetProjectAgentPoolsParamsWithHTTPClient(client *http.Client) *SetProjectAgentPoolsParams {
	var ()
	return &SetProjectAgentPoolsParams{
		HTTPClient: client,
	}
}

/*SetProjectAgentPoolsParams contains all the parameters to send to the API endpoint
for the set project agent pools operation typically these are written to a http.Request
*/
type SetProjectAgentPoolsParams struct {

	/*Body*/
	Body *models.AgentPools
	/*Fields*/
	Fields *string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set project agent pools params
func (o *SetProjectAgentPoolsParams) WithTimeout(timeout time.Duration) *SetProjectAgentPoolsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set project agent pools params
func (o *SetProjectAgentPoolsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set project agent pools params
func (o *SetProjectAgentPoolsParams) WithContext(ctx context.Context) *SetProjectAgentPoolsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set project agent pools params
func (o *SetProjectAgentPoolsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set project agent pools params
func (o *SetProjectAgentPoolsParams) WithHTTPClient(client *http.Client) *SetProjectAgentPoolsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set project agent pools params
func (o *SetProjectAgentPoolsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set project agent pools params
func (o *SetProjectAgentPoolsParams) WithBody(body *models.AgentPools) *SetProjectAgentPoolsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set project agent pools params
func (o *SetProjectAgentPoolsParams) SetBody(body *models.AgentPools) {
	o.Body = body
}

// WithFields adds the fields to the set project agent pools params
func (o *SetProjectAgentPoolsParams) WithFields(fields *string) *SetProjectAgentPoolsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the set project agent pools params
func (o *SetProjectAgentPoolsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectLocator adds the projectLocator to the set project agent pools params
func (o *SetProjectAgentPoolsParams) WithProjectLocator(projectLocator string) *SetProjectAgentPoolsParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the set project agent pools params
func (o *SetProjectAgentPoolsParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *SetProjectAgentPoolsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}