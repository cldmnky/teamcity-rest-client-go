// Code generated by go-swagger; DO NOT EDIT.

package debug

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

// NewGetScrambledParams creates a new GetScrambledParams object
// with the default values initialized.
func NewGetScrambledParams() *GetScrambledParams {
	var ()
	return &GetScrambledParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScrambledParamsWithTimeout creates a new GetScrambledParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScrambledParamsWithTimeout(timeout time.Duration) *GetScrambledParams {
	var ()
	return &GetScrambledParams{

		timeout: timeout,
	}
}

// NewGetScrambledParamsWithContext creates a new GetScrambledParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScrambledParamsWithContext(ctx context.Context) *GetScrambledParams {
	var ()
	return &GetScrambledParams{

		Context: ctx,
	}
}

// NewGetScrambledParamsWithHTTPClient creates a new GetScrambledParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScrambledParamsWithHTTPClient(client *http.Client) *GetScrambledParams {
	var ()
	return &GetScrambledParams{
		HTTPClient: client,
	}
}

/*GetScrambledParams contains all the parameters to send to the API endpoint
for the get scrambled operation typically these are written to a http.Request
*/
type GetScrambledParams struct {

	/*Value*/
	Value *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scrambled params
func (o *GetScrambledParams) WithTimeout(timeout time.Duration) *GetScrambledParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scrambled params
func (o *GetScrambledParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scrambled params
func (o *GetScrambledParams) WithContext(ctx context.Context) *GetScrambledParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scrambled params
func (o *GetScrambledParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scrambled params
func (o *GetScrambledParams) WithHTTPClient(client *http.Client) *GetScrambledParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scrambled params
func (o *GetScrambledParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValue adds the value to the get scrambled params
func (o *GetScrambledParams) WithValue(value *string) *GetScrambledParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the get scrambled params
func (o *GetScrambledParams) SetValue(value *string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *GetScrambledParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Value != nil {

		// query param value
		var qrValue string
		if o.Value != nil {
			qrValue = *o.Value
		}
		qValue := qrValue
		if qValue != "" {
			if err := r.SetQueryParam("value", qValue); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
