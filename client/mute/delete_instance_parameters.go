// Code generated by go-swagger; DO NOT EDIT.

package mute

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

// NewDeleteInstanceParams creates a new DeleteInstanceParams object
// with the default values initialized.
func NewDeleteInstanceParams() *DeleteInstanceParams {
	var ()
	return &DeleteInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInstanceParamsWithTimeout creates a new DeleteInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInstanceParamsWithTimeout(timeout time.Duration) *DeleteInstanceParams {
	var ()
	return &DeleteInstanceParams{

		timeout: timeout,
	}
}

// NewDeleteInstanceParamsWithContext creates a new DeleteInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInstanceParamsWithContext(ctx context.Context) *DeleteInstanceParams {
	var ()
	return &DeleteInstanceParams{

		Context: ctx,
	}
}

// NewDeleteInstanceParamsWithHTTPClient creates a new DeleteInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInstanceParamsWithHTTPClient(client *http.Client) *DeleteInstanceParams {
	var ()
	return &DeleteInstanceParams{
		HTTPClient: client,
	}
}

/*DeleteInstanceParams contains all the parameters to send to the API endpoint
for the delete instance operation typically these are written to a http.Request
*/
type DeleteInstanceParams struct {

	/*Body*/
	Body string
	/*MuteLocator*/
	MuteLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete instance params
func (o *DeleteInstanceParams) WithTimeout(timeout time.Duration) *DeleteInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete instance params
func (o *DeleteInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete instance params
func (o *DeleteInstanceParams) WithContext(ctx context.Context) *DeleteInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete instance params
func (o *DeleteInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete instance params
func (o *DeleteInstanceParams) WithHTTPClient(client *http.Client) *DeleteInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete instance params
func (o *DeleteInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the delete instance params
func (o *DeleteInstanceParams) WithBody(body string) *DeleteInstanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the delete instance params
func (o *DeleteInstanceParams) SetBody(body string) {
	o.Body = body
}

// WithMuteLocator adds the muteLocator to the delete instance params
func (o *DeleteInstanceParams) WithMuteLocator(muteLocator string) *DeleteInstanceParams {
	o.SetMuteLocator(muteLocator)
	return o
}

// SetMuteLocator adds the muteLocator to the delete instance params
func (o *DeleteInstanceParams) SetMuteLocator(muteLocator string) {
	o.MuteLocator = muteLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param muteLocator
	if err := r.SetPathParam("muteLocator", o.MuteLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}