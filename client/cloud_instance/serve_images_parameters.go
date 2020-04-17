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

// NewServeImagesParams creates a new ServeImagesParams object
// with the default values initialized.
func NewServeImagesParams() *ServeImagesParams {
	var ()
	return &ServeImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeImagesParamsWithTimeout creates a new ServeImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeImagesParamsWithTimeout(timeout time.Duration) *ServeImagesParams {
	var ()
	return &ServeImagesParams{

		timeout: timeout,
	}
}

// NewServeImagesParamsWithContext creates a new ServeImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeImagesParamsWithContext(ctx context.Context) *ServeImagesParams {
	var ()
	return &ServeImagesParams{

		Context: ctx,
	}
}

// NewServeImagesParamsWithHTTPClient creates a new ServeImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeImagesParamsWithHTTPClient(client *http.Client) *ServeImagesParams {
	var ()
	return &ServeImagesParams{
		HTTPClient: client,
	}
}

/*ServeImagesParams contains all the parameters to send to the API endpoint
for the serve images operation typically these are written to a http.Request
*/
type ServeImagesParams struct {

	/*Fields*/
	Fields *string
	/*Locator*/
	Locator *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve images params
func (o *ServeImagesParams) WithTimeout(timeout time.Duration) *ServeImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve images params
func (o *ServeImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve images params
func (o *ServeImagesParams) WithContext(ctx context.Context) *ServeImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve images params
func (o *ServeImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve images params
func (o *ServeImagesParams) WithHTTPClient(client *http.Client) *ServeImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve images params
func (o *ServeImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the serve images params
func (o *ServeImagesParams) WithFields(fields *string) *ServeImagesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve images params
func (o *ServeImagesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the serve images params
func (o *ServeImagesParams) WithLocator(locator *string) *ServeImagesParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the serve images params
func (o *ServeImagesParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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