// Code generated by go-swagger; DO NOT EDIT.

package vcs_root

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// ServeRootInstanceReader is a Reader for the ServeRootInstance structure.
type ServeRootInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeRootInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServeRootInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeRootInstanceOK creates a ServeRootInstanceOK with default headers values
func NewServeRootInstanceOK() *ServeRootInstanceOK {
	return &ServeRootInstanceOK{}
}

/*ServeRootInstanceOK handles this case with default header values.

successful operation
*/
type ServeRootInstanceOK struct {
	Payload *models.VcsRootInstance
}

func (o *ServeRootInstanceOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/vcs-roots/{vcsRootLocator}/instances/{vcsRootInstanceLocator}][%d] serveRootInstanceOK  %+v", 200, o.Payload)
}

func (o *ServeRootInstanceOK) GetPayload() *models.VcsRootInstance {
	return o.Payload
}

func (o *ServeRootInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VcsRootInstance)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
