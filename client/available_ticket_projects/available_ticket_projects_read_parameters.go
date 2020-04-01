// Code generated by go-swagger; DO NOT EDIT.

package available_ticket_projects

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

// NewAvailableTicketProjectsReadParams creates a new AvailableTicketProjectsReadParams object
// with the default values initialized.
func NewAvailableTicketProjectsReadParams() *AvailableTicketProjectsReadParams {
	var ()
	return &AvailableTicketProjectsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAvailableTicketProjectsReadParamsWithTimeout creates a new AvailableTicketProjectsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAvailableTicketProjectsReadParamsWithTimeout(timeout time.Duration) *AvailableTicketProjectsReadParams {
	var ()
	return &AvailableTicketProjectsReadParams{

		timeout: timeout,
	}
}

// NewAvailableTicketProjectsReadParamsWithContext creates a new AvailableTicketProjectsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewAvailableTicketProjectsReadParamsWithContext(ctx context.Context) *AvailableTicketProjectsReadParams {
	var ()
	return &AvailableTicketProjectsReadParams{

		Context: ctx,
	}
}

// NewAvailableTicketProjectsReadParamsWithHTTPClient creates a new AvailableTicketProjectsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAvailableTicketProjectsReadParamsWithHTTPClient(client *http.Client) *AvailableTicketProjectsReadParams {
	var ()
	return &AvailableTicketProjectsReadParams{
		HTTPClient: client,
	}
}

/*AvailableTicketProjectsReadParams contains all the parameters to send to the API endpoint
for the available ticket projects read operation typically these are written to a http.Request
*/
type AvailableTicketProjectsReadParams struct {

	/*ID
	  A unique integer value identifying this source project cache.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the available ticket projects read params
func (o *AvailableTicketProjectsReadParams) WithTimeout(timeout time.Duration) *AvailableTicketProjectsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the available ticket projects read params
func (o *AvailableTicketProjectsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the available ticket projects read params
func (o *AvailableTicketProjectsReadParams) WithContext(ctx context.Context) *AvailableTicketProjectsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the available ticket projects read params
func (o *AvailableTicketProjectsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the available ticket projects read params
func (o *AvailableTicketProjectsReadParams) WithHTTPClient(client *http.Client) *AvailableTicketProjectsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the available ticket projects read params
func (o *AvailableTicketProjectsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the available ticket projects read params
func (o *AvailableTicketProjectsReadParams) WithID(id int64) *AvailableTicketProjectsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the available ticket projects read params
func (o *AvailableTicketProjectsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AvailableTicketProjectsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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