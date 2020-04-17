// Code generated by go-swagger; DO NOT EDIT.

package mute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// GetMutesReader is a Reader for the GetMutes structure.
type GetMutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMutesOK creates a GetMutesOK with default headers values
func NewGetMutesOK() *GetMutesOK {
	return &GetMutesOK{}
}

/*GetMutesOK handles this case with default header values.

successful operation
*/
type GetMutesOK struct {
	Payload *models.Mutes
}

func (o *GetMutesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/mutes][%d] getMutesOK  %+v", 200, o.Payload)
}

func (o *GetMutesOK) GetPayload() *models.Mutes {
	return o.Payload
}

func (o *GetMutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Mutes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
