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

// PullRequestsReadReader is a Reader for the PullRequestsRead structure.
type PullRequestsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PullRequestsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPullRequestsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPullRequestsReadOK creates a PullRequestsReadOK with default headers values
func NewPullRequestsReadOK() *PullRequestsReadOK {
	return &PullRequestsReadOK{}
}

/*PullRequestsReadOK handles this case with default header values.

PullRequestsReadOK pull requests read o k
*/
type PullRequestsReadOK struct {
	Payload interface{}
}

func (o *PullRequestsReadOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/pull_requests/{id}/][%d] pullRequestsReadOK  %+v", 200, o.Payload)
}

func (o *PullRequestsReadOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PullRequestsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
