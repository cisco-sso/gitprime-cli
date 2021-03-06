// Code generated by go-swagger; DO NOT EDIT.

package integrations

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

// NewIntegrationsPartialUpdateParams creates a new IntegrationsPartialUpdateParams object
// with the default values initialized.
func NewIntegrationsPartialUpdateParams() *IntegrationsPartialUpdateParams {
	var ()
	return &IntegrationsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIntegrationsPartialUpdateParamsWithTimeout creates a new IntegrationsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIntegrationsPartialUpdateParamsWithTimeout(timeout time.Duration) *IntegrationsPartialUpdateParams {
	var ()
	return &IntegrationsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewIntegrationsPartialUpdateParamsWithContext creates a new IntegrationsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIntegrationsPartialUpdateParamsWithContext(ctx context.Context) *IntegrationsPartialUpdateParams {
	var ()
	return &IntegrationsPartialUpdateParams{

		Context: ctx,
	}
}

// NewIntegrationsPartialUpdateParamsWithHTTPClient creates a new IntegrationsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIntegrationsPartialUpdateParamsWithHTTPClient(client *http.Client) *IntegrationsPartialUpdateParams {
	var ()
	return &IntegrationsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*IntegrationsPartialUpdateParams contains all the parameters to send to the API endpoint
for the integrations partial update operation typically these are written to a http.Request
*/
type IntegrationsPartialUpdateParams struct {

	/*Data*/
	Data IntegrationsPartialUpdateBody
	/*ID
	  A unique integer value identifying this integration.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the integrations partial update params
func (o *IntegrationsPartialUpdateParams) WithTimeout(timeout time.Duration) *IntegrationsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the integrations partial update params
func (o *IntegrationsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the integrations partial update params
func (o *IntegrationsPartialUpdateParams) WithContext(ctx context.Context) *IntegrationsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the integrations partial update params
func (o *IntegrationsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the integrations partial update params
func (o *IntegrationsPartialUpdateParams) WithHTTPClient(client *http.Client) *IntegrationsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the integrations partial update params
func (o *IntegrationsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the integrations partial update params
func (o *IntegrationsPartialUpdateParams) WithData(data IntegrationsPartialUpdateBody) *IntegrationsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the integrations partial update params
func (o *IntegrationsPartialUpdateParams) SetData(data IntegrationsPartialUpdateBody) {
	o.Data = data
}

// WithID adds the id to the integrations partial update params
func (o *IntegrationsPartialUpdateParams) WithID(id int64) *IntegrationsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the integrations partial update params
func (o *IntegrationsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IntegrationsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
