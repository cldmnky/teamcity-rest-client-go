// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// ServeGroupsReader is a Reader for the ServeGroups structure.
type ServeGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServeGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeGroupsOK creates a ServeGroupsOK with default headers values
func NewServeGroupsOK() *ServeGroupsOK {
	return &ServeGroupsOK{}
}

/*ServeGroupsOK handles this case with default header values.

successful operation
*/
type ServeGroupsOK struct {
	Payload *models.Groups
}

func (o *ServeGroupsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/userGroups][%d] serveGroupsOK  %+v", 200, o.Payload)
}

func (o *ServeGroupsOK) GetPayload() *models.Groups {
	return o.Payload
}

func (o *ServeGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Groups)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
