// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// Returned if the request is malformed or contains an error such as an invalid
// parameter value or a missing required parameter.
type BadRequestException struct {
	Message *string

	Code *string
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string             { return "BadRequestException" }
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *BadRequestException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *BadRequestException) HasMessage() bool {
	return e.Message != nil
}
func (e *BadRequestException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *BadRequestException) HasCode() bool {
	return e.Code != nil
}

// Returned if the session doesn't exist anymore because it timed out or expired.
type InvalidSessionException struct {
	Message *string

	Code *string
}

func (e *InvalidSessionException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidSessionException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidSessionException) ErrorCode() string             { return "InvalidSessionException" }
func (e *InvalidSessionException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidSessionException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *InvalidSessionException) HasCode() bool {
	return e.Code != nil
}
func (e *InvalidSessionException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidSessionException) HasMessage() bool {
	return e.Message != nil
}

// Returned if a resource limit such as number of active sessions is exceeded.
type LimitExceededException struct {
	Message *string
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string             { return "LimitExceededException" }
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *LimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// Returned when a transaction cannot be written to the journal due to a failure in
// the verification phase of optimistic concurrency control (OCC).
type OccConflictException struct {
	Message *string
}

func (e *OccConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OccConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OccConflictException) ErrorCode() string             { return "OccConflictException" }
func (e *OccConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *OccConflictException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *OccConflictException) HasMessage() bool {
	return e.Message != nil
}

// Returned when the rate of requests exceeds the allowed throughput.
type RateExceededException struct {
	Message *string
}

func (e *RateExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RateExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RateExceededException) ErrorCode() string             { return "RateExceededException" }
func (e *RateExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *RateExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *RateExceededException) HasMessage() bool {
	return e.Message != nil
}
