// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// GetParameterTypeReader is a Reader for the GetParameterType structure.
type GetParameterTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetParameterTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetParameterTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetParameterTypeOK creates a GetParameterTypeOK with default headers values
func NewGetParameterTypeOK() *GetParameterTypeOK {
	return &GetParameterTypeOK{}
}

/*GetParameterTypeOK handles this case with default header values.

successful operation
*/
type GetParameterTypeOK struct {
	Payload *models.Type
}

func (o *GetParameterTypeOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/projects/{projectLocator}/parameters/{name}/type][%d] getParameterTypeOK  %+v", 200, o.Payload)
}

func (o *GetParameterTypeOK) GetPayload() *models.Type {
	return o.Payload
}

func (o *GetParameterTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Type)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
