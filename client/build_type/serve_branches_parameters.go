// Code generated by go-swagger; DO NOT EDIT.

package build_type

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

// NewServeBranchesParams creates a new ServeBranchesParams object
// with the default values initialized.
func NewServeBranchesParams() *ServeBranchesParams {
	var ()
	return &ServeBranchesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeBranchesParamsWithTimeout creates a new ServeBranchesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeBranchesParamsWithTimeout(timeout time.Duration) *ServeBranchesParams {
	var ()
	return &ServeBranchesParams{

		timeout: timeout,
	}
}

// NewServeBranchesParamsWithContext creates a new ServeBranchesParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeBranchesParamsWithContext(ctx context.Context) *ServeBranchesParams {
	var ()
	return &ServeBranchesParams{

		Context: ctx,
	}
}

// NewServeBranchesParamsWithHTTPClient creates a new ServeBranchesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeBranchesParamsWithHTTPClient(client *http.Client) *ServeBranchesParams {
	var ()
	return &ServeBranchesParams{
		HTTPClient: client,
	}
}

/*ServeBranchesParams contains all the parameters to send to the API endpoint
for the serve branches operation typically these are written to a http.Request
*/
type ServeBranchesParams struct {

	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string
	/*Locator*/
	Locator *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve branches params
func (o *ServeBranchesParams) WithTimeout(timeout time.Duration) *ServeBranchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve branches params
func (o *ServeBranchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve branches params
func (o *ServeBranchesParams) WithContext(ctx context.Context) *ServeBranchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve branches params
func (o *ServeBranchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve branches params
func (o *ServeBranchesParams) WithHTTPClient(client *http.Client) *ServeBranchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve branches params
func (o *ServeBranchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the serve branches params
func (o *ServeBranchesParams) WithBtLocator(btLocator string) *ServeBranchesParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the serve branches params
func (o *ServeBranchesParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the serve branches params
func (o *ServeBranchesParams) WithFields(fields *string) *ServeBranchesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve branches params
func (o *ServeBranchesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the serve branches params
func (o *ServeBranchesParams) WithLocator(locator *string) *ServeBranchesParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the serve branches params
func (o *ServeBranchesParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeBranchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
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
