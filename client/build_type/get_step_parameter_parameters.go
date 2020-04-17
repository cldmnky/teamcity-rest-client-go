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

// NewGetStepParameterParams creates a new GetStepParameterParams object
// with the default values initialized.
func NewGetStepParameterParams() *GetStepParameterParams {
	var ()
	return &GetStepParameterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStepParameterParamsWithTimeout creates a new GetStepParameterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStepParameterParamsWithTimeout(timeout time.Duration) *GetStepParameterParams {
	var ()
	return &GetStepParameterParams{

		timeout: timeout,
	}
}

// NewGetStepParameterParamsWithContext creates a new GetStepParameterParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStepParameterParamsWithContext(ctx context.Context) *GetStepParameterParams {
	var ()
	return &GetStepParameterParams{

		Context: ctx,
	}
}

// NewGetStepParameterParamsWithHTTPClient creates a new GetStepParameterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStepParameterParamsWithHTTPClient(client *http.Client) *GetStepParameterParams {
	var ()
	return &GetStepParameterParams{
		HTTPClient: client,
	}
}

/*GetStepParameterParams contains all the parameters to send to the API endpoint
for the get step parameter operation typically these are written to a http.Request
*/
type GetStepParameterParams struct {

	/*BtLocator*/
	BtLocator string
	/*ParameterName*/
	ParameterName string
	/*StepID*/
	StepID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get step parameter params
func (o *GetStepParameterParams) WithTimeout(timeout time.Duration) *GetStepParameterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get step parameter params
func (o *GetStepParameterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get step parameter params
func (o *GetStepParameterParams) WithContext(ctx context.Context) *GetStepParameterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get step parameter params
func (o *GetStepParameterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get step parameter params
func (o *GetStepParameterParams) WithHTTPClient(client *http.Client) *GetStepParameterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get step parameter params
func (o *GetStepParameterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get step parameter params
func (o *GetStepParameterParams) WithBtLocator(btLocator string) *GetStepParameterParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get step parameter params
func (o *GetStepParameterParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithParameterName adds the parameterName to the get step parameter params
func (o *GetStepParameterParams) WithParameterName(parameterName string) *GetStepParameterParams {
	o.SetParameterName(parameterName)
	return o
}

// SetParameterName adds the parameterName to the get step parameter params
func (o *GetStepParameterParams) SetParameterName(parameterName string) {
	o.ParameterName = parameterName
}

// WithStepID adds the stepID to the get step parameter params
func (o *GetStepParameterParams) WithStepID(stepID string) *GetStepParameterParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the get step parameter params
func (o *GetStepParameterParams) SetStepID(stepID string) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStepParameterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param parameterName
	if err := r.SetPathParam("parameterName", o.ParameterName); err != nil {
		return err
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
