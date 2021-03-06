// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewTeamsCreateParams creates a new TeamsCreateParams object
// with the default values initialized.
func NewTeamsCreateParams() *TeamsCreateParams {
	var ()
	return &TeamsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTeamsCreateParamsWithTimeout creates a new TeamsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTeamsCreateParamsWithTimeout(timeout time.Duration) *TeamsCreateParams {
	var ()
	return &TeamsCreateParams{

		timeout: timeout,
	}
}

// NewTeamsCreateParamsWithContext creates a new TeamsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewTeamsCreateParamsWithContext(ctx context.Context) *TeamsCreateParams {
	var ()
	return &TeamsCreateParams{

		Context: ctx,
	}
}

// NewTeamsCreateParamsWithHTTPClient creates a new TeamsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTeamsCreateParamsWithHTTPClient(client *http.Client) *TeamsCreateParams {
	var ()
	return &TeamsCreateParams{
		HTTPClient: client,
	}
}

/*TeamsCreateParams contains all the parameters to send to the API endpoint
for the teams create operation typically these are written to a http.Request
*/
type TeamsCreateParams struct {

	/*Data*/
	Data TeamsCreateBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the teams create params
func (o *TeamsCreateParams) WithTimeout(timeout time.Duration) *TeamsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the teams create params
func (o *TeamsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the teams create params
func (o *TeamsCreateParams) WithContext(ctx context.Context) *TeamsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the teams create params
func (o *TeamsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the teams create params
func (o *TeamsCreateParams) WithHTTPClient(client *http.Client) *TeamsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the teams create params
func (o *TeamsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the teams create params
func (o *TeamsCreateParams) WithData(data TeamsCreateBody) *TeamsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the teams create params
func (o *TeamsCreateParams) SetData(data TeamsCreateBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TeamsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
