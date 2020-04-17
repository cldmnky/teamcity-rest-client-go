// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// AddTagsMultipleReader is a Reader for the AddTagsMultiple structure.
type AddTagsMultipleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddTagsMultipleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddTagsMultipleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddTagsMultipleOK creates a AddTagsMultipleOK with default headers values
func NewAddTagsMultipleOK() *AddTagsMultipleOK {
	return &AddTagsMultipleOK{}
}

/*AddTagsMultipleOK handles this case with default header values.

successful operation
*/
type AddTagsMultipleOK struct {
	Payload *models.MultipleOperationResult
}

func (o *AddTagsMultipleOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/builds/multiple/{buildLocator}/tags][%d] addTagsMultipleOK  %+v", 200, o.Payload)
}

func (o *AddTagsMultipleOK) GetPayload() *models.MultipleOperationResult {
	return o.Payload
}

func (o *AddTagsMultipleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MultipleOperationResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
