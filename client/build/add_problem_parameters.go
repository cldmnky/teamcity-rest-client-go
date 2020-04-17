// Code generated by go-swagger; DO NOT EDIT.

package build

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

// NewAddProblemParams creates a new AddProblemParams object
// with the default values initialized.
func NewAddProblemParams() *AddProblemParams {
	var ()
	return &AddProblemParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddProblemParamsWithTimeout creates a new AddProblemParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddProblemParamsWithTimeout(timeout time.Duration) *AddProblemParams {
	var ()
	return &AddProblemParams{

		timeout: timeout,
	}
}

// NewAddProblemParamsWithContext creates a new AddProblemParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddProblemParamsWithContext(ctx context.Context) *AddProblemParams {
	var ()
	return &AddProblemParams{

		Context: ctx,
	}
}

// NewAddProblemParamsWithHTTPClient creates a new AddProblemParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddProblemParamsWithHTTPClient(client *http.Client) *AddProblemParams {
	var ()
	return &AddProblemParams{
		HTTPClient: client,
	}
}

/*AddProblemParams contains all the parameters to send to the API endpoint
for the add problem operation typically these are written to a http.Request
*/
type AddProblemParams struct {

	/*Body*/
	Body string
	/*BuildLocator*/
	BuildLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add problem params
func (o *AddProblemParams) WithTimeout(timeout time.Duration) *AddProblemParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add problem params
func (o *AddProblemParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add problem params
func (o *AddProblemParams) WithContext(ctx context.Context) *AddProblemParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add problem params
func (o *AddProblemParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add problem params
func (o *AddProblemParams) WithHTTPClient(client *http.Client) *AddProblemParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add problem params
func (o *AddProblemParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add problem params
func (o *AddProblemParams) WithBody(body string) *AddProblemParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add problem params
func (o *AddProblemParams) SetBody(body string) {
	o.Body = body
}

// WithBuildLocator adds the buildLocator to the add problem params
func (o *AddProblemParams) WithBuildLocator(buildLocator string) *AddProblemParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the add problem params
func (o *AddProblemParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the add problem params
func (o *AddProblemParams) WithFields(fields *string) *AddProblemParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add problem params
func (o *AddProblemParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddProblemParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
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