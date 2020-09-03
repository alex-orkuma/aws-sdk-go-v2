// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// Information about any problems encountered while processing an upload request.
type DocumentServiceException struct {
	Message *string

	Status *string
}

func (e *DocumentServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DocumentServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DocumentServiceException) ErrorCode() string             { return "DocumentServiceException" }
func (e *DocumentServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *DocumentServiceException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *DocumentServiceException) HasMessage() bool {
	return e.Message != nil
}
func (e *DocumentServiceException) GetStatus() string {
	return ptr.ToString(e.Status)
}
func (e *DocumentServiceException) HasStatus() bool {
	return e.Status != nil
}

// Information about any problems encountered while processing a search request.
type SearchException struct {
	Message *string
}

func (e *SearchException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *SearchException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *SearchException) ErrorCode() string             { return "SearchException" }
func (e *SearchException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *SearchException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *SearchException) HasMessage() bool {
	return e.Message != nil
}
