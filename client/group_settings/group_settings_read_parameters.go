// Code generated by go-swagger; DO NOT EDIT.

package group_settings

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

// NewGroupSettingsReadParams creates a new GroupSettingsReadParams object
// with the default values initialized.
func NewGroupSettingsReadParams() *GroupSettingsReadParams {
	var ()
	return &GroupSettingsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGroupSettingsReadParamsWithTimeout creates a new GroupSettingsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGroupSettingsReadParamsWithTimeout(timeout time.Duration) *GroupSettingsReadParams {
	var ()
	return &GroupSettingsReadParams{

		timeout: timeout,
	}
}

// NewGroupSettingsReadParamsWithContext creates a new GroupSettingsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewGroupSettingsReadParamsWithContext(ctx context.Context) *GroupSettingsReadParams {
	var ()
	return &GroupSettingsReadParams{

		Context: ctx,
	}
}

// NewGroupSettingsReadParamsWithHTTPClient creates a new GroupSettingsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGroupSettingsReadParamsWithHTTPClient(client *http.Client) *GroupSettingsReadParams {
	var ()
	return &GroupSettingsReadParams{
		HTTPClient: client,
	}
}

/*GroupSettingsReadParams contains all the parameters to send to the API endpoint
for the group settings read operation typically these are written to a http.Request
*/
type GroupSettingsReadParams struct {

	/*ID
	  A unique integer value identifying this group settings.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the group settings read params
func (o *GroupSettingsReadParams) WithTimeout(timeout time.Duration) *GroupSettingsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the group settings read params
func (o *GroupSettingsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the group settings read params
func (o *GroupSettingsReadParams) WithContext(ctx context.Context) *GroupSettingsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the group settings read params
func (o *GroupSettingsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the group settings read params
func (o *GroupSettingsReadParams) WithHTTPClient(client *http.Client) *GroupSettingsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the group settings read params
func (o *GroupSettingsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the group settings read params
func (o *GroupSettingsReadParams) WithID(id int64) *GroupSettingsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the group settings read params
func (o *GroupSettingsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GroupSettingsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}