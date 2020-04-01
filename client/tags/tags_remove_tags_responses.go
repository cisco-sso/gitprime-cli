// Code generated by go-swagger; DO NOT EDIT.

package tags

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

// TagsRemoveTagsReader is a Reader for the TagsRemoveTags structure.
type TagsRemoveTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TagsRemoveTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTagsRemoveTagsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTagsRemoveTagsCreated creates a TagsRemoveTagsCreated with default headers values
func NewTagsRemoveTagsCreated() *TagsRemoveTagsCreated {
	return &TagsRemoveTagsCreated{}
}

/*TagsRemoveTagsCreated handles this case with default header values.

TagsRemoveTagsCreated tags remove tags created
*/
type TagsRemoveTagsCreated struct {
	Payload interface{}
}

func (o *TagsRemoveTagsCreated) Error() string {
	return fmt.Sprintf("[POST /v3/customer/core/tags/remove_tags/][%d] tagsRemoveTagsCreated  %+v", 201, o.Payload)
}

func (o *TagsRemoveTagsCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *TagsRemoveTagsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TagsRemoveTagsBody tags remove tags body
swagger:model TagsRemoveTagsBody
*/
type TagsRemoveTagsBody struct {

	// remove unused
	RemoveUnused bool `json:"remove_unused,omitempty"`

	// repo ids
	// Required: true
	RepoIds []string `json:"repo_ids"`

	// tag ids
	// Required: true
	TagIds []string `json:"tag_ids"`
}

// Validate validates this tags remove tags body
func (o *TagsRemoveTagsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRepoIds(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTagIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TagsRemoveTagsBody) validateRepoIds(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"repo_ids", "body", o.RepoIds); err != nil {
		return err
	}

	return nil
}

func (o *TagsRemoveTagsBody) validateTagIds(formats strfmt.Registry) error {

	if err := validate.Required("data"+"."+"tag_ids", "body", o.TagIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *TagsRemoveTagsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TagsRemoveTagsBody) UnmarshalBinary(b []byte) error {
	var res TagsRemoveTagsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
