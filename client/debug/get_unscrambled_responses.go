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

// GetUnscrambledReader is a Reader for the GetUnscrambled structure.
type GetUnscrambledReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUnscrambledReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUnscrambledOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUnscrambledOK creates a GetUnscrambledOK with default headers values
func NewGetUnscrambledOK() *GetUnscrambledOK {
	return &GetUnscrambledOK{}
}

/*GetUnscrambledOK handles this case with default header values.

successful operation
*/
type GetUnscrambledOK struct {
	Payload string
}

func (o *GetUnscrambledOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/debug/values/password/unscrambled][%d] getUnscrambledOK  %+v", 200, o.Payload)
}

func (o *GetUnscrambledOK) GetPayload() string {
	return o.Payload
}

func (o *GetUnscrambledOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
