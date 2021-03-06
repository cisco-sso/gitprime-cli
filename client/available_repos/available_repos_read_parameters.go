// Code generated by go-swagger; DO NOT EDIT.

package available_repos

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

// NewAvailableReposReadParams creates a new AvailableReposReadParams object
// with the default values initialized.
func NewAvailableReposReadParams() *AvailableReposReadParams {
	var ()
	return &AvailableReposReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAvailableReposReadParamsWithTimeout creates a new AvailableReposReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAvailableReposReadParamsWithTimeout(timeout time.Duration) *AvailableReposReadParams {
	var ()
	return &AvailableReposReadParams{

		timeout: timeout,
	}
}

// NewAvailableReposReadParamsWithContext creates a new AvailableReposReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewAvailableReposReadParamsWithContext(ctx context.Context) *AvailableReposReadParams {
	var ()
	return &AvailableReposReadParams{

		Context: ctx,
	}
}

// NewAvailableReposReadParamsWithHTTPClient creates a new AvailableReposReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAvailableReposReadParamsWithHTTPClient(client *http.Client) *AvailableReposReadParams {
	var ()
	return &AvailableReposReadParams{
		HTTPClient: client,
	}
}

/*AvailableReposReadParams contains all the parameters to send to the API endpoint
for the available repos read operation typically these are written to a http.Request
*/
type AvailableReposReadParams struct {

	/*ID
	  A unique integer value identifying this source project cache.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the available repos read params
func (o *AvailableReposReadParams) WithTimeout(timeout time.Duration) *AvailableReposReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the available repos read params
func (o *AvailableReposReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the available repos read params
func (o *AvailableReposReadParams) WithContext(ctx context.Context) *AvailableReposReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the available repos read params
func (o *AvailableReposReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the available repos read params
func (o *AvailableReposReadParams) WithHTTPClient(client *http.Client) *AvailableReposReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the available repos read params
func (o *AvailableReposReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the available repos read params
func (o *AvailableReposReadParams) WithID(id int64) *AvailableReposReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the available repos read params
func (o *AvailableReposReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AvailableReposReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
