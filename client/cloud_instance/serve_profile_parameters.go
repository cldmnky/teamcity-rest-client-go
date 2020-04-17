// Code generated by go-swagger; DO NOT EDIT.

package cloud_instance

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

// NewServeProfileParams creates a new ServeProfileParams object
// with the default values initialized.
func NewServeProfileParams() *ServeProfileParams {
	var ()
	return &ServeProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeProfileParamsWithTimeout creates a new ServeProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeProfileParamsWithTimeout(timeout time.Duration) *ServeProfileParams {
	var ()
	return &ServeProfileParams{

		timeout: timeout,
	}
}

// NewServeProfileParamsWithContext creates a new ServeProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeProfileParamsWithContext(ctx context.Context) *ServeProfileParams {
	var ()
	return &ServeProfileParams{

		Context: ctx,
	}
}

// NewServeProfileParamsWithHTTPClient creates a new ServeProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeProfileParamsWithHTTPClient(client *http.Client) *ServeProfileParams {
	var ()
	return &ServeProfileParams{
		HTTPClient: client,
	}
}

/*ServeProfileParams contains all the parameters to send to the API endpoint
for the serve profile operation typically these are written to a http.Request
*/
type ServeProfileParams struct {

	/*Fields*/
	Fields *string
	/*ProfileLocator*/
	ProfileLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve profile params
func (o *ServeProfileParams) WithTimeout(timeout time.Duration) *ServeProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve profile params
func (o *ServeProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve profile params
func (o *ServeProfileParams) WithContext(ctx context.Context) *ServeProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve profile params
func (o *ServeProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve profile params
func (o *ServeProfileParams) WithHTTPClient(client *http.Client) *ServeProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve profile params
func (o *ServeProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the serve profile params
func (o *ServeProfileParams) WithFields(fields *string) *ServeProfileParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve profile params
func (o *ServeProfileParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProfileLocator adds the profileLocator to the serve profile params
func (o *ServeProfileParams) WithProfileLocator(profileLocator string) *ServeProfileParams {
	o.SetProfileLocator(profileLocator)
	return o
}

// SetProfileLocator adds the profileLocator to the serve profile params
func (o *ServeProfileParams) SetProfileLocator(profileLocator string) {
	o.ProfileLocator = profileLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param profileLocator
	if err := r.SetPathParam("profileLocator", o.ProfileLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
