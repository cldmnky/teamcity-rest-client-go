// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

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

// NewGetPoolsParams creates a new GetPoolsParams object
// with the default values initialized.
func NewGetPoolsParams() *GetPoolsParams {
	var ()
	return &GetPoolsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPoolsParamsWithTimeout creates a new GetPoolsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPoolsParamsWithTimeout(timeout time.Duration) *GetPoolsParams {
	var ()
	return &GetPoolsParams{

		timeout: timeout,
	}
}

// NewGetPoolsParamsWithContext creates a new GetPoolsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPoolsParamsWithContext(ctx context.Context) *GetPoolsParams {
	var ()
	return &GetPoolsParams{

		Context: ctx,
	}
}

// NewGetPoolsParamsWithHTTPClient creates a new GetPoolsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPoolsParamsWithHTTPClient(client *http.Client) *GetPoolsParams {
	var ()
	return &GetPoolsParams{
		HTTPClient: client,
	}
}

/*GetPoolsParams contains all the parameters to send to the API endpoint
for the get pools operation typically these are written to a http.Request
*/
type GetPoolsParams struct {

	/*Fields*/
	Fields *string
	/*Locator*/
	Locator *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get pools params
func (o *GetPoolsParams) WithTimeout(timeout time.Duration) *GetPoolsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get pools params
func (o *GetPoolsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get pools params
func (o *GetPoolsParams) WithContext(ctx context.Context) *GetPoolsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get pools params
func (o *GetPoolsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get pools params
func (o *GetPoolsParams) WithHTTPClient(client *http.Client) *GetPoolsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get pools params
func (o *GetPoolsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get pools params
func (o *GetPoolsParams) WithFields(fields *string) *GetPoolsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get pools params
func (o *GetPoolsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the get pools params
func (o *GetPoolsParams) WithLocator(locator *string) *GetPoolsParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the get pools params
func (o *GetPoolsParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WriteToRequest writes these params to a swagger request
func (o *GetPoolsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
