// Code generated by go-swagger; DO NOT EDIT.

package vcs_root

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

// NewAddRootParams creates a new AddRootParams object
// with the default values initialized.
func NewAddRootParams() *AddRootParams {
	var ()
	return &AddRootParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddRootParamsWithTimeout creates a new AddRootParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddRootParamsWithTimeout(timeout time.Duration) *AddRootParams {
	var ()
	return &AddRootParams{

		timeout: timeout,
	}
}

// NewAddRootParamsWithContext creates a new AddRootParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddRootParamsWithContext(ctx context.Context) *AddRootParams {
	var ()
	return &AddRootParams{

		Context: ctx,
	}
}

// NewAddRootParamsWithHTTPClient creates a new AddRootParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddRootParamsWithHTTPClient(client *http.Client) *AddRootParams {
	var ()
	return &AddRootParams{
		HTTPClient: client,
	}
}

/*AddRootParams contains all the parameters to send to the API endpoint
for the add root operation typically these are written to a http.Request
*/
type AddRootParams struct {

	/*Body*/
	Body *models.VcsRoot
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add root params
func (o *AddRootParams) WithTimeout(timeout time.Duration) *AddRootParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add root params
func (o *AddRootParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add root params
func (o *AddRootParams) WithContext(ctx context.Context) *AddRootParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add root params
func (o *AddRootParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add root params
func (o *AddRootParams) WithHTTPClient(client *http.Client) *AddRootParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add root params
func (o *AddRootParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add root params
func (o *AddRootParams) WithBody(body *models.VcsRoot) *AddRootParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add root params
func (o *AddRootParams) SetBody(body *models.VcsRoot) {
	o.Body = body
}

// WithFields adds the fields to the add root params
func (o *AddRootParams) WithFields(fields *string) *AddRootParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the add root params
func (o *AddRootParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *AddRootParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
