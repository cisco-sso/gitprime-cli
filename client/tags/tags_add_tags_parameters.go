// Code generated by go-swagger; DO NOT EDIT.

package tags

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

// NewTagsAddTagsParams creates a new TagsAddTagsParams object
// with the default values initialized.
func NewTagsAddTagsParams() *TagsAddTagsParams {
	var ()
	return &TagsAddTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTagsAddTagsParamsWithTimeout creates a new TagsAddTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTagsAddTagsParamsWithTimeout(timeout time.Duration) *TagsAddTagsParams {
	var ()
	return &TagsAddTagsParams{

		timeout: timeout,
	}
}

// NewTagsAddTagsParamsWithContext creates a new TagsAddTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewTagsAddTagsParamsWithContext(ctx context.Context) *TagsAddTagsParams {
	var ()
	return &TagsAddTagsParams{

		Context: ctx,
	}
}

// NewTagsAddTagsParamsWithHTTPClient creates a new TagsAddTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTagsAddTagsParamsWithHTTPClient(client *http.Client) *TagsAddTagsParams {
	var ()
	return &TagsAddTagsParams{
		HTTPClient: client,
	}
}

/*TagsAddTagsParams contains all the parameters to send to the API endpoint
for the tags add tags operation typically these are written to a http.Request
*/
type TagsAddTagsParams struct {

	/*Data*/
	Data TagsAddTagsBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tags add tags params
func (o *TagsAddTagsParams) WithTimeout(timeout time.Duration) *TagsAddTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tags add tags params
func (o *TagsAddTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tags add tags params
func (o *TagsAddTagsParams) WithContext(ctx context.Context) *TagsAddTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tags add tags params
func (o *TagsAddTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tags add tags params
func (o *TagsAddTagsParams) WithHTTPClient(client *http.Client) *TagsAddTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tags add tags params
func (o *TagsAddTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the tags add tags params
func (o *TagsAddTagsParams) WithData(data TagsAddTagsBody) *TagsAddTagsParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the tags add tags params
func (o *TagsAddTagsParams) SetData(data TagsAddTagsBody) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *TagsAddTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
