// Code generated by go-swagger; DO NOT EDIT.

package build

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

// NewReplaceCommentParams creates a new ReplaceCommentParams object
// with the default values initialized.
func NewReplaceCommentParams() *ReplaceCommentParams {
	var ()
	return &ReplaceCommentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceCommentParamsWithTimeout creates a new ReplaceCommentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceCommentParamsWithTimeout(timeout time.Duration) *ReplaceCommentParams {
	var ()
	return &ReplaceCommentParams{

		timeout: timeout,
	}
}

// NewReplaceCommentParamsWithContext creates a new ReplaceCommentParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceCommentParamsWithContext(ctx context.Context) *ReplaceCommentParams {
	var ()
	return &ReplaceCommentParams{

		Context: ctx,
	}
}

// NewReplaceCommentParamsWithHTTPClient creates a new ReplaceCommentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceCommentParamsWithHTTPClient(client *http.Client) *ReplaceCommentParams {
	var ()
	return &ReplaceCommentParams{
		HTTPClient: client,
	}
}

/*ReplaceCommentParams contains all the parameters to send to the API endpoint
for the replace comment operation typically these are written to a http.Request
*/
type ReplaceCommentParams struct {

	/*Body*/
	Body string
	/*BuildLocator*/
	BuildLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace comment params
func (o *ReplaceCommentParams) WithTimeout(timeout time.Duration) *ReplaceCommentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace comment params
func (o *ReplaceCommentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace comment params
func (o *ReplaceCommentParams) WithContext(ctx context.Context) *ReplaceCommentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace comment params
func (o *ReplaceCommentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace comment params
func (o *ReplaceCommentParams) WithHTTPClient(client *http.Client) *ReplaceCommentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace comment params
func (o *ReplaceCommentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace comment params
func (o *ReplaceCommentParams) WithBody(body string) *ReplaceCommentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace comment params
func (o *ReplaceCommentParams) SetBody(body string) {
	o.Body = body
}

// WithBuildLocator adds the buildLocator to the replace comment params
func (o *ReplaceCommentParams) WithBuildLocator(buildLocator string) *ReplaceCommentParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the replace comment params
func (o *ReplaceCommentParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceCommentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
