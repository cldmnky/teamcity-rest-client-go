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

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// NewReplaceVcsRootEntriesParams creates a new ReplaceVcsRootEntriesParams object
// with the default values initialized.
func NewReplaceVcsRootEntriesParams() *ReplaceVcsRootEntriesParams {
	var ()
	return &ReplaceVcsRootEntriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceVcsRootEntriesParamsWithTimeout creates a new ReplaceVcsRootEntriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceVcsRootEntriesParamsWithTimeout(timeout time.Duration) *ReplaceVcsRootEntriesParams {
	var ()
	return &ReplaceVcsRootEntriesParams{

		timeout: timeout,
	}
}

// NewReplaceVcsRootEntriesParamsWithContext creates a new ReplaceVcsRootEntriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceVcsRootEntriesParamsWithContext(ctx context.Context) *ReplaceVcsRootEntriesParams {
	var ()
	return &ReplaceVcsRootEntriesParams{

		Context: ctx,
	}
}

// NewReplaceVcsRootEntriesParamsWithHTTPClient creates a new ReplaceVcsRootEntriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceVcsRootEntriesParamsWithHTTPClient(client *http.Client) *ReplaceVcsRootEntriesParams {
	var ()
	return &ReplaceVcsRootEntriesParams{
		HTTPClient: client,
	}
}

/*ReplaceVcsRootEntriesParams contains all the parameters to send to the API endpoint
for the replace vcs root entries operation typically these are written to a http.Request
*/
type ReplaceVcsRootEntriesParams struct {

	/*Body*/
	Body *models.VcsRootEntries
	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace vcs root entries params
func (o *ReplaceVcsRootEntriesParams) WithTimeout(timeout time.Duration) *ReplaceVcsRootEntriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace vcs root entries params
func (o *ReplaceVcsRootEntriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace vcs root entries params
func (o *ReplaceVcsRootEntriesParams) WithContext(ctx context.Context) *ReplaceVcsRootEntriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace vcs root entries params
func (o *ReplaceVcsRootEntriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace vcs root entries params
func (o *ReplaceVcsRootEntriesParams) WithHTTPClient(client *http.Client) *ReplaceVcsRootEntriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace vcs root entries params
func (o *ReplaceVcsRootEntriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace vcs root entries params
func (o *ReplaceVcsRootEntriesParams) WithBody(body *models.VcsRootEntries) *ReplaceVcsRootEntriesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace vcs root entries params
func (o *ReplaceVcsRootEntriesParams) SetBody(body *models.VcsRootEntries) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the replace vcs root entries params
func (o *ReplaceVcsRootEntriesParams) WithBtLocator(btLocator string) *ReplaceVcsRootEntriesParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the replace vcs root entries params
func (o *ReplaceVcsRootEntriesParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the replace vcs root entries params
func (o *ReplaceVcsRootEntriesParams) WithFields(fields *string) *ReplaceVcsRootEntriesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the replace vcs root entries params
func (o *ReplaceVcsRootEntriesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceVcsRootEntriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
