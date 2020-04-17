// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// GetParentGroupsReader is a Reader for the GetParentGroups structure.
type GetParentGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetParentGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetParentGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetParentGroupsOK creates a GetParentGroupsOK with default headers values
func NewGetParentGroupsOK() *GetParentGroupsOK {
	return &GetParentGroupsOK{}
}

/*GetParentGroupsOK handles this case with default header values.

successful operation
*/
type GetParentGroupsOK struct {
	Payload *models.Groups
}

func (o *GetParentGroupsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/userGroups/{groupLocator}/parent-groups][%d] getParentGroupsOK  %+v", 200, o.Payload)
}

func (o *GetParentGroupsOK) GetPayload() *models.Groups {
	return o.Payload
}

func (o *GetParentGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Groups)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
