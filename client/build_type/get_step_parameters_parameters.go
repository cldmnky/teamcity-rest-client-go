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

// NewGetStepParametersParams creates a new GetStepParametersParams object
// with the default values initialized.
func NewGetStepParametersParams() *GetStepParametersParams {
	var ()
	return &GetStepParametersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStepParametersParamsWithTimeout creates a new GetStepParametersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStepParametersParamsWithTimeout(timeout time.Duration) *GetStepParametersParams {
	var ()
	return &GetStepParametersParams{

		timeout: timeout,
	}
}

// NewGetStepParametersParamsWithContext creates a new GetStepParametersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStepParametersParamsWithContext(ctx context.Context) *GetStepParametersParams {
	var ()
	return &GetStepParametersParams{

		Context: ctx,
	}
}

// NewGetStepParametersParamsWithHTTPClient creates a new GetStepParametersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStepParametersParamsWithHTTPClient(client *http.Client) *GetStepParametersParams {
	var ()
	return &GetStepParametersParams{
		HTTPClient: client,
	}
}

/*GetStepParametersParams contains all the parameters to send to the API endpoint
for the get step parameters operation typically these are written to a http.Request
*/
type GetStepParametersParams struct {

	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string
	/*StepID*/
	StepID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get step parameters params
func (o *GetStepParametersParams) WithTimeout(timeout time.Duration) *GetStepParametersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get step parameters params
func (o *GetStepParametersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get step parameters params
func (o *GetStepParametersParams) WithContext(ctx context.Context) *GetStepParametersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get step parameters params
func (o *GetStepParametersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get step parameters params
func (o *GetStepParametersParams) WithHTTPClient(client *http.Client) *GetStepParametersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get step parameters params
func (o *GetStepParametersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get step parameters params
func (o *GetStepParametersParams) WithBtLocator(btLocator string) *GetStepParametersParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get step parameters params
func (o *GetStepParametersParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the get step parameters params
func (o *GetStepParametersParams) WithFields(fields *string) *GetStepParametersParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get step parameters params
func (o *GetStepParametersParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithStepID adds the stepID to the get step parameters params
func (o *GetStepParametersParams) WithStepID(stepID string) *GetStepParametersParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the get step parameters params
func (o *GetStepParametersParams) SetStepID(stepID string) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStepParametersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param stepId
	if err := r.SetPathParam("stepId", o.StepID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
