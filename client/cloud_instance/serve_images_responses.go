// Code generated by go-swagger; DO NOT EDIT.

package cloud_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// ServeImagesReader is a Reader for the ServeImages structure.
type ServeImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServeImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeImagesOK creates a ServeImagesOK with default headers values
func NewServeImagesOK() *ServeImagesOK {
	return &ServeImagesOK{}
}

/*ServeImagesOK handles this case with default header values.

successful operation
*/
type ServeImagesOK struct {
	Payload *models.CloudImages
}

func (o *ServeImagesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/cloud/images][%d] serveImagesOK  %+v", 200, o.Payload)
}

func (o *ServeImagesOK) GetPayload() *models.CloudImages {
	return o.Payload
}

func (o *ServeImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudImages)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
