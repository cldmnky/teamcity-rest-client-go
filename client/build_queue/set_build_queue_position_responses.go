// Code generated by go-swagger; DO NOT EDIT.

package build_queue

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// SetBuildQueuePositionReader is a Reader for the SetBuildQueuePosition structure.
type SetBuildQueuePositionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetBuildQueuePositionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetBuildQueuePositionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetBuildQueuePositionOK creates a SetBuildQueuePositionOK with default headers values
func NewSetBuildQueuePositionOK() *SetBuildQueuePositionOK {
	return &SetBuildQueuePositionOK{}
}

/*SetBuildQueuePositionOK handles this case with default header values.

successful operation
*/
type SetBuildQueuePositionOK struct {
	Payload *models.Build
}

func (o *SetBuildQueuePositionOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildQueue/order/{queuePosition}][%d] setBuildQueuePositionOK  %+v", 200, o.Payload)
}

func (o *SetBuildQueuePositionOK) GetPayload() *models.Build {
	return o.Payload
}

func (o *SetBuildQueuePositionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Build)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
