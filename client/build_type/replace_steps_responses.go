// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// ReplaceStepsReader is a Reader for the ReplaceSteps structure.
type ReplaceStepsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceStepsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceStepsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceStepsOK creates a ReplaceStepsOK with default headers values
func NewReplaceStepsOK() *ReplaceStepsOK {
	return &ReplaceStepsOK{}
}

/*ReplaceStepsOK handles this case with default header values.

successful operation
*/
type ReplaceStepsOK struct {
	Payload *models.Steps
}

func (o *ReplaceStepsOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/steps][%d] replaceStepsOK  %+v", 200, o.Payload)
}

func (o *ReplaceStepsOK) GetPayload() *models.Steps {
	return o.Payload
}

func (o *ReplaceStepsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Steps)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
