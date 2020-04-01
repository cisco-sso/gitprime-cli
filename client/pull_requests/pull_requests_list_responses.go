// Code generated by go-swagger; DO NOT EDIT.

package pull_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PullRequestsListReader is a Reader for the PullRequestsList structure.
type PullRequestsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PullRequestsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPullRequestsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPullRequestsListOK creates a PullRequestsListOK with default headers values
func NewPullRequestsListOK() *PullRequestsListOK {
	return &PullRequestsListOK{}
}

/*PullRequestsListOK handles this case with default header values.

PullRequestsListOK pull requests list o k
*/
type PullRequestsListOK struct {
	Payload interface{}
}

func (o *PullRequestsListOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/pull_requests/][%d] pullRequestsListOK  %+v", 200, o.Payload)
}

func (o *PullRequestsListOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PullRequestsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
