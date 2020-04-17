// Code generated by go-swagger; DO NOT EDIT.

package agent

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// GetAuthorizedInfoReader is a Reader for the GetAuthorizedInfo structure.
type GetAuthorizedInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizedInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizedInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAuthorizedInfoOK creates a GetAuthorizedInfoOK with default headers values
func NewGetAuthorizedInfoOK() *GetAuthorizedInfoOK {
	return &GetAuthorizedInfoOK{}
}

/*GetAuthorizedInfoOK handles this case with default header values.

successful operation
*/
type GetAuthorizedInfoOK struct {
	Payload *models.AuthorizedInfo
}

func (o *GetAuthorizedInfoOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/agents/{agentLocator}/authorizedInfo][%d] getAuthorizedInfoOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizedInfoOK) GetPayload() *models.AuthorizedInfo {
	return o.Payload
}

func (o *GetAuthorizedInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizedInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
