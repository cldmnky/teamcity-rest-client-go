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

// NewServeUserPropertyParams creates a new ServeUserPropertyParams object
// with the default values initialized.
func NewServeUserPropertyParams() *ServeUserPropertyParams {
	var ()
	return &ServeUserPropertyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeUserPropertyParamsWithTimeout creates a new ServeUserPropertyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeUserPropertyParamsWithTimeout(timeout time.Duration) *ServeUserPropertyParams {
	var ()
	return &ServeUserPropertyParams{

		timeout: timeout,
	}
}

// NewServeUserPropertyParamsWithContext creates a new ServeUserPropertyParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeUserPropertyParamsWithContext(ctx context.Context) *ServeUserPropertyParams {
	var ()
	return &ServeUserPropertyParams{

		Context: ctx,
	}
}

// NewServeUserPropertyParamsWithHTTPClient creates a new ServeUserPropertyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeUserPropertyParamsWithHTTPClient(client *http.Client) *ServeUserPropertyParams {
	var ()
	return &ServeUserPropertyParams{
		HTTPClient: client,
	}
}

/*ServeUserPropertyParams contains all the parameters to send to the API endpoint
for the serve user property operation typically these are written to a http.Request
*/
type ServeUserPropertyParams struct {

	/*Name*/
	Name string
	/*UserLocator*/
	UserLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve user property params
func (o *ServeUserPropertyParams) WithTimeout(timeout time.Duration) *ServeUserPropertyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve user property params
func (o *ServeUserPropertyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve user property params
func (o *ServeUserPropertyParams) WithContext(ctx context.Context) *ServeUserPropertyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve user property params
func (o *ServeUserPropertyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve user property params
func (o *ServeUserPropertyParams) WithHTTPClient(client *http.Client) *ServeUserPropertyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve user property params
func (o *ServeUserPropertyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the serve user property params
func (o *ServeUserPropertyParams) WithName(name string) *ServeUserPropertyParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the serve user property params
func (o *ServeUserPropertyParams) SetName(name string) {
	o.Name = name
}

// WithUserLocator adds the userLocator to the serve user property params
func (o *ServeUserPropertyParams) WithUserLocator(userLocator string) *ServeUserPropertyParams {
	o.SetUserLocator(userLocator)
	return o
}

// SetUserLocator adds the userLocator to the serve user property params
func (o *ServeUserPropertyParams) SetUserLocator(userLocator string) {
	o.UserLocator = userLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeUserPropertyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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
