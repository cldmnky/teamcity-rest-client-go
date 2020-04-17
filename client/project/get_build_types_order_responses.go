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

// GetBuildTypesOrderReader is a Reader for the GetBuildTypesOrder structure.
type GetBuildTypesOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBuildTypesOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBuildTypesOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBuildTypesOrderOK creates a GetBuildTypesOrderOK with default headers values
func NewGetBuildTypesOrderOK() *GetBuildTypesOrderOK {
	return &GetBuildTypesOrderOK{}
}

/*GetBuildTypesOrderOK handles this case with default header values.

successful operation
*/
type GetBuildTypesOrderOK struct {
	Payload *models.BuildTypes
}

func (o *GetBuildTypesOrderOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/projects/{projectLocator}/order/buildTypes][%d] getBuildTypesOrderOK  %+v", 200, o.Payload)
}

func (o *GetBuildTypesOrderOK) GetPayload() *models.BuildTypes {
	return o.Payload
}

func (o *GetBuildTypesOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildTypes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
