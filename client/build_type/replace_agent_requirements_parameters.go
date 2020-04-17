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

// NewReplaceAgentRequirementsParams creates a new ReplaceAgentRequirementsParams object
// with the default values initialized.
func NewReplaceAgentRequirementsParams() *ReplaceAgentRequirementsParams {
	var ()
	return &ReplaceAgentRequirementsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceAgentRequirementsParamsWithTimeout creates a new ReplaceAgentRequirementsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceAgentRequirementsParamsWithTimeout(timeout time.Duration) *ReplaceAgentRequirementsParams {
	var ()
	return &ReplaceAgentRequirementsParams{

		timeout: timeout,
	}
}

// NewReplaceAgentRequirementsParamsWithContext creates a new ReplaceAgentRequirementsParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceAgentRequirementsParamsWithContext(ctx context.Context) *ReplaceAgentRequirementsParams {
	var ()
	return &ReplaceAgentRequirementsParams{

		Context: ctx,
	}
}

// NewReplaceAgentRequirementsParamsWithHTTPClient creates a new ReplaceAgentRequirementsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceAgentRequirementsParamsWithHTTPClient(client *http.Client) *ReplaceAgentRequirementsParams {
	var ()
	return &ReplaceAgentRequirementsParams{
		HTTPClient: client,
	}
}

/*ReplaceAgentRequirementsParams contains all the parameters to send to the API endpoint
for the replace agent requirements operation typically these are written to a http.Request
*/
type ReplaceAgentRequirementsParams struct {

	/*Body*/
	Body *models.AgentRequirements
	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace agent requirements params
func (o *ReplaceAgentRequirementsParams) WithTimeout(timeout time.Duration) *ReplaceAgentRequirementsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace agent requirements params
func (o *ReplaceAgentRequirementsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace agent requirements params
func (o *ReplaceAgentRequirementsParams) WithContext(ctx context.Context) *ReplaceAgentRequirementsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace agent requirements params
func (o *ReplaceAgentRequirementsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace agent requirements params
func (o *ReplaceAgentRequirementsParams) WithHTTPClient(client *http.Client) *ReplaceAgentRequirementsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace agent requirements params
func (o *ReplaceAgentRequirementsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace agent requirements params
func (o *ReplaceAgentRequirementsParams) WithBody(body *models.AgentRequirements) *ReplaceAgentRequirementsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace agent requirements params
func (o *ReplaceAgentRequirementsParams) SetBody(body *models.AgentRequirements) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the replace agent requirements params
func (o *ReplaceAgentRequirementsParams) WithBtLocator(btLocator string) *ReplaceAgentRequirementsParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the replace agent requirements params
func (o *ReplaceAgentRequirementsParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the replace agent requirements params
func (o *ReplaceAgentRequirementsParams) WithFields(fields *string) *ReplaceAgentRequirementsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the replace agent requirements params
func (o *ReplaceAgentRequirementsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceAgentRequirementsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
