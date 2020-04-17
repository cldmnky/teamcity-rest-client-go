// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetResolvedParameterReader is a Reader for the GetResolvedParameter structure.
type GetResolvedParameterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResolvedParameterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResolvedParameterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetResolvedParameterOK creates a GetResolvedParameterOK with default headers values
func NewGetResolvedParameterOK() *GetResolvedParameterOK {
	return &GetResolvedParameterOK{}
}

/*GetResolvedParameterOK handles this case with default header values.

successful operation
*/
type GetResolvedParameterOK struct {
	Payload string
}

func (o *GetResolvedParameterOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/builds/{buildLocator}/resolved/{value}][%d] getResolvedParameterOK  %+v", 200, o.Payload)
}

func (o *GetResolvedParameterOK) GetPayload() string {
	return o.Payload
}

func (o *GetResolvedParameterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}