// Code generated by go-swagger; DO NOT EDIT.

package users

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

// UsersLoginCreateReader is a Reader for the UsersLoginCreate structure.
type UsersLoginCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersLoginCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersLoginCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersLoginCreateCreated creates a UsersLoginCreateCreated with default headers values
func NewUsersLoginCreateCreated() *UsersLoginCreateCreated {
	return &UsersLoginCreateCreated{}
}

/*UsersLoginCreateCreated handles this case with default header values.

UsersLoginCreateCreated users login create created
*/
type UsersLoginCreateCreated struct {
	Payload interface{}
}

func (o *UsersLoginCreateCreated) Error() string {
	return fmt.Sprintf("[POST /v3/customer/core/users/{id}/login/][%d] usersLoginCreateCreated  %+v", 201, o.Payload)
}

func (o *UsersLoginCreateCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersLoginCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UsersLoginCreateBody users login create body
swagger:model UsersLoginCreateBody
*/
type UsersLoginCreateBody struct {

	// access ids
	AccessIds []string `json:"access_ids"`

	// login email
	// Required: true
	LoginEmail *string `json:"login_email"`

	// Designates whether this user should be treated as active. Unselect this instead of deleting accounts.
	LoginEnabled bool `json:"login_enabled,omitempty"`

	// login password
	// Required: true
	LoginPassword *string `json:"login_password"`

	// login preferred name
	// Required: true
	LoginPreferredName *string `json:"login_preferred_name"`

	// needs password reset
	NeedsPasswordReset bool `json:"needs_password_reset,omitempty"`

	// role ids
	RoleIds []string `json:"role_ids"`

	// team visibility
	TeamVisibility string `json:"team_visibility,omitempty"`

	// user visibility
	UserVisibility string `json:"user_visibility,omitempty"`
}

// Validate validates this users login create body
func (o *UsersLoginCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLoginEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLoginPassword(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLoginPreferredName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UsersLoginCreateBody) validateLoginEmail(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"login_email", "body", o.LoginEmail); err != nil {
		return err
	}

	return nil
}

func (o *UsersLoginCreateBody) validateLoginPassword(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"login_password", "body", o.LoginPassword); err != nil {
		return err
	}

	return nil
}

func (o *UsersLoginCreateBody) validateLoginPreferredName(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"login_preferred_name", "body", o.LoginPreferredName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UsersLoginCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UsersLoginCreateBody) UnmarshalBinary(b []byte) error {
	var res UsersLoginCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
