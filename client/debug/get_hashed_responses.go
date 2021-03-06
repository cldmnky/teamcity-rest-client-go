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

// GetHashedReader is a Reader for the GetHashed structure.
type GetHashedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHashedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHashedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetHashedOK creates a GetHashedOK with default headers values
func NewGetHashedOK() *GetHashedOK {
	return &GetHashedOK{}
}

/*GetHashedOK handles this case with default header values.

successful operation
*/
type GetHashedOK struct {
	Payload string
}

func (o *GetHashedOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/debug/values/transform/{method}][%d] getHashedOK  %+v", 200, o.Payload)
}

func (o *GetHashedOK) GetPayload() string {
	return o.Payload
}

func (o *GetHashedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
