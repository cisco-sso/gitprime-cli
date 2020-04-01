// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewUsersBulkDeleteParams creates a new UsersBulkDeleteParams object
// with the default values initialized.
func NewUsersBulkDeleteParams() *UsersBulkDeleteParams {
	var ()
	return &UsersBulkDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersBulkDeleteParamsWithTimeout creates a new UsersBulkDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersBulkDeleteParamsWithTimeout(timeout time.Duration) *UsersBulkDeleteParams {
	var ()
	return &UsersBulkDeleteParams{

		timeout: timeout,
	}
}

// NewUsersBulkDeleteParamsWithContext creates a new UsersBulkDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersBulkDeleteParamsWithContext(ctx context.Context) *UsersBulkDeleteParams {
	var ()
	return &UsersBulkDeleteParams{

		Context: ctx,
	}
}

// NewUsersBulkDeleteParamsWithHTTPClient creates a new UsersBulkDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersBulkDeleteParamsWithHTTPClient(client *http.Client) *UsersBulkDeleteParams {
	var ()
	return &UsersBulkDeleteParams{
		HTTPClient: client,
	}
}

/*UsersBulkDeleteParams contains all the parameters to send to the API endpoint
for the users bulk delete operation typically these are written to a http.Request
*/
type UsersBulkDeleteParams struct {

	/*ApexUserIds
	  A json array of ids: [1,2,3]

	*/
	ApexUserIds interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users bulk delete params
func (o *UsersBulkDeleteParams) WithTimeout(timeout time.Duration) *UsersBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users bulk delete params
func (o *UsersBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users bulk delete params
func (o *UsersBulkDeleteParams) WithContext(ctx context.Context) *UsersBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users bulk delete params
func (o *UsersBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users bulk delete params
func (o *UsersBulkDeleteParams) WithHTTPClient(client *http.Client) *UsersBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users bulk delete params
func (o *UsersBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApexUserIds adds the apexUserIds to the users bulk delete params
func (o *UsersBulkDeleteParams) WithApexUserIds(apexUserIds interface{}) *UsersBulkDeleteParams {
	o.SetApexUserIds(apexUserIds)
	return o
}

// SetApexUserIds adds the apexUserIds to the users bulk delete params
func (o *UsersBulkDeleteParams) SetApexUserIds(apexUserIds interface{}) {
	o.ApexUserIds = apexUserIds
}

// WriteToRequest writes these params to a swagger request
func (o *UsersBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApexUserIds != nil {
		if err := r.SetBodyParam(o.ApexUserIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}