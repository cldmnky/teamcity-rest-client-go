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

// GetArtifactDepSettingReader is a Reader for the GetArtifactDepSetting structure.
type GetArtifactDepSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArtifactDepSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArtifactDepSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetArtifactDepSettingOK creates a GetArtifactDepSettingOK with default headers values
func NewGetArtifactDepSettingOK() *GetArtifactDepSettingOK {
	return &GetArtifactDepSettingOK{}
}

/*GetArtifactDepSettingOK handles this case with default header values.

successful operation
*/
type GetArtifactDepSettingOK struct {
	Payload string
}

func (o *GetArtifactDepSettingOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/artifact-dependencies/{artifactDepLocator}/{fieldName}][%d] getArtifactDepSettingOK  %+v", 200, o.Payload)
}

func (o *GetArtifactDepSettingOK) GetPayload() string {
	return o.Payload
}

func (o *GetArtifactDepSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
