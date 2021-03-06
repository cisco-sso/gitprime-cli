// Code generated by go-swagger; DO NOT EDIT.

package org_membership

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// OrgMembershipListReader is a Reader for the OrgMembershipList structure.
type OrgMembershipListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrgMembershipListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOrgMembershipListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrgMembershipListOK creates a OrgMembershipListOK with default headers values
func NewOrgMembershipListOK() *OrgMembershipListOK {
	return &OrgMembershipListOK{}
}

/*OrgMembershipListOK handles this case with default header values.

OrgMembershipListOK org membership list o k
*/
type OrgMembershipListOK struct {
	Payload interface{}
}

func (o *OrgMembershipListOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/org_membership/][%d] orgMembershipListOK  %+v", 200, o.Payload)
}

func (o *OrgMembershipListOK) GetPayload() interface{} {
	return o.Payload
}

func (o *OrgMembershipListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
