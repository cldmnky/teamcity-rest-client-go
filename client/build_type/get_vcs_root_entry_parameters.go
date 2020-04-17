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

// NewGetVcsRootEntryParams creates a new GetVcsRootEntryParams object
// with the default values initialized.
func NewGetVcsRootEntryParams() *GetVcsRootEntryParams {
	var ()
	return &GetVcsRootEntryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVcsRootEntryParamsWithTimeout creates a new GetVcsRootEntryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVcsRootEntryParamsWithTimeout(timeout time.Duration) *GetVcsRootEntryParams {
	var ()
	return &GetVcsRootEntryParams{

		timeout: timeout,
	}
}

// NewGetVcsRootEntryParamsWithContext creates a new GetVcsRootEntryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVcsRootEntryParamsWithContext(ctx context.Context) *GetVcsRootEntryParams {
	var ()
	return &GetVcsRootEntryParams{

		Context: ctx,
	}
}

// NewGetVcsRootEntryParamsWithHTTPClient creates a new GetVcsRootEntryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVcsRootEntryParamsWithHTTPClient(client *http.Client) *GetVcsRootEntryParams {
	var ()
	return &GetVcsRootEntryParams{
		HTTPClient: client,
	}
}

/*GetVcsRootEntryParams contains all the parameters to send to the API endpoint
for the get vcs root entry operation typically these are written to a http.Request
*/
type GetVcsRootEntryParams struct {

	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string
	/*VcsRootLocator*/
	VcsRootLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get vcs root entry params
func (o *GetVcsRootEntryParams) WithTimeout(timeout time.Duration) *GetVcsRootEntryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get vcs root entry params
func (o *GetVcsRootEntryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get vcs root entry params
func (o *GetVcsRootEntryParams) WithContext(ctx context.Context) *GetVcsRootEntryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get vcs root entry params
func (o *GetVcsRootEntryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get vcs root entry params
func (o *GetVcsRootEntryParams) WithHTTPClient(client *http.Client) *GetVcsRootEntryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get vcs root entry params
func (o *GetVcsRootEntryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get vcs root entry params
func (o *GetVcsRootEntryParams) WithBtLocator(btLocator string) *GetVcsRootEntryParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get vcs root entry params
func (o *GetVcsRootEntryParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the get vcs root entry params
func (o *GetVcsRootEntryParams) WithFields(fields *string) *GetVcsRootEntryParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get vcs root entry params
func (o *GetVcsRootEntryParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithVcsRootLocator adds the vcsRootLocator to the get vcs root entry params
func (o *GetVcsRootEntryParams) WithVcsRootLocator(vcsRootLocator string) *GetVcsRootEntryParams {
	o.SetVcsRootLocator(vcsRootLocator)
	return o
}

// SetVcsRootLocator adds the vcsRootLocator to the get vcs root entry params
func (o *GetVcsRootEntryParams) SetVcsRootLocator(vcsRootLocator string) {
	o.VcsRootLocator = vcsRootLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetVcsRootEntryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param vcsRootLocator
	if err := r.SetPathParam("vcsRootLocator", o.VcsRootLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
