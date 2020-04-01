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
	"github.com/go-openapi/swag"
)

// NewUsersLoginCreateParams creates a new UsersLoginCreateParams object
// with the default values initialized.
func NewUsersLoginCreateParams() *UsersLoginCreateParams {
	var ()
	return &UsersLoginCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersLoginCreateParamsWithTimeout creates a new UsersLoginCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersLoginCreateParamsWithTimeout(timeout time.Duration) *UsersLoginCreateParams {
	var ()
	return &UsersLoginCreateParams{

		timeout: timeout,
	}
}

// NewUsersLoginCreateParamsWithContext creates a new UsersLoginCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersLoginCreateParamsWithContext(ctx context.Context) *UsersLoginCreateParams {
	var ()
	return &UsersLoginCreateParams{

		Context: ctx,
	}
}

// NewUsersLoginCreateParamsWithHTTPClient creates a new UsersLoginCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersLoginCreateParamsWithHTTPClient(client *http.Client) *UsersLoginCreateParams {
	var ()
	return &UsersLoginCreateParams{
		HTTPClient: client,
	}
}

/*UsersLoginCreateParams contains all the parameters to send to the API endpoint
for the users login create operation typically these are written to a http.Request
*/
type UsersLoginCreateParams struct {

	/*Data*/
	Data UsersLoginCreateBody
	/*ID
	  A unique integer value identifying this user.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users login create params
func (o *UsersLoginCreateParams) WithTimeout(timeout time.Duration) *UsersLoginCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users login create params
func (o *UsersLoginCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users login create params
func (o *UsersLoginCreateParams) WithContext(ctx context.Context) *UsersLoginCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users login create params
func (o *UsersLoginCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users login create params
func (o *UsersLoginCreateParams) WithHTTPClient(client *http.Client) *UsersLoginCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users login create params
func (o *UsersLoginCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users login create params
func (o *UsersLoginCreateParams) WithData(data UsersLoginCreateBody) *UsersLoginCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users login create params
func (o *UsersLoginCreateParams) SetData(data UsersLoginCreateBody) {
	o.Data = data
}

// WithID adds the id to the users login create params
func (o *UsersLoginCreateParams) WithID(id int64) *UsersLoginCreateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users login create params
func (o *UsersLoginCreateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersLoginCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Data); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}