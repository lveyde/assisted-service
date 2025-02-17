// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// UploadHostLogsReader is a Reader for the UploadHostLogs structure.
type UploadHostLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadHostLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUploadHostLogsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadHostLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUploadHostLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUploadHostLogsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUploadHostLogsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUploadHostLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewUploadHostLogsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadHostLogsNoContent creates a UploadHostLogsNoContent with default headers values
func NewUploadHostLogsNoContent() *UploadHostLogsNoContent {
	return &UploadHostLogsNoContent{}
}

/* UploadHostLogsNoContent describes a response with status code 204, with default header values.

Success.
*/
type UploadHostLogsNoContent struct {
}

func (o *UploadHostLogsNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/logs][%d] uploadHostLogsNoContent ", 204)
}

func (o *UploadHostLogsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadHostLogsBadRequest creates a UploadHostLogsBadRequest with default headers values
func NewUploadHostLogsBadRequest() *UploadHostLogsBadRequest {
	return &UploadHostLogsBadRequest{}
}

/* UploadHostLogsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UploadHostLogsBadRequest struct {
	Payload *models.InfraError
}

func (o *UploadHostLogsBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/logs][%d] uploadHostLogsBadRequest  %+v", 400, o.Payload)
}
func (o *UploadHostLogsBadRequest) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UploadHostLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadHostLogsUnauthorized creates a UploadHostLogsUnauthorized with default headers values
func NewUploadHostLogsUnauthorized() *UploadHostLogsUnauthorized {
	return &UploadHostLogsUnauthorized{}
}

/* UploadHostLogsUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type UploadHostLogsUnauthorized struct {
	Payload *models.InfraError
}

func (o *UploadHostLogsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/logs][%d] uploadHostLogsUnauthorized  %+v", 401, o.Payload)
}
func (o *UploadHostLogsUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UploadHostLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadHostLogsForbidden creates a UploadHostLogsForbidden with default headers values
func NewUploadHostLogsForbidden() *UploadHostLogsForbidden {
	return &UploadHostLogsForbidden{}
}

/* UploadHostLogsForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type UploadHostLogsForbidden struct {
	Payload *models.InfraError
}

func (o *UploadHostLogsForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/logs][%d] uploadHostLogsForbidden  %+v", 403, o.Payload)
}
func (o *UploadHostLogsForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UploadHostLogsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadHostLogsNotFound creates a UploadHostLogsNotFound with default headers values
func NewUploadHostLogsNotFound() *UploadHostLogsNotFound {
	return &UploadHostLogsNotFound{}
}

/* UploadHostLogsNotFound describes a response with status code 404, with default header values.

Error.
*/
type UploadHostLogsNotFound struct {
	Payload *models.Error
}

func (o *UploadHostLogsNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/logs][%d] uploadHostLogsNotFound  %+v", 404, o.Payload)
}
func (o *UploadHostLogsNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UploadHostLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadHostLogsInternalServerError creates a UploadHostLogsInternalServerError with default headers values
func NewUploadHostLogsInternalServerError() *UploadHostLogsInternalServerError {
	return &UploadHostLogsInternalServerError{}
}

/* UploadHostLogsInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type UploadHostLogsInternalServerError struct {
	Payload *models.Error
}

func (o *UploadHostLogsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/logs][%d] uploadHostLogsInternalServerError  %+v", 500, o.Payload)
}
func (o *UploadHostLogsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UploadHostLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadHostLogsServiceUnavailable creates a UploadHostLogsServiceUnavailable with default headers values
func NewUploadHostLogsServiceUnavailable() *UploadHostLogsServiceUnavailable {
	return &UploadHostLogsServiceUnavailable{}
}

/* UploadHostLogsServiceUnavailable describes a response with status code 503, with default header values.

Unavailable.
*/
type UploadHostLogsServiceUnavailable struct {
	Payload *models.Error
}

func (o *UploadHostLogsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{cluster_id}/hosts/{host_id}/logs][%d] uploadHostLogsServiceUnavailable  %+v", 503, o.Payload)
}
func (o *UploadHostLogsServiceUnavailable) GetPayload() *models.Error {
	return o.Payload
}

func (o *UploadHostLogsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
