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

// NewDeleteStepParams creates a new DeleteStepParams object
// with the default values initialized.
func NewDeleteStepParams() *DeleteStepParams {
	var ()
	return &DeleteStepParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteStepParamsWithTimeout creates a new DeleteStepParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteStepParamsWithTimeout(timeout time.Duration) *DeleteStepParams {
	var ()
	return &DeleteStepParams{

		timeout: timeout,
	}
}

// NewDeleteStepParamsWithContext creates a new DeleteStepParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteStepParamsWithContext(ctx context.Context) *DeleteStepParams {
	var ()
	return &DeleteStepParams{

		Context: ctx,
	}
}

// NewDeleteStepParamsWithHTTPClient creates a new DeleteStepParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteStepParamsWithHTTPClient(client *http.Client) *DeleteStepParams {
	var ()
	return &DeleteStepParams{
		HTTPClient: client,
	}
}

/*DeleteStepParams contains all the parameters to send to the API endpoint
for the delete step operation typically these are written to a http.Request
*/
type DeleteStepParams struct {

	/*BtLocator*/
	BtLocator string
	/*StepID*/
	StepID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete step params
func (o *DeleteStepParams) WithTimeout(timeout time.Duration) *DeleteStepParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete step params
func (o *DeleteStepParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete step params
func (o *DeleteStepParams) WithContext(ctx context.Context) *DeleteStepParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete step params
func (o *DeleteStepParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete step params
func (o *DeleteStepParams) WithHTTPClient(client *http.Client) *DeleteStepParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete step params
func (o *DeleteStepParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the delete step params
func (o *DeleteStepParams) WithBtLocator(btLocator string) *DeleteStepParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the delete step params
func (o *DeleteStepParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithStepID adds the stepID to the delete step params
func (o *DeleteStepParams) WithStepID(stepID string) *DeleteStepParams {
	o.SetStepID(stepID)
	return o
}

// SetStepID adds the stepId to the delete step params
func (o *DeleteStepParams) SetStepID(stepID string) {
	o.StepID = stepID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStepParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
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
