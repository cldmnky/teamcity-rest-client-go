// Code generated by go-swagger; DO NOT EDIT.

package debug

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// GetCachedBuildPromotionsStatsReader is a Reader for the GetCachedBuildPromotionsStats structure.
type GetCachedBuildPromotionsStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCachedBuildPromotionsStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCachedBuildPromotionsStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCachedBuildPromotionsStatsOK creates a GetCachedBuildPromotionsStatsOK with default headers values
func NewGetCachedBuildPromotionsStatsOK() *GetCachedBuildPromotionsStatsOK {
	return &GetCachedBuildPromotionsStatsOK{}
}

/*GetCachedBuildPromotionsStatsOK handles this case with default header values.

successful operation
*/
type GetCachedBuildPromotionsStatsOK struct {
	Payload *models.Properties
}

func (o *GetCachedBuildPromotionsStatsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/debug/caches/buildPromotions/stats][%d] getCachedBuildPromotionsStatsOK  %+v", 200, o.Payload)
}

func (o *GetCachedBuildPromotionsStatsOK) GetPayload() *models.Properties {
	return o.Payload
}

func (o *GetCachedBuildPromotionsStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Properties)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
