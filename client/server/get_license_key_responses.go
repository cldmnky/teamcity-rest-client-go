// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// GetLicenseKeyReader is a Reader for the GetLicenseKey structure.
type GetLicenseKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLicenseKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLicenseKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLicenseKeyOK creates a GetLicenseKeyOK with default headers values
func NewGetLicenseKeyOK() *GetLicenseKeyOK {
	return &GetLicenseKeyOK{}
}

/*GetLicenseKeyOK handles this case with default header values.

successful operation
*/
type GetLicenseKeyOK struct {
	Payload *models.LicenseKey
}

func (o *GetLicenseKeyOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/server/licensingData/licenseKeys/{licenseKey}][%d] getLicenseKeyOK  %+v", 200, o.Payload)
}

func (o *GetLicenseKeyOK) GetPayload() *models.LicenseKey {
	return o.Payload
}

func (o *GetLicenseKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseKey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}