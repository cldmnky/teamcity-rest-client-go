// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewServeUserFieldParams creates a new ServeUserFieldParams object
// with the default values initialized.
func NewServeUserFieldParams() *ServeUserFieldParams {
	var ()
	return &ServeUserFieldParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeUserFieldParamsWithTimeout creates a new ServeUserFieldParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeUserFieldParamsWithTimeout(timeout time.Duration) *ServeUserFieldParams {
	var ()
	return &ServeUserFieldParams{

		timeout: timeout,
	}
}

// NewServeUserFieldParamsWithContext creates a new ServeUserFieldParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeUserFieldParamsWithContext(ctx context.Context) *ServeUserFieldParams {
	var ()
	return &ServeUserFieldParams{

		Context: ctx,
	}
}

// NewServeUserFieldParamsWithHTTPClient creates a new ServeUserFieldParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeUserFieldParamsWithHTTPClient(client *http.Client) *ServeUserFieldParams {
	var ()
	return &ServeUserFieldParams{
		HTTPClient: client,
	}
}

/*ServeUserFieldParams contains all the parameters to send to the API endpoint
for the serve user field operation typically these are written to a http.Request
*/
type ServeUserFieldParams struct {

	/*Field*/
	Field string
	/*UserLocator*/
	UserLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve user field params
func (o *ServeUserFieldParams) WithTimeout(timeout time.Duration) *ServeUserFieldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve user field params
func (o *ServeUserFieldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve user field params
func (o *ServeUserFieldParams) WithContext(ctx context.Context) *ServeUserFieldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve user field params
func (o *ServeUserFieldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve user field params
func (o *ServeUserFieldParams) WithHTTPClient(client *http.Client) *ServeUserFieldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve user field params
func (o *ServeUserFieldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithField adds the field to the serve user field params
func (o *ServeUserFieldParams) WithField(field string) *ServeUserFieldParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the serve user field params
func (o *ServeUserFieldParams) SetField(field string) {
	o.Field = field
}

// WithUserLocator adds the userLocator to the serve user field params
func (o *ServeUserFieldParams) WithUserLocator(userLocator string) *ServeUserFieldParams {
	o.SetUserLocator(userLocator)
	return o
}

// SetUserLocator adds the userLocator to the serve user field params
func (o *ServeUserFieldParams) SetUserLocator(userLocator string) {
	o.UserLocator = userLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeUserFieldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
		return err
	}

	// path param userLocator
	if err := r.SetPathParam("userLocator", o.UserLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
