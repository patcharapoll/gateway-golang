package utils

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
)

// ParseErrorResponse ...
func ParseErrorResponse(err error) *gqlerror.Error {
	errStatus, _ := status.FromError(err)

	// Initial error response
	errorResponse := &gqlerror.Error{
		Message:    errStatus.Message(),
		Extensions: make(map[string]interface{}),
	}

	// Extract the error status detail and insert to extensions.
	errorDetail := errStatus.Details()

	// Retrive error code from top of the hierarchy
	errorResponse.Extensions["code"] = ToUpperSnakeCase(errStatus.Code().String())
	if len(errorDetail) > 0 {
		for _, detail := range errorDetail {
			switch t := detail.(type) {

			case *errdetails.ErrorInfo:
				// Override the error code when ErrorInfo available
				errorResponse.Extensions["code"] = ToUpperSnakeCase(t.GetReason())

				metadata := t.GetMetadata()
				for k, v := range metadata {
					errorResponse.Extensions[k] = v
				}

			case *errdetails.BadRequest:
				fields := []map[string]string{}
				for _, violation := range t.GetFieldViolations() {
					fields = append(fields, map[string]string{
						"field":   violation.GetField(),
						"message": violation.GetDescription(),
					})
				}
				errorResponse.Extensions["fields"] = fields
			}
		}
	}

	return errorResponse
}

// ParseError ...
func ParseError(err error) map[string]interface{} {
	errMap := make(map[string]interface{})

	errStatus, _ := status.FromError(err)

	errMap["code"] = ToUpperSnakeCase(errStatus.Code().String())
	errMap["message"] = errStatus.Message()

	// Extract the error status detail and insert to extensions.
	errorDetails := errStatus.Details()
	if len(errorDetails) > 0 {
		for _, detail := range errorDetails {
			switch t := detail.(type) {

			case *errdetails.ErrorInfo:
				// Override the error code when ErrorInfo available
				errMap["code"] = ToUpperSnakeCase(t.GetReason())

				metadata := t.GetMetadata()
				for k, v := range metadata {
					errMap[k] = v
				}

			case *errdetails.BadRequest:
				var fields []map[string]string
				for _, violation := range t.GetFieldViolations() {
					fields = append(fields, map[string]string{
						"field":   violation.GetField(),
						"message": violation.GetDescription(),
					})
				}
				errMap["fields"] = fields
			}
		}
	}

	return errMap
}
