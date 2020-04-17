// Code generated by go-swagger; DO NOT EDIT.

package change

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

// NewServeChangesParams creates a new ServeChangesParams object
// with the default values initialized.
func NewServeChangesParams() *ServeChangesParams {
	var ()
	return &ServeChangesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeChangesParamsWithTimeout creates a new ServeChangesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeChangesParamsWithTimeout(timeout time.Duration) *ServeChangesParams {
	var ()
	return &ServeChangesParams{

		timeout: timeout,
	}
}

// NewServeChangesParamsWithContext creates a new ServeChangesParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeChangesParamsWithContext(ctx context.Context) *ServeChangesParams {
	var ()
	return &ServeChangesParams{

		Context: ctx,
	}
}

// NewServeChangesParamsWithHTTPClient creates a new ServeChangesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeChangesParamsWithHTTPClient(client *http.Client) *ServeChangesParams {
	var ()
	return &ServeChangesParams{
		HTTPClient: client,
	}
}

/*ServeChangesParams contains all the parameters to send to the API endpoint
for the serve changes operation typically these are written to a http.Request
*/
type ServeChangesParams struct {

	/*Build*/
	Build *string
	/*BuildType*/
	BuildType *string
	/*Count*/
	Count *int32
	/*Fields*/
	Fields *string
	/*Locator*/
	Locator *string
	/*Project*/
	Project *string
	/*SinceChange*/
	SinceChange *string
	/*Start*/
	Start *int64
	/*VcsRoot*/
	VcsRoot *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve changes params
func (o *ServeChangesParams) WithTimeout(timeout time.Duration) *ServeChangesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve changes params
func (o *ServeChangesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve changes params
func (o *ServeChangesParams) WithContext(ctx context.Context) *ServeChangesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve changes params
func (o *ServeChangesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve changes params
func (o *ServeChangesParams) WithHTTPClient(client *http.Client) *ServeChangesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve changes params
func (o *ServeChangesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuild adds the build to the serve changes params
func (o *ServeChangesParams) WithBuild(build *string) *ServeChangesParams {
	o.SetBuild(build)
	return o
}

// SetBuild adds the build to the serve changes params
func (o *ServeChangesParams) SetBuild(build *string) {
	o.Build = build
}

// WithBuildType adds the buildType to the serve changes params
func (o *ServeChangesParams) WithBuildType(buildType *string) *ServeChangesParams {
	o.SetBuildType(buildType)
	return o
}

// SetBuildType adds the buildType to the serve changes params
func (o *ServeChangesParams) SetBuildType(buildType *string) {
	o.BuildType = buildType
}

// WithCount adds the count to the serve changes params
func (o *ServeChangesParams) WithCount(count *int32) *ServeChangesParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the serve changes params
func (o *ServeChangesParams) SetCount(count *int32) {
	o.Count = count
}

// WithFields adds the fields to the serve changes params
func (o *ServeChangesParams) WithFields(fields *string) *ServeChangesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve changes params
func (o *ServeChangesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the serve changes params
func (o *ServeChangesParams) WithLocator(locator *string) *ServeChangesParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the serve changes params
func (o *ServeChangesParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WithProject adds the project to the serve changes params
func (o *ServeChangesParams) WithProject(project *string) *ServeChangesParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the serve changes params
func (o *ServeChangesParams) SetProject(project *string) {
	o.Project = project
}

// WithSinceChange adds the sinceChange to the serve changes params
func (o *ServeChangesParams) WithSinceChange(sinceChange *string) *ServeChangesParams {
	o.SetSinceChange(sinceChange)
	return o
}

// SetSinceChange adds the sinceChange to the serve changes params
func (o *ServeChangesParams) SetSinceChange(sinceChange *string) {
	o.SinceChange = sinceChange
}

// WithStart adds the start to the serve changes params
func (o *ServeChangesParams) WithStart(start *int64) *ServeChangesParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the serve changes params
func (o *ServeChangesParams) SetStart(start *int64) {
	o.Start = start
}

// WithVcsRoot adds the vcsRoot to the serve changes params
func (o *ServeChangesParams) WithVcsRoot(vcsRoot *string) *ServeChangesParams {
	o.SetVcsRoot(vcsRoot)
	return o
}

// SetVcsRoot adds the vcsRoot to the serve changes params
func (o *ServeChangesParams) SetVcsRoot(vcsRoot *string) {
	o.VcsRoot = vcsRoot
}

// WriteToRequest writes these params to a swagger request
func (o *ServeChangesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Build != nil {

		// query param build
		var qrBuild string
		if o.Build != nil {
			qrBuild = *o.Build
		}
		qBuild := qrBuild
		if qBuild != "" {
			if err := r.SetQueryParam("build", qBuild); err != nil {
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

	if o.Project != nil {

		// query param project
		var qrProject string
		if o.Project != nil {
			qrProject = *o.Project
		}
		qProject := qrProject
		if qProject != "" {
			if err := r.SetQueryParam("project", qProject); err != nil {
				return err
			}
		}

	}

	if o.SinceChange != nil {

		// query param sinceChange
		var qrSinceChange string
		if o.SinceChange != nil {
			qrSinceChange = *o.SinceChange
		}
		qSinceChange := qrSinceChange
		if qSinceChange != "" {
			if err := r.SetQueryParam("sinceChange", qSinceChange); err != nil {
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

	if o.VcsRoot != nil {

		// query param vcsRoot
		var qrVcsRoot string
		if o.VcsRoot != nil {
			qrVcsRoot = *o.VcsRoot
		}
		qVcsRoot := qrVcsRoot
		if qVcsRoot != "" {
			if err := r.SetQueryParam("vcsRoot", qVcsRoot); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
