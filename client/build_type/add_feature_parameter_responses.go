// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddFeatureParameterReader is a Reader for the AddFeatureParameter structure.
type AddFeatureParameterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddFeatureParameterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddFeatureParameterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddFeatureParameterOK creates a AddFeatureParameterOK with default headers values
func NewAddFeatureParameterOK() *AddFeatureParameterOK {
	return &AddFeatureParameterOK{}
}

/*AddFeatureParameterOK handles this case with default header values.

successful operation
*/
type AddFeatureParameterOK struct {
	Payload string
}

func (o *AddFeatureParameterOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/features/{featureId}/parameters/{parameterName}][%d] addFeatureParameterOK  %+v", 200, o.Payload)
}

func (o *AddFeatureParameterOK) GetPayload() string {
	return o.Payload
}

func (o *AddFeatureParameterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}