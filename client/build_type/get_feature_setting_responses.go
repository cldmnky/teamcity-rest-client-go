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

// GetFeatureSettingReader is a Reader for the GetFeatureSetting structure.
type GetFeatureSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFeatureSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFeatureSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFeatureSettingOK creates a GetFeatureSettingOK with default headers values
func NewGetFeatureSettingOK() *GetFeatureSettingOK {
	return &GetFeatureSettingOK{}
}

/*GetFeatureSettingOK handles this case with default header values.

successful operation
*/
type GetFeatureSettingOK struct {
	Payload string
}

func (o *GetFeatureSettingOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/features/{featureId}/{name}][%d] getFeatureSettingOK  %+v", 200, o.Payload)
}

func (o *GetFeatureSettingOK) GetPayload() string {
	return o.Payload
}

func (o *GetFeatureSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
