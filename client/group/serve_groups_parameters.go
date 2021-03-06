// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewServeGroupsParams creates a new ServeGroupsParams object
// with the default values initialized.
func NewServeGroupsParams() *ServeGroupsParams {
	var ()
	return &ServeGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeGroupsParamsWithTimeout creates a new ServeGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeGroupsParamsWithTimeout(timeout time.Duration) *ServeGroupsParams {
	var ()
	return &ServeGroupsParams{

		timeout: timeout,
	}
}

// NewServeGroupsParamsWithContext creates a new ServeGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeGroupsParamsWithContext(ctx context.Context) *ServeGroupsParams {
	var ()
	return &ServeGroupsParams{

		Context: ctx,
	}
}

// NewServeGroupsParamsWithHTTPClient creates a new ServeGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeGroupsParamsWithHTTPClient(client *http.Client) *ServeGroupsParams {
	var ()
	return &ServeGroupsParams{
		HTTPClient: client,
	}
}

/*ServeGroupsParams contains all the parameters to send to the API endpoint
for the serve groups operation typically these are written to a http.Request
*/
type ServeGroupsParams struct {

	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve groups params
func (o *ServeGroupsParams) WithTimeout(timeout time.Duration) *ServeGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve groups params
func (o *ServeGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve groups params
func (o *ServeGroupsParams) WithContext(ctx context.Context) *ServeGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve groups params
func (o *ServeGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve groups params
func (o *ServeGroupsParams) WithHTTPClient(client *http.Client) *ServeGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve groups params
func (o *ServeGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the serve groups params
func (o *ServeGroupsParams) WithFields(fields *string) *ServeGroupsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve groups params
func (o *ServeGroupsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ServeGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
