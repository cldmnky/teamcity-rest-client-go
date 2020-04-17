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

// EmptyTaskReader is a Reader for the EmptyTask structure.
type EmptyTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmptyTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmptyTaskOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmptyTaskOK creates a EmptyTaskOK with default headers values
func NewEmptyTaskOK() *EmptyTaskOK {
	return &EmptyTaskOK{}
}

/*EmptyTaskOK handles this case with default header values.

successful operation
*/
type EmptyTaskOK struct {
	Payload string
}

func (o *EmptyTaskOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/debug/emptyTask][%d] emptyTaskOK  %+v", 200, o.Payload)
}

func (o *EmptyTaskOK) GetPayload() string {
	return o.Payload
}

func (o *EmptyTaskOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
