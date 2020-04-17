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
)

// NewGetProjectAgentPoolsParams creates a new GetProjectAgentPoolsParams object
// with the default values initialized.
func NewGetProjectAgentPoolsParams() *GetProjectAgentPoolsParams {
	var ()
	return &GetProjectAgentPoolsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetProjectAgentPoolsParamsWithTimeout creates a new GetProjectAgentPoolsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetProjectAgentPoolsParamsWithTimeout(timeout time.Duration) *GetProjectAgentPoolsParams {
	var ()
	return &GetProjectAgentPoolsParams{

		timeout: timeout,
	}
}

// NewGetProjectAgentPoolsParamsWithContext creates a new GetProjectAgentPoolsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetProjectAgentPoolsParamsWithContext(ctx context.Context) *GetProjectAgentPoolsParams {
	var ()
	return &GetProjectAgentPoolsParams{

		Context: ctx,
	}
}

// NewGetProjectAgentPoolsParamsWithHTTPClient creates a new GetProjectAgentPoolsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetProjectAgentPoolsParamsWithHTTPClient(client *http.Client) *GetProjectAgentPoolsParams {
	var ()
	return &GetProjectAgentPoolsParams{
		HTTPClient: client,
	}
}

/*GetProjectAgentPoolsParams contains all the parameters to send to the API endpoint
for the get project agent pools operation typically these are written to a http.Request
*/
type GetProjectAgentPoolsParams struct {

	/*Fields*/
	Fields *string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get project agent pools params
func (o *GetProjectAgentPoolsParams) WithTimeout(timeout time.Duration) *GetProjectAgentPoolsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get project agent pools params
func (o *GetProjectAgentPoolsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get project agent pools params
func (o *GetProjectAgentPoolsParams) WithContext(ctx context.Context) *GetProjectAgentPoolsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get project agent pools params
func (o *GetProjectAgentPoolsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get project agent pools params
func (o *GetProjectAgentPoolsParams) WithHTTPClient(client *http.Client) *GetProjectAgentPoolsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get project agent pools params
func (o *GetProjectAgentPoolsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get project agent pools params
func (o *GetProjectAgentPoolsParams) WithFields(fields *string) *GetProjectAgentPoolsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get project agent pools params
func (o *GetProjectAgentPoolsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectLocator adds the projectLocator to the get project agent pools params
func (o *GetProjectAgentPoolsParams) WithProjectLocator(projectLocator string) *GetProjectAgentPoolsParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the get project agent pools params
func (o *GetProjectAgentPoolsParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetProjectAgentPoolsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
