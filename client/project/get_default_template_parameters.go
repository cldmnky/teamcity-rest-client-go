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

// NewGetDefaultTemplateParams creates a new GetDefaultTemplateParams object
// with the default values initialized.
func NewGetDefaultTemplateParams() *GetDefaultTemplateParams {
	var ()
	return &GetDefaultTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDefaultTemplateParamsWithTimeout creates a new GetDefaultTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDefaultTemplateParamsWithTimeout(timeout time.Duration) *GetDefaultTemplateParams {
	var ()
	return &GetDefaultTemplateParams{

		timeout: timeout,
	}
}

// NewGetDefaultTemplateParamsWithContext creates a new GetDefaultTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDefaultTemplateParamsWithContext(ctx context.Context) *GetDefaultTemplateParams {
	var ()
	return &GetDefaultTemplateParams{

		Context: ctx,
	}
}

// NewGetDefaultTemplateParamsWithHTTPClient creates a new GetDefaultTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDefaultTemplateParamsWithHTTPClient(client *http.Client) *GetDefaultTemplateParams {
	var ()
	return &GetDefaultTemplateParams{
		HTTPClient: client,
	}
}

/*GetDefaultTemplateParams contains all the parameters to send to the API endpoint
for the get default template operation typically these are written to a http.Request
*/
type GetDefaultTemplateParams struct {

	/*Fields*/
	Fields *string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get default template params
func (o *GetDefaultTemplateParams) WithTimeout(timeout time.Duration) *GetDefaultTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get default template params
func (o *GetDefaultTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get default template params
func (o *GetDefaultTemplateParams) WithContext(ctx context.Context) *GetDefaultTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get default template params
func (o *GetDefaultTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get default template params
func (o *GetDefaultTemplateParams) WithHTTPClient(client *http.Client) *GetDefaultTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get default template params
func (o *GetDefaultTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get default template params
func (o *GetDefaultTemplateParams) WithFields(fields *string) *GetDefaultTemplateParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get default template params
func (o *GetDefaultTemplateParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectLocator adds the projectLocator to the get default template params
func (o *GetDefaultTemplateParams) WithProjectLocator(projectLocator string) *GetDefaultTemplateParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the get default template params
func (o *GetDefaultTemplateParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetDefaultTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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