// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// GetFeatureReader is a Reader for the GetFeature structure.
type GetFeatureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFeatureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFeatureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFeatureOK creates a GetFeatureOK with default headers values
func NewGetFeatureOK() *GetFeatureOK {
	return &GetFeatureOK{}
}

/*GetFeatureOK handles this case with default header values.

successful operation
*/
type GetFeatureOK struct {
	Payload *models.Feature
}

func (o *GetFeatureOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/buildTypes/{btLocator}/features/{featureId}][%d] getFeatureOK  %+v", 200, o.Payload)
}

func (o *GetFeatureOK) GetPayload() *models.Feature {
	return o.Payload
}

func (o *GetFeatureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Feature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
