// Code generated by go-swagger; DO NOT EDIT.

package investigation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// ReplaceInstanceReader is a Reader for the ReplaceInstance structure.
type ReplaceInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceInstanceOK creates a ReplaceInstanceOK with default headers values
func NewReplaceInstanceOK() *ReplaceInstanceOK {
	return &ReplaceInstanceOK{}
}

/*ReplaceInstanceOK handles this case with default header values.

successful operation
*/
type ReplaceInstanceOK struct {
	Payload *models.Investigation
}

func (o *ReplaceInstanceOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/investigations/{investigationLocator}][%d] replaceInstanceOK  %+v", 200, o.Payload)
}

func (o *ReplaceInstanceOK) GetPayload() *models.Investigation {
	return o.Payload
}

func (o *ReplaceInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Investigation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
