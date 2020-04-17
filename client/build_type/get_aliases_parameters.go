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

// NewGetAliasesParams creates a new GetAliasesParams object
// with the default values initialized.
func NewGetAliasesParams() *GetAliasesParams {
	var ()
	return &GetAliasesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAliasesParamsWithTimeout creates a new GetAliasesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAliasesParamsWithTimeout(timeout time.Duration) *GetAliasesParams {
	var ()
	return &GetAliasesParams{

		timeout: timeout,
	}
}

// NewGetAliasesParamsWithContext creates a new GetAliasesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAliasesParamsWithContext(ctx context.Context) *GetAliasesParams {
	var ()
	return &GetAliasesParams{

		Context: ctx,
	}
}

// NewGetAliasesParamsWithHTTPClient creates a new GetAliasesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAliasesParamsWithHTTPClient(client *http.Client) *GetAliasesParams {
	var ()
	return &GetAliasesParams{
		HTTPClient: client,
	}
}

/*GetAliasesParams contains all the parameters to send to the API endpoint
for the get aliases operation typically these are written to a http.Request
*/
type GetAliasesParams struct {

	/*BtLocator*/
	BtLocator string
	/*Field*/
	Field string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get aliases params
func (o *GetAliasesParams) WithTimeout(timeout time.Duration) *GetAliasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get aliases params
func (o *GetAliasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get aliases params
func (o *GetAliasesParams) WithContext(ctx context.Context) *GetAliasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get aliases params
func (o *GetAliasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get aliases params
func (o *GetAliasesParams) WithHTTPClient(client *http.Client) *GetAliasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get aliases params
func (o *GetAliasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get aliases params
func (o *GetAliasesParams) WithBtLocator(btLocator string) *GetAliasesParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get aliases params
func (o *GetAliasesParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithField adds the field to the get aliases params
func (o *GetAliasesParams) WithField(field string) *GetAliasesParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the get aliases params
func (o *GetAliasesParams) SetField(field string) {
	o.Field = field
}

// WriteToRequest writes these params to a swagger request
func (o *GetAliasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
