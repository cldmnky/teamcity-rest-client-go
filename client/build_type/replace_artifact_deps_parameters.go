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

// NewReplaceArtifactDepsParams creates a new ReplaceArtifactDepsParams object
// with the default values initialized.
func NewReplaceArtifactDepsParams() *ReplaceArtifactDepsParams {
	var ()
	return &ReplaceArtifactDepsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceArtifactDepsParamsWithTimeout creates a new ReplaceArtifactDepsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceArtifactDepsParamsWithTimeout(timeout time.Duration) *ReplaceArtifactDepsParams {
	var ()
	return &ReplaceArtifactDepsParams{

		timeout: timeout,
	}
}

// NewReplaceArtifactDepsParamsWithContext creates a new ReplaceArtifactDepsParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceArtifactDepsParamsWithContext(ctx context.Context) *ReplaceArtifactDepsParams {
	var ()
	return &ReplaceArtifactDepsParams{

		Context: ctx,
	}
}

// NewReplaceArtifactDepsParamsWithHTTPClient creates a new ReplaceArtifactDepsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceArtifactDepsParamsWithHTTPClient(client *http.Client) *ReplaceArtifactDepsParams {
	var ()
	return &ReplaceArtifactDepsParams{
		HTTPClient: client,
	}
}

/*ReplaceArtifactDepsParams contains all the parameters to send to the API endpoint
for the replace artifact deps operation typically these are written to a http.Request
*/
type ReplaceArtifactDepsParams struct {

	/*Body*/
	Body *models.ArtifactDependencies
	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace artifact deps params
func (o *ReplaceArtifactDepsParams) WithTimeout(timeout time.Duration) *ReplaceArtifactDepsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace artifact deps params
func (o *ReplaceArtifactDepsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace artifact deps params
func (o *ReplaceArtifactDepsParams) WithContext(ctx context.Context) *ReplaceArtifactDepsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace artifact deps params
func (o *ReplaceArtifactDepsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace artifact deps params
func (o *ReplaceArtifactDepsParams) WithHTTPClient(client *http.Client) *ReplaceArtifactDepsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace artifact deps params
func (o *ReplaceArtifactDepsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace artifact deps params
func (o *ReplaceArtifactDepsParams) WithBody(body *models.ArtifactDependencies) *ReplaceArtifactDepsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace artifact deps params
func (o *ReplaceArtifactDepsParams) SetBody(body *models.ArtifactDependencies) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the replace artifact deps params
func (o *ReplaceArtifactDepsParams) WithBtLocator(btLocator string) *ReplaceArtifactDepsParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the replace artifact deps params
func (o *ReplaceArtifactDepsParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the replace artifact deps params
func (o *ReplaceArtifactDepsParams) WithFields(fields *string) *ReplaceArtifactDepsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the replace artifact deps params
func (o *ReplaceArtifactDepsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceArtifactDepsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
