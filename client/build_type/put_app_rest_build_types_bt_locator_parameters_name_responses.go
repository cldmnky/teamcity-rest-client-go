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

// PutAppRestBuildTypesBtLocatorParametersNameReader is a Reader for the PutAppRestBuildTypesBtLocatorParametersName structure.
type PutAppRestBuildTypesBtLocatorParametersNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAppRestBuildTypesBtLocatorParametersNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutAppRestBuildTypesBtLocatorParametersNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutAppRestBuildTypesBtLocatorParametersNameOK creates a PutAppRestBuildTypesBtLocatorParametersNameOK with default headers values
func NewPutAppRestBuildTypesBtLocatorParametersNameOK() *PutAppRestBuildTypesBtLocatorParametersNameOK {
	return &PutAppRestBuildTypesBtLocatorParametersNameOK{}
}

/*PutAppRestBuildTypesBtLocatorParametersNameOK handles this case with default header values.

successful operation
*/
type PutAppRestBuildTypesBtLocatorParametersNameOK struct {
	Payload *models.Property
}

func (o *PutAppRestBuildTypesBtLocatorParametersNameOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/parameters/{name}][%d] putAppRestBuildTypesBtLocatorParametersNameOK  %+v", 200, o.Payload)
}

func (o *PutAppRestBuildTypesBtLocatorParametersNameOK) GetPayload() *models.Property {
	return o.Payload
}

func (o *PutAppRestBuildTypesBtLocatorParametersNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Property)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
