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

// NewServeUsersParams creates a new ServeUsersParams object
// with the default values initialized.
func NewServeUsersParams() *ServeUsersParams {
	var ()
	return &ServeUsersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeUsersParamsWithTimeout creates a new ServeUsersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeUsersParamsWithTimeout(timeout time.Duration) *ServeUsersParams {
	var ()
	return &ServeUsersParams{

		timeout: timeout,
	}
}

// NewServeUsersParamsWithContext creates a new ServeUsersParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeUsersParamsWithContext(ctx context.Context) *ServeUsersParams {
	var ()
	return &ServeUsersParams{

		Context: ctx,
	}
}

// NewServeUsersParamsWithHTTPClient creates a new ServeUsersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeUsersParamsWithHTTPClient(client *http.Client) *ServeUsersParams {
	var ()
	return &ServeUsersParams{
		HTTPClient: client,
	}
}

/*ServeUsersParams contains all the parameters to send to the API endpoint
for the serve users operation typically these are written to a http.Request
*/
type ServeUsersParams struct {

	/*Fields*/
	Fields *string
	/*Locator*/
	Locator *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve users params
func (o *ServeUsersParams) WithTimeout(timeout time.Duration) *ServeUsersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve users params
func (o *ServeUsersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve users params
func (o *ServeUsersParams) WithContext(ctx context.Context) *ServeUsersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve users params
func (o *ServeUsersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve users params
func (o *ServeUsersParams) WithHTTPClient(client *http.Client) *ServeUsersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve users params
func (o *ServeUsersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the serve users params
func (o *ServeUsersParams) WithFields(fields *string) *ServeUsersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve users params
func (o *ServeUsersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the serve users params
func (o *ServeUsersParams) WithLocator(locator *string) *ServeUsersParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the serve users params
func (o *ServeUsersParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeUsersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Locator != nil {

		// query param locator
		var qrLocator string
		if o.Locator != nil {
			qrLocator = *o.Locator
		}
		qLocator := qrLocator
		if qLocator != "" {
			if err := r.SetQueryParam("locator", qLocator); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
