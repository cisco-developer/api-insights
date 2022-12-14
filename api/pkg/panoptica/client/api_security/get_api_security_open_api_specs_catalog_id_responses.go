// Code generated by go-swagger; DO NOT EDIT.

package api_security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"io"
	"github.com/cisco-developer/api-insights/api/pkg/panoptica/models"
)

// GetAPISecurityOpenAPISpecsCatalogIDReader is a Reader for the GetAPISecurityOpenAPISpecsCatalogID structure.
type GetAPISecurityOpenAPISpecsCatalogIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPISecurityOpenAPISpecsCatalogIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPISecurityOpenAPISpecsCatalogIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAPISecurityOpenAPISpecsCatalogIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPISecurityOpenAPISpecsCatalogIDOK creates a GetAPISecurityOpenAPISpecsCatalogIDOK with default headers values
func NewGetAPISecurityOpenAPISpecsCatalogIDOK() *GetAPISecurityOpenAPISpecsCatalogIDOK {
	return &GetAPISecurityOpenAPISpecsCatalogIDOK{}
}

/* GetAPISecurityOpenAPISpecsCatalogIDOK describes a response with status code 200, with default header values.

Success
*/
type GetAPISecurityOpenAPISpecsCatalogIDOK struct {
	Payload *models.OpenAPISpec
}

func (o *GetAPISecurityOpenAPISpecsCatalogIDOK) Error() string {
	return fmt.Sprintf("[GET /apiSecurity/openApiSpecs/{catalogId}][%d] getApiSecurityOpenApiSpecsCatalogIdOK  %+v", 200, o.Payload)
}
func (o *GetAPISecurityOpenAPISpecsCatalogIDOK) GetPayload() *models.OpenAPISpec {
	return o.Payload
}

func (o *GetAPISecurityOpenAPISpecsCatalogIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenAPISpec)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPISecurityOpenAPISpecsCatalogIDUnauthorized creates a GetAPISecurityOpenAPISpecsCatalogIDUnauthorized with default headers values
func NewGetAPISecurityOpenAPISpecsCatalogIDUnauthorized() *GetAPISecurityOpenAPISpecsCatalogIDUnauthorized {
	return &GetAPISecurityOpenAPISpecsCatalogIDUnauthorized{}
}

/* GetAPISecurityOpenAPISpecsCatalogIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type GetAPISecurityOpenAPISpecsCatalogIDUnauthorized struct {
}

func (o *GetAPISecurityOpenAPISpecsCatalogIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /apiSecurity/openApiSpecs/{catalogId}][%d] getApiSecurityOpenApiSpecsCatalogIdUnauthorized ", 401)
}

func (o *GetAPISecurityOpenAPISpecsCatalogIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
