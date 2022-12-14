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

// DeleteAPISecurityOpenAPISpecsCatalogIDReader is a Reader for the DeleteAPISecurityOpenAPISpecsCatalogID structure.
type DeleteAPISecurityOpenAPISpecsCatalogIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPISecurityOpenAPISpecsCatalogIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteAPISecurityOpenAPISpecsCatalogIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteAPISecurityOpenAPISpecsCatalogIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteAPISecurityOpenAPISpecsCatalogIDNoContent creates a DeleteAPISecurityOpenAPISpecsCatalogIDNoContent with default headers values
func NewDeleteAPISecurityOpenAPISpecsCatalogIDNoContent() *DeleteAPISecurityOpenAPISpecsCatalogIDNoContent {
	return &DeleteAPISecurityOpenAPISpecsCatalogIDNoContent{}
}

/* DeleteAPISecurityOpenAPISpecsCatalogIDNoContent describes a response with status code 204, with default header values.

Success
*/
type DeleteAPISecurityOpenAPISpecsCatalogIDNoContent struct {
	Payload models.OpenAPISpecScoreStatus
}

func (o *DeleteAPISecurityOpenAPISpecsCatalogIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /apiSecurity/openApiSpecs/{catalogId}][%d] deleteApiSecurityOpenApiSpecsCatalogIdNoContent  %+v", 204, o.Payload)
}
func (o *DeleteAPISecurityOpenAPISpecsCatalogIDNoContent) GetPayload() models.OpenAPISpecScoreStatus {
	return o.Payload
}

func (o *DeleteAPISecurityOpenAPISpecsCatalogIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAPISecurityOpenAPISpecsCatalogIDUnauthorized creates a DeleteAPISecurityOpenAPISpecsCatalogIDUnauthorized with default headers values
func NewDeleteAPISecurityOpenAPISpecsCatalogIDUnauthorized() *DeleteAPISecurityOpenAPISpecsCatalogIDUnauthorized {
	return &DeleteAPISecurityOpenAPISpecsCatalogIDUnauthorized{}
}

/* DeleteAPISecurityOpenAPISpecsCatalogIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type DeleteAPISecurityOpenAPISpecsCatalogIDUnauthorized struct {
}

func (o *DeleteAPISecurityOpenAPISpecsCatalogIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /apiSecurity/openApiSpecs/{catalogId}][%d] deleteApiSecurityOpenApiSpecsCatalogIdUnauthorized ", 401)
}

func (o *DeleteAPISecurityOpenAPISpecsCatalogIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
