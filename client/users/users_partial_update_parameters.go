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

// NewUsersPartialUpdateParams creates a new UsersPartialUpdateParams object
// with the default values initialized.
func NewUsersPartialUpdateParams() *UsersPartialUpdateParams {
	var ()
	return &UsersPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUsersPartialUpdateParamsWithTimeout creates a new UsersPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUsersPartialUpdateParamsWithTimeout(timeout time.Duration) *UsersPartialUpdateParams {
	var ()
	return &UsersPartialUpdateParams{

		timeout: timeout,
	}
}

// NewUsersPartialUpdateParamsWithContext creates a new UsersPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUsersPartialUpdateParamsWithContext(ctx context.Context) *UsersPartialUpdateParams {
	var ()
	return &UsersPartialUpdateParams{

		Context: ctx,
	}
}

// NewUsersPartialUpdateParamsWithHTTPClient creates a new UsersPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUsersPartialUpdateParamsWithHTTPClient(client *http.Client) *UsersPartialUpdateParams {
	var ()
	return &UsersPartialUpdateParams{
		HTTPClient: client,
	}
}

/*UsersPartialUpdateParams contains all the parameters to send to the API endpoint
for the users partial update operation typically these are written to a http.Request
*/
type UsersPartialUpdateParams struct {

	/*Data*/
	Data UsersPartialUpdateBody
	/*ID
	  A unique integer value identifying this apex user.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the users partial update params
func (o *UsersPartialUpdateParams) WithTimeout(timeout time.Duration) *UsersPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users partial update params
func (o *UsersPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users partial update params
func (o *UsersPartialUpdateParams) WithContext(ctx context.Context) *UsersPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users partial update params
func (o *UsersPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users partial update params
func (o *UsersPartialUpdateParams) WithHTTPClient(client *http.Client) *UsersPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users partial update params
func (o *UsersPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the users partial update params
func (o *UsersPartialUpdateParams) WithData(data UsersPartialUpdateBody) *UsersPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the users partial update params
func (o *UsersPartialUpdateParams) SetData(data UsersPartialUpdateBody) {
	o.Data = data
}

// WithID adds the id to the users partial update params
func (o *UsersPartialUpdateParams) WithID(id int64) *UsersPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users partial update params
func (o *UsersPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UsersPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
