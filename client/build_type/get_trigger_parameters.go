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

// NewGetTriggerParams creates a new GetTriggerParams object
// with the default values initialized.
func NewGetTriggerParams() *GetTriggerParams {
	var ()
	return &GetTriggerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTriggerParamsWithTimeout creates a new GetTriggerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTriggerParamsWithTimeout(timeout time.Duration) *GetTriggerParams {
	var ()
	return &GetTriggerParams{

		timeout: timeout,
	}
}

// NewGetTriggerParamsWithContext creates a new GetTriggerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTriggerParamsWithContext(ctx context.Context) *GetTriggerParams {
	var ()
	return &GetTriggerParams{

		Context: ctx,
	}
}

// NewGetTriggerParamsWithHTTPClient creates a new GetTriggerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTriggerParamsWithHTTPClient(client *http.Client) *GetTriggerParams {
	var ()
	return &GetTriggerParams{
		HTTPClient: client,
	}
}

/*GetTriggerParams contains all the parameters to send to the API endpoint
for the get trigger operation typically these are written to a http.Request
*/
type GetTriggerParams struct {

	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string
	/*TriggerLocator*/
	TriggerLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get trigger params
func (o *GetTriggerParams) WithTimeout(timeout time.Duration) *GetTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get trigger params
func (o *GetTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get trigger params
func (o *GetTriggerParams) WithContext(ctx context.Context) *GetTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get trigger params
func (o *GetTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get trigger params
func (o *GetTriggerParams) WithHTTPClient(client *http.Client) *GetTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get trigger params
func (o *GetTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get trigger params
func (o *GetTriggerParams) WithBtLocator(btLocator string) *GetTriggerParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get trigger params
func (o *GetTriggerParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the get trigger params
func (o *GetTriggerParams) WithFields(fields *string) *GetTriggerParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get trigger params
func (o *GetTriggerParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithTriggerLocator adds the triggerLocator to the get trigger params
func (o *GetTriggerParams) WithTriggerLocator(triggerLocator string) *GetTriggerParams {
	o.SetTriggerLocator(triggerLocator)
	return o
}

// SetTriggerLocator adds the triggerLocator to the get trigger params
func (o *GetTriggerParams) SetTriggerLocator(triggerLocator string) {
	o.TriggerLocator = triggerLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param triggerLocator
	if err := r.SetPathParam("triggerLocator", o.TriggerLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
