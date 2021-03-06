// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UsersPartialUpdateReader is a Reader for the UsersPartialUpdate structure.
type UsersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersPartialUpdateOK creates a UsersPartialUpdateOK with default headers values
func NewUsersPartialUpdateOK() *UsersPartialUpdateOK {
	return &UsersPartialUpdateOK{}
}

/*UsersPartialUpdateOK handles this case with default header values.

UsersPartialUpdateOK users partial update o k
*/
type UsersPartialUpdateOK struct {
	Payload interface{}
}

func (o *UsersPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v3/customer/core/users/{id}/][%d] usersPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersPartialUpdateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UsersPartialUpdateBody users partial update body
swagger:model UsersPartialUpdateBody
*/
type UsersPartialUpdateBody struct {

	// email
	Email string `json:"email,omitempty"`

	// first activity at
	FirstActivityAt string `json:"first_activity_at,omitempty"`

	// hidden from reports
	HiddenFromReports bool `json:"hidden_from_reports,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// team ids
	TeamIds []string `json:"team_ids"`
}

// Validate validates this users partial update body
func (o *UsersPartialUpdateBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UsersPartialUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersPartialUpdateBody) UnmarshalBinary(b []byte) error {
	var res UsersPartialUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
