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

// GetStepSettingReader is a Reader for the GetStepSetting structure.
type GetStepSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStepSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStepSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStepSettingOK creates a GetStepSettingOK with default headers values
func NewGetStepSettingOK() *GetStepSettingOK {
	return &GetStepSettingOK{}
}

/*GetStepSettingOK handles this case with default header values.

successful operation
*/
type GetStepSettingOK struct {
	Payload string
}

func (o *GetStepSettingOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/steps/{stepId}/{fieldName}][%d] getStepSettingOK  %+v", 200, o.Payload)
}

func (o *GetStepSettingOK) GetPayload() string {
	return o.Payload
}

func (o *GetStepSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
