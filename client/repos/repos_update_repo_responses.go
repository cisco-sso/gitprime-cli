// Code generated by go-swagger; DO NOT EDIT.

package repos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ReposUpdateRepoReader is a Reader for the ReposUpdateRepo structure.
type ReposUpdateRepoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReposUpdateRepoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewReposUpdateRepoCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReposUpdateRepoCreated creates a ReposUpdateRepoCreated with default headers values
func NewReposUpdateRepoCreated() *ReposUpdateRepoCreated {
	return &ReposUpdateRepoCreated{}
}

/*ReposUpdateRepoCreated handles this case with default header values.

ReposUpdateRepoCreated repos update repo created
*/
type ReposUpdateRepoCreated struct {
	Payload interface{}
}

func (o *ReposUpdateRepoCreated) Error() string {
	return fmt.Sprintf("[POST /v3/customer/core/repos/{project}/update_repo/][%d] reposUpdateRepoCreated  %+v", 201, o.Payload)
}

func (o *ReposUpdateRepoCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *ReposUpdateRepoCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}