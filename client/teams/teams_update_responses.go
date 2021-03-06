// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TeamsUpdateReader is a Reader for the TeamsUpdate structure.
type TeamsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTeamsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamsUpdateOK creates a TeamsUpdateOK with default headers values
func NewTeamsUpdateOK() *TeamsUpdateOK {
	return &TeamsUpdateOK{}
}

/*TeamsUpdateOK handles this case with default header values.

TeamsUpdateOK teams update o k
*/
type TeamsUpdateOK struct {
	Payload interface{}
}

func (o *TeamsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /v3/customer/core/teams/{id}/][%d] teamsUpdateOK  %+v", 200, o.Payload)
}

func (o *TeamsUpdateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TeamsUpdateBody teams update body
swagger:model TeamsUpdateBody
*/
type TeamsUpdateBody struct {

	// avatar
	Avatar int64 `json:"avatar,omitempty"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// depth
	// Required: true
	Depth *string `json:"depth"`

	// description
	Description string `json:"description,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// org
	// Required: true
	Org *int64 `json:"org"`

	// parent
	Parent int64 `json:"parent,omitempty"`

	// vendor
	Vendor string `json:"vendor,omitempty"`

	// Valid values: SHOW or HIDE
	// Required: true
	Visibility *string `json:"visibility"`
}

// Validate validates this teams update body
func (o *TeamsUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDepth(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOrg(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVisibility(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TeamsUpdateBody) validateDepth(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"depth", "body", o.Depth); err != nil {
		return err
	}

	return nil
}

func (o *TeamsUpdateBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *TeamsUpdateBody) validateOrg(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"org", "body", o.Org); err != nil {
		return err
	}

	return nil
}

func (o *TeamsUpdateBody) validateVisibility(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"visibility", "body", o.Visibility); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TeamsUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TeamsUpdateBody) UnmarshalBinary(b []byte) error {
	var res TeamsUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
