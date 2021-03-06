// Code generated by go-swagger; DO NOT EDIT.

package debug

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetRequestDetailsReader is a Reader for the GetRequestDetails structure.
type GetRequestDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRequestDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRequestDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRequestDetailsOK creates a GetRequestDetailsOK with default headers values
func NewGetRequestDetailsOK() *GetRequestDetailsOK {
	return &GetRequestDetailsOK{}
}

/*GetRequestDetailsOK handles this case with default header values.

successful operation
*/
type GetRequestDetailsOK struct {
	Payload string
}

func (o *GetRequestDetailsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/debug/currentRequest/details{extra}][%d] getRequestDetailsOK  %+v", 200, o.Payload)
}

func (o *GetRequestDetailsOK) GetPayload() string {
	return o.Payload
}

func (o *GetRequestDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
