// Code generated by go-swagger; DO NOT EDIT.

package federation

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

// NewServersParams creates a new ServersParams object
// with the default values initialized.
func NewServersParams() *ServersParams {
	var ()
	return &ServersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServersParamsWithTimeout creates a new ServersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServersParamsWithTimeout(timeout time.Duration) *ServersParams {
	var ()
	return &ServersParams{

		timeout: timeout,
	}
}

// NewServersParamsWithContext creates a new ServersParams object
// with the default values initialized, and the ability to set a context for a request
func NewServersParamsWithContext(ctx context.Context) *ServersParams {
	var ()
	return &ServersParams{

		Context: ctx,
	}
}

// NewServersParamsWithHTTPClient creates a new ServersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServersParamsWithHTTPClient(client *http.Client) *ServersParams {
	var ()
	return &ServersParams{
		HTTPClient: client,
	}
}

/*ServersParams contains all the parameters to send to the API endpoint
for the servers operation typically these are written to a http.Request
*/
type ServersParams struct {

	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the servers params
func (o *ServersParams) WithTimeout(timeout time.Duration) *ServersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the servers params
func (o *ServersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the servers params
func (o *ServersParams) WithContext(ctx context.Context) *ServersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the servers params
func (o *ServersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the servers params
func (o *ServersParams) WithHTTPClient(client *http.Client) *ServersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the servers params
func (o *ServersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the servers params
func (o *ServersParams) WithFields(fields *string) *ServersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the servers params
func (o *ServersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ServersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
