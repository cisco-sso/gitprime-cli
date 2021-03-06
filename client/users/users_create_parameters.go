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

// NewUsersCreateParams creates a new UsersCreateParams object
// with the default values initialized.
func NewUsersCreateParams() *UsersCreateParams {
	var ()
	return &UsersCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersCreateParamsWithTimeout creates a new UsersCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersCreateParamsWithTimeout(timeout time.Duration) *UsersCreateParams {
	var ()
	return &UsersCreateParams{

		timeout: timeout,
	}
}

// NewUsersCreateParamsWithContext creates a new UsersCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersCreateParamsWithContext(ctx context.Context) *UsersCreateParams {
	var ()
	return &UsersCreateParams{

		Context: ctx,
	}
}

// NewUsersCreateParamsWithHTTPClient creates a new UsersCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersCreateParamsWithHTTPClient(client *http.Client) *UsersCreateParams {
	var ()
	return &UsersCreateParams{
		HTTPClient: client,
	}
}

/*UsersCreateParams contains all the parameters to send to the API endpoint
for the users create operation typically these are written to a http.Request
*/
type UsersCreateParams struct {

	/*Data*/
	Data UsersCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users create params
func (o *UsersCreateParams) WithTimeout(timeout time.Duration) *UsersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users create params
func (o *UsersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users create params
func (o *UsersCreateParams) WithContext(ctx context.Context) *UsersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users create params
func (o *UsersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users create params
func (o *UsersCreateParams) WithHTTPClient(client *http.Client) *UsersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users create params
func (o *UsersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users create params
func (o *UsersCreateParams) WithData(data UsersCreateBody) *UsersCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users create params
func (o *UsersCreateParams) SetData(data UsersCreateBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *UsersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
