// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// GetPinDataReader is a Reader for the GetPinData structure.
type GetPinDataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPinDataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPinDataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPinDataOK creates a GetPinDataOK with default headers values
func NewGetPinDataOK() *GetPinDataOK {
	return &GetPinDataOK{}
}

/*GetPinDataOK handles this case with default header values.

successful operation
*/
type GetPinDataOK struct {
	Payload *models.PinInfo
}

func (o *GetPinDataOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/builds/{buildLocator}/pinInfo][%d] getPinDataOK  %+v", 200, o.Payload)
}

func (o *GetPinDataOK) GetPayload() *models.PinInfo {
	return o.Payload
}

func (o *GetPinDataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PinInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}