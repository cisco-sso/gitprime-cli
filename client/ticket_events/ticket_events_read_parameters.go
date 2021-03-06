// Code generated by go-swagger; DO NOT EDIT.

package ticket_events

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

// NewTicketEventsReadParams creates a new TicketEventsReadParams object
// with the default values initialized.
func NewTicketEventsReadParams() *TicketEventsReadParams {
	var ()
	return &TicketEventsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTicketEventsReadParamsWithTimeout creates a new TicketEventsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTicketEventsReadParamsWithTimeout(timeout time.Duration) *TicketEventsReadParams {
	var ()
	return &TicketEventsReadParams{

		timeout: timeout,
	}
}

// NewTicketEventsReadParamsWithContext creates a new TicketEventsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewTicketEventsReadParamsWithContext(ctx context.Context) *TicketEventsReadParams {
	var ()
	return &TicketEventsReadParams{

		Context: ctx,
	}
}

// NewTicketEventsReadParamsWithHTTPClient creates a new TicketEventsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTicketEventsReadParamsWithHTTPClient(client *http.Client) *TicketEventsReadParams {
	var ()
	return &TicketEventsReadParams{
		HTTPClient: client,
	}
}

/*TicketEventsReadParams contains all the parameters to send to the API endpoint
for the ticket events read operation typically these are written to a http.Request
*/
type TicketEventsReadParams struct {

	/*ID
	  A unique integer value identifying this ticket event.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ticket events read params
func (o *TicketEventsReadParams) WithTimeout(timeout time.Duration) *TicketEventsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ticket events read params
func (o *TicketEventsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ticket events read params
func (o *TicketEventsReadParams) WithContext(ctx context.Context) *TicketEventsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ticket events read params
func (o *TicketEventsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ticket events read params
func (o *TicketEventsReadParams) WithHTTPClient(client *http.Client) *TicketEventsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ticket events read params
func (o *TicketEventsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ticket events read params
func (o *TicketEventsReadParams) WithID(id int64) *TicketEventsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ticket events read params
func (o *TicketEventsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TicketEventsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
