// Code generated by go-swagger; DO NOT EDIT.

package ticket_projects

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

// NewTicketProjectsImportProjectParams creates a new TicketProjectsImportProjectParams object
// with the default values initialized.
func NewTicketProjectsImportProjectParams() *TicketProjectsImportProjectParams {
	var ()
	return &TicketProjectsImportProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTicketProjectsImportProjectParamsWithTimeout creates a new TicketProjectsImportProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTicketProjectsImportProjectParamsWithTimeout(timeout time.Duration) *TicketProjectsImportProjectParams {
	var ()
	return &TicketProjectsImportProjectParams{

		timeout: timeout,
	}
}

// NewTicketProjectsImportProjectParamsWithContext creates a new TicketProjectsImportProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewTicketProjectsImportProjectParamsWithContext(ctx context.Context) *TicketProjectsImportProjectParams {
	var ()
	return &TicketProjectsImportProjectParams{

		Context: ctx,
	}
}

// NewTicketProjectsImportProjectParamsWithHTTPClient creates a new TicketProjectsImportProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTicketProjectsImportProjectParamsWithHTTPClient(client *http.Client) *TicketProjectsImportProjectParams {
	var ()
	return &TicketProjectsImportProjectParams{
		HTTPClient: client,
	}
}

/*TicketProjectsImportProjectParams contains all the parameters to send to the API endpoint
for the ticket projects import project operation typically these are written to a http.Request
*/
type TicketProjectsImportProjectParams struct {

	/*Data*/
	Data TicketProjectsImportProjectBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ticket projects import project params
func (o *TicketProjectsImportProjectParams) WithTimeout(timeout time.Duration) *TicketProjectsImportProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ticket projects import project params
func (o *TicketProjectsImportProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ticket projects import project params
func (o *TicketProjectsImportProjectParams) WithContext(ctx context.Context) *TicketProjectsImportProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ticket projects import project params
func (o *TicketProjectsImportProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ticket projects import project params
func (o *TicketProjectsImportProjectParams) WithHTTPClient(client *http.Client) *TicketProjectsImportProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ticket projects import project params
func (o *TicketProjectsImportProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ticket projects import project params
func (o *TicketProjectsImportProjectParams) WithData(data TicketProjectsImportProjectBody) *TicketProjectsImportProjectParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ticket projects import project params
func (o *TicketProjectsImportProjectParams) SetData(data TicketProjectsImportProjectBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TicketProjectsImportProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
