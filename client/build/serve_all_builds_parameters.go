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
	"github.com/go-openapi/swag"
)

// NewServeAllBuildsParams creates a new ServeAllBuildsParams object
// with the default values initialized.
func NewServeAllBuildsParams() *ServeAllBuildsParams {
	var ()
	return &ServeAllBuildsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeAllBuildsParamsWithTimeout creates a new ServeAllBuildsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeAllBuildsParamsWithTimeout(timeout time.Duration) *ServeAllBuildsParams {
	var ()
	return &ServeAllBuildsParams{

		timeout: timeout,
	}
}

// NewServeAllBuildsParamsWithContext creates a new ServeAllBuildsParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeAllBuildsParamsWithContext(ctx context.Context) *ServeAllBuildsParams {
	var ()
	return &ServeAllBuildsParams{

		Context: ctx,
	}
}

// NewServeAllBuildsParamsWithHTTPClient creates a new ServeAllBuildsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeAllBuildsParamsWithHTTPClient(client *http.Client) *ServeAllBuildsParams {
	var ()
	return &ServeAllBuildsParams{
		HTTPClient: client,
	}
}

/*ServeAllBuildsParams contains all the parameters to send to the API endpoint
for the serve all builds operation typically these are written to a http.Request
*/
type ServeAllBuildsParams struct {

	/*AgentName*/
	AgentName *string
	/*BuildType*/
	BuildType *string
	/*Count*/
	Count *int32
	/*Fields*/
	Fields *string
	/*IncludeCanceled*/
	IncludeCanceled *bool
	/*IncludePersonal*/
	IncludePersonal *bool
	/*Locator*/
	Locator *string
	/*OnlyPinned*/
	OnlyPinned *bool
	/*SinceBuild*/
	SinceBuild *string
	/*SinceDate*/
	SinceDate *string
	/*Start*/
	Start *int64
	/*Status*/
	Status *string
	/*Tag*/
	Tag []string
	/*TriggeredByUser*/
	TriggeredByUser *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve all builds params
func (o *ServeAllBuildsParams) WithTimeout(timeout time.Duration) *ServeAllBuildsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve all builds params
func (o *ServeAllBuildsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve all builds params
func (o *ServeAllBuildsParams) WithContext(ctx context.Context) *ServeAllBuildsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve all builds params
func (o *ServeAllBuildsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve all builds params
func (o *ServeAllBuildsParams) WithHTTPClient(client *http.Client) *ServeAllBuildsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve all builds params
func (o *ServeAllBuildsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAgentName adds the agentName to the serve all builds params
func (o *ServeAllBuildsParams) WithAgentName(agentName *string) *ServeAllBuildsParams {
	o.SetAgentName(agentName)
	return o
}

// SetAgentName adds the agentName to the serve all builds params
func (o *ServeAllBuildsParams) SetAgentName(agentName *string) {
	o.AgentName = agentName
}

// WithBuildType adds the buildType to the serve all builds params
func (o *ServeAllBuildsParams) WithBuildType(buildType *string) *ServeAllBuildsParams {
	o.SetBuildType(buildType)
	return o
}

// SetBuildType adds the buildType to the serve all builds params
func (o *ServeAllBuildsParams) SetBuildType(buildType *string) {
	o.BuildType = buildType
}

// WithCount adds the count to the serve all builds params
func (o *ServeAllBuildsParams) WithCount(count *int32) *ServeAllBuildsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the serve all builds params
func (o *ServeAllBuildsParams) SetCount(count *int32) {
	o.Count = count
}

// WithFields adds the fields to the serve all builds params
func (o *ServeAllBuildsParams) WithFields(fields *string) *ServeAllBuildsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve all builds params
func (o *ServeAllBuildsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithIncludeCanceled adds the includeCanceled to the serve all builds params
func (o *ServeAllBuildsParams) WithIncludeCanceled(includeCanceled *bool) *ServeAllBuildsParams {
	o.SetIncludeCanceled(includeCanceled)
	return o
}

// SetIncludeCanceled adds the includeCanceled to the serve all builds params
func (o *ServeAllBuildsParams) SetIncludeCanceled(includeCanceled *bool) {
	o.IncludeCanceled = includeCanceled
}

// WithIncludePersonal adds the includePersonal to the serve all builds params
func (o *ServeAllBuildsParams) WithIncludePersonal(includePersonal *bool) *ServeAllBuildsParams {
	o.SetIncludePersonal(includePersonal)
	return o
}

// SetIncludePersonal adds the includePersonal to the serve all builds params
func (o *ServeAllBuildsParams) SetIncludePersonal(includePersonal *bool) {
	o.IncludePersonal = includePersonal
}

// WithLocator adds the locator to the serve all builds params
func (o *ServeAllBuildsParams) WithLocator(locator *string) *ServeAllBuildsParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the serve all builds params
func (o *ServeAllBuildsParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WithOnlyPinned adds the onlyPinned to the serve all builds params
func (o *ServeAllBuildsParams) WithOnlyPinned(onlyPinned *bool) *ServeAllBuildsParams {
	o.SetOnlyPinned(onlyPinned)
	return o
}

// SetOnlyPinned adds the onlyPinned to the serve all builds params
func (o *ServeAllBuildsParams) SetOnlyPinned(onlyPinned *bool) {
	o.OnlyPinned = onlyPinned
}

// WithSinceBuild adds the sinceBuild to the serve all builds params
func (o *ServeAllBuildsParams) WithSinceBuild(sinceBuild *string) *ServeAllBuildsParams {
	o.SetSinceBuild(sinceBuild)
	return o
}

// SetSinceBuild adds the sinceBuild to the serve all builds params
func (o *ServeAllBuildsParams) SetSinceBuild(sinceBuild *string) {
	o.SinceBuild = sinceBuild
}

// WithSinceDate adds the sinceDate to the serve all builds params
func (o *ServeAllBuildsParams) WithSinceDate(sinceDate *string) *ServeAllBuildsParams {
	o.SetSinceDate(sinceDate)
	return o
}

// SetSinceDate adds the sinceDate to the serve all builds params
func (o *ServeAllBuildsParams) SetSinceDate(sinceDate *string) {
	o.SinceDate = sinceDate
}

// WithStart adds the start to the serve all builds params
func (o *ServeAllBuildsParams) WithStart(start *int64) *ServeAllBuildsParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the serve all builds params
func (o *ServeAllBuildsParams) SetStart(start *int64) {
	o.Start = start
}

// WithStatus adds the status to the serve all builds params
func (o *ServeAllBuildsParams) WithStatus(status *string) *ServeAllBuildsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the serve all builds params
func (o *ServeAllBuildsParams) SetStatus(status *string) {
	o.Status = status
}

// WithTag adds the tag to the serve all builds params
func (o *ServeAllBuildsParams) WithTag(tag []string) *ServeAllBuildsParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the serve all builds params
func (o *ServeAllBuildsParams) SetTag(tag []string) {
	o.Tag = tag
}

// WithTriggeredByUser adds the triggeredByUser to the serve all builds params
func (o *ServeAllBuildsParams) WithTriggeredByUser(triggeredByUser *string) *ServeAllBuildsParams {
	o.SetTriggeredByUser(triggeredByUser)
	return o
}

// SetTriggeredByUser adds the triggeredByUser to the serve all builds params
func (o *ServeAllBuildsParams) SetTriggeredByUser(triggeredByUser *string) {
	o.TriggeredByUser = triggeredByUser
}

// WriteToRequest writes these params to a swagger request
func (o *ServeAllBuildsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AgentName != nil {

		// query param agentName
		var qrAgentName string
		if o.AgentName != nil {
			qrAgentName = *o.AgentName
		}
		qAgentName := qrAgentName
		if qAgentName != "" {
			if err := r.SetQueryParam("agentName", qAgentName); err != nil {
				return err
			}
		}

	}

	if o.BuildType != nil {

		// query param buildType
		var qrBuildType string
		if o.BuildType != nil {
			qrBuildType = *o.BuildType
		}
		qBuildType := qrBuildType
		if qBuildType != "" {
			if err := r.SetQueryParam("buildType", qBuildType); err != nil {
				return err
			}
		}

	}

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

	if o.IncludeCanceled != nil {

		// query param includeCanceled
		var qrIncludeCanceled bool
		if o.IncludeCanceled != nil {
			qrIncludeCanceled = *o.IncludeCanceled
		}
		qIncludeCanceled := swag.FormatBool(qrIncludeCanceled)
		if qIncludeCanceled != "" {
			if err := r.SetQueryParam("includeCanceled", qIncludeCanceled); err != nil {
				return err
			}
		}

	}

	if o.IncludePersonal != nil {

		// query param includePersonal
		var qrIncludePersonal bool
		if o.IncludePersonal != nil {
			qrIncludePersonal = *o.IncludePersonal
		}
		qIncludePersonal := swag.FormatBool(qrIncludePersonal)
		if qIncludePersonal != "" {
			if err := r.SetQueryParam("includePersonal", qIncludePersonal); err != nil {
				return err
			}
		}

	}

	if o.Locator != nil {

		// query param locator
		var qrLocator string
		if o.Locator != nil {
			qrLocator = *o.Locator
		}
		qLocator := qrLocator
		if qLocator != "" {
			if err := r.SetQueryParam("locator", qLocator); err != nil {
				return err
			}
		}

	}

	if o.OnlyPinned != nil {

		// query param onlyPinned
		var qrOnlyPinned bool
		if o.OnlyPinned != nil {
			qrOnlyPinned = *o.OnlyPinned
		}
		qOnlyPinned := swag.FormatBool(qrOnlyPinned)
		if qOnlyPinned != "" {
			if err := r.SetQueryParam("onlyPinned", qOnlyPinned); err != nil {
				return err
			}
		}

	}

	if o.SinceBuild != nil {

		// query param sinceBuild
		var qrSinceBuild string
		if o.SinceBuild != nil {
			qrSinceBuild = *o.SinceBuild
		}
		qSinceBuild := qrSinceBuild
		if qSinceBuild != "" {
			if err := r.SetQueryParam("sinceBuild", qSinceBuild); err != nil {
				return err
			}
		}

	}

	if o.SinceDate != nil {

		// query param sinceDate
		var qrSinceDate string
		if o.SinceDate != nil {
			qrSinceDate = *o.SinceDate
		}
		qSinceDate := qrSinceDate
		if qSinceDate != "" {
			if err := r.SetQueryParam("sinceDate", qSinceDate); err != nil {
				return err
			}
		}

	}

	if o.Start != nil {

		// query param start
		var qrStart int64
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt64(qrStart)
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if o.Status != nil {

		// query param status
		var qrStatus string
		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {
			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}

	}

	valuesTag := o.Tag

	joinedTag := swag.JoinByFormat(valuesTag, "multi")
	// query array param tag
	if err := r.SetQueryParam("tag", joinedTag...); err != nil {
		return err
	}

	if o.TriggeredByUser != nil {

		// query param triggeredByUser
		var qrTriggeredByUser string
		if o.TriggeredByUser != nil {
			qrTriggeredByUser = *o.TriggeredByUser
		}
		qTriggeredByUser := qrTriggeredByUser
		if qTriggeredByUser != "" {
			if err := r.SetQueryParam("triggeredByUser", qTriggeredByUser); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
