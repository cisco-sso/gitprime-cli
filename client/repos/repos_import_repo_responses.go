// Code generated by go-swagger; DO NOT EDIT.

package repos

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

// ReposImportRepoReader is a Reader for the ReposImportRepo structure.
type ReposImportRepoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReposImportRepoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewReposImportRepoCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReposImportRepoCreated creates a ReposImportRepoCreated with default headers values
func NewReposImportRepoCreated() *ReposImportRepoCreated {
	return &ReposImportRepoCreated{}
}

/*ReposImportRepoCreated handles this case with default header values.

ReposImportRepoCreated repos import repo created
*/
type ReposImportRepoCreated struct {
	Payload interface{}
}

func (o *ReposImportRepoCreated) Error() string {
	return fmt.Sprintf("[POST /v3/customer/core/repos/import_repo/][%d] reposImportRepoCreated  %+v", 201, o.Payload)
}

func (o *ReposImportRepoCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *ReposImportRepoCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ReposImportRepoBody repos import repo body
swagger:model ReposImportRepoBody
*/
type ReposImportRepoBody struct {

	// available repo id
	// Required: true
	AvailableRepoID *string `json:"available_repo_id"`
}

// Validate validates this repos import repo body
func (o *ReposImportRepoBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvailableRepoID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ReposImportRepoBody) validateAvailableRepoID(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"available_repo_id", "body", o.AvailableRepoID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ReposImportRepoBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ReposImportRepoBody) UnmarshalBinary(b []byte) error {
	var res ReposImportRepoBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
