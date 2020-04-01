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

// ReposDeleteReader is a Reader for the ReposDelete structure.
type ReposDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReposDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewReposDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReposDeleteNoContent creates a ReposDeleteNoContent with default headers values
func NewReposDeleteNoContent() *ReposDeleteNoContent {
	return &ReposDeleteNoContent{}
}

/*ReposDeleteNoContent handles this case with default header values.

ReposDeleteNoContent repos delete no content
*/
type ReposDeleteNoContent struct {
	Payload interface{}
}

func (o *ReposDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v3/customer/core/repos/{project}/][%d] reposDeleteNoContent  %+v", 204, o.Payload)
}

func (o *ReposDeleteNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ReposDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}