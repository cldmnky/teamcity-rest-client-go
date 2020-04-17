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
	"github.com/go-openapi/swag"
)

// NewExecuteDBQueryParams creates a new ExecuteDBQueryParams object
// with the default values initialized.
func NewExecuteDBQueryParams() *ExecuteDBQueryParams {
	var (
		countDefault          = int32(1000)
		fieldDelimiterDefault = string(", ")
	)
	return &ExecuteDBQueryParams{
		Count:          &countDefault,
		FieldDelimiter: &fieldDelimiterDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewExecuteDBQueryParamsWithTimeout creates a new ExecuteDBQueryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExecuteDBQueryParamsWithTimeout(timeout time.Duration) *ExecuteDBQueryParams {
	var (
		countDefault          = int32(1000)
		fieldDelimiterDefault = string(", ")
	)
	return &ExecuteDBQueryParams{
		Count:          &countDefault,
		FieldDelimiter: &fieldDelimiterDefault,

		timeout: timeout,
	}
}

// NewExecuteDBQueryParamsWithContext creates a new ExecuteDBQueryParams object
// with the default values initialized, and the ability to set a context for a request
func NewExecuteDBQueryParamsWithContext(ctx context.Context) *ExecuteDBQueryParams {
	var (
		countDefault          = int32(1000)
		fieldDelimiterDefault = string(", ")
	)
	return &ExecuteDBQueryParams{
		Count:          &countDefault,
		FieldDelimiter: &fieldDelimiterDefault,

		Context: ctx,
	}
}

// NewExecuteDBQueryParamsWithHTTPClient creates a new ExecuteDBQueryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExecuteDBQueryParamsWithHTTPClient(client *http.Client) *ExecuteDBQueryParams {
	var (
		countDefault          = int32(1000)
		fieldDelimiterDefault = string(", ")
	)
	return &ExecuteDBQueryParams{
		Count:          &countDefault,
		FieldDelimiter: &fieldDelimiterDefault,
		HTTPClient:     client,
	}
}

/*ExecuteDBQueryParams contains all the parameters to send to the API endpoint
for the execute d b query operation typically these are written to a http.Request
*/
type ExecuteDBQueryParams struct {

	/*Count*/
	Count *int32
	/*DataRetrieveQuery*/
	DataRetrieveQuery *string
	/*FieldDelimiter*/
	FieldDelimiter *string
	/*Query*/
	Query string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the execute d b query params
func (o *ExecuteDBQueryParams) WithTimeout(timeout time.Duration) *ExecuteDBQueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the execute d b query params
func (o *ExecuteDBQueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the execute d b query params
func (o *ExecuteDBQueryParams) WithContext(ctx context.Context) *ExecuteDBQueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the execute d b query params
func (o *ExecuteDBQueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the execute d b query params
func (o *ExecuteDBQueryParams) WithHTTPClient(client *http.Client) *ExecuteDBQueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the execute d b query params
func (o *ExecuteDBQueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCount adds the count to the execute d b query params
func (o *ExecuteDBQueryParams) WithCount(count *int32) *ExecuteDBQueryParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the execute d b query params
func (o *ExecuteDBQueryParams) SetCount(count *int32) {
	o.Count = count
}

// WithDataRetrieveQuery adds the dataRetrieveQuery to the execute d b query params
func (o *ExecuteDBQueryParams) WithDataRetrieveQuery(dataRetrieveQuery *string) *ExecuteDBQueryParams {
	o.SetDataRetrieveQuery(dataRetrieveQuery)
	return o
}

// SetDataRetrieveQuery adds the dataRetrieveQuery to the execute d b query params
func (o *ExecuteDBQueryParams) SetDataRetrieveQuery(dataRetrieveQuery *string) {
	o.DataRetrieveQuery = dataRetrieveQuery
}

// WithFieldDelimiter adds the fieldDelimiter to the execute d b query params
func (o *ExecuteDBQueryParams) WithFieldDelimiter(fieldDelimiter *string) *ExecuteDBQueryParams {
	o.SetFieldDelimiter(fieldDelimiter)
	return o
}

// SetFieldDelimiter adds the fieldDelimiter to the execute d b query params
func (o *ExecuteDBQueryParams) SetFieldDelimiter(fieldDelimiter *string) {
	o.FieldDelimiter = fieldDelimiter
}

// WithQuery adds the query to the execute d b query params
func (o *ExecuteDBQueryParams) WithQuery(query string) *ExecuteDBQueryParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the execute d b query params
func (o *ExecuteDBQueryParams) SetQuery(query string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *ExecuteDBQueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Count != nil {

		// query param count
		var qrCount int32
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	if o.DataRetrieveQuery != nil {

		// query param dataRetrieveQuery
		var qrDataRetrieveQuery string
		if o.DataRetrieveQuery != nil {
			qrDataRetrieveQuery = *o.DataRetrieveQuery
		}
		qDataRetrieveQuery := qrDataRetrieveQuery
		if qDataRetrieveQuery != "" {
			if err := r.SetQueryParam("dataRetrieveQuery", qDataRetrieveQuery); err != nil {
				return err
			}
		}

	}

	if o.FieldDelimiter != nil {

		// query param fieldDelimiter
		var qrFieldDelimiter string
		if o.FieldDelimiter != nil {
			qrFieldDelimiter = *o.FieldDelimiter
		}
		qFieldDelimiter := qrFieldDelimiter
		if qFieldDelimiter != "" {
			if err := r.SetQueryParam("fieldDelimiter", qFieldDelimiter); err != nil {
				return err
			}
		}

	}

	// path param query
	if err := r.SetPathParam("query", o.Query); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
