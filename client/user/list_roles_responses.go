// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// ListRolesReader is a Reader for the ListRoles structure.
type ListRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListRolesOK creates a ListRolesOK with default headers values
func NewListRolesOK() *ListRolesOK {
	return &ListRolesOK{}
}

/*ListRolesOK handles this case with default header values.

successful operation
*/
type ListRolesOK struct {
	Payload *models.Roles
}

func (o *ListRolesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/users/{userLocator}/roles][%d] listRolesOK  %+v", 200, o.Payload)
}

func (o *ListRolesOK) GetPayload() *models.Roles {
	return o.Payload
}

func (o *ListRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Roles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
