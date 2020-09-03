// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// Lightsail throws this exception when the user cannot be authenticated or uses
// invalid credentials to access a resource.
type AccessDeniedException struct {
	Message *string

	Tip  *string
	Code *string
	Docs *string
}

func (e *AccessDeniedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccessDeniedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccessDeniedException) ErrorCode() string             { return "AccessDeniedException" }
func (e *AccessDeniedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *AccessDeniedException) GetTip() string {
	return ptr.ToString(e.Tip)
}
func (e *AccessDeniedException) HasTip() bool {
	return e.Tip != nil
}
func (e *AccessDeniedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *AccessDeniedException) HasMessage() bool {
	return e.Message != nil
}
func (e *AccessDeniedException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *AccessDeniedException) HasCode() bool {
	return e.Code != nil
}
func (e *AccessDeniedException) GetDocs() string {
	return ptr.ToString(e.Docs)
}
func (e *AccessDeniedException) HasDocs() bool {
	return e.Docs != nil
}

// Lightsail throws this exception when an account is still in the setup in
// progress state.
type AccountSetupInProgressException struct {
	Message *string

	Docs *string
	Code *string
	Tip  *string
}

func (e *AccountSetupInProgressException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *AccountSetupInProgressException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *AccountSetupInProgressException) ErrorCode() string {
	return "AccountSetupInProgressException"
}
func (e *AccountSetupInProgressException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *AccountSetupInProgressException) GetDocs() string {
	return ptr.ToString(e.Docs)
}
func (e *AccountSetupInProgressException) HasDocs() bool {
	return e.Docs != nil
}
func (e *AccountSetupInProgressException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *AccountSetupInProgressException) HasCode() bool {
	return e.Code != nil
}
func (e *AccountSetupInProgressException) GetTip() string {
	return ptr.ToString(e.Tip)
}
func (e *AccountSetupInProgressException) HasTip() bool {
	return e.Tip != nil
}
func (e *AccountSetupInProgressException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *AccountSetupInProgressException) HasMessage() bool {
	return e.Message != nil
}

// Lightsail throws this exception when user input does not conform to the
// validation rules of an input field. Domain-related APIs are only available in
// the N. Virginia (us-east-1) Region. Please set your AWS Region configuration to
// us-east-1 to create, view, or edit these resources.
type InvalidInputException struct {
	Message *string

	Tip  *string
	Code *string
	Docs *string
}

func (e *InvalidInputException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidInputException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidInputException) ErrorCode() string             { return "InvalidInputException" }
func (e *InvalidInputException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidInputException) GetTip() string {
	return ptr.ToString(e.Tip)
}
func (e *InvalidInputException) HasTip() bool {
	return e.Tip != nil
}
func (e *InvalidInputException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *InvalidInputException) HasCode() bool {
	return e.Code != nil
}
func (e *InvalidInputException) GetDocs() string {
	return ptr.ToString(e.Docs)
}
func (e *InvalidInputException) HasDocs() bool {
	return e.Docs != nil
}
func (e *InvalidInputException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidInputException) HasMessage() bool {
	return e.Message != nil
}

// Lightsail throws this exception when it cannot find a resource.
type NotFoundException struct {
	Message *string

	Tip  *string
	Docs *string
	Code *string
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string             { return "NotFoundException" }
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *NotFoundException) GetTip() string {
	return ptr.ToString(e.Tip)
}
func (e *NotFoundException) HasTip() bool {
	return e.Tip != nil
}
func (e *NotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *NotFoundException) HasMessage() bool {
	return e.Message != nil
}
func (e *NotFoundException) GetDocs() string {
	return ptr.ToString(e.Docs)
}
func (e *NotFoundException) HasDocs() bool {
	return e.Docs != nil
}
func (e *NotFoundException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *NotFoundException) HasCode() bool {
	return e.Code != nil
}

// Lightsail throws this exception when an operation fails to execute.
type OperationFailureException struct {
	Message *string

	Code *string
	Tip  *string
	Docs *string
}

func (e *OperationFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *OperationFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *OperationFailureException) ErrorCode() string             { return "OperationFailureException" }
func (e *OperationFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *OperationFailureException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *OperationFailureException) HasCode() bool {
	return e.Code != nil
}
func (e *OperationFailureException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *OperationFailureException) HasMessage() bool {
	return e.Message != nil
}
func (e *OperationFailureException) GetTip() string {
	return ptr.ToString(e.Tip)
}
func (e *OperationFailureException) HasTip() bool {
	return e.Tip != nil
}
func (e *OperationFailureException) GetDocs() string {
	return ptr.ToString(e.Docs)
}
func (e *OperationFailureException) HasDocs() bool {
	return e.Docs != nil
}

// A general service exception.
type ServiceException struct {
	Message *string

	Docs *string
	Tip  *string
	Code *string
}

func (e *ServiceException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceException) ErrorCode() string             { return "ServiceException" }
func (e *ServiceException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *ServiceException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ServiceException) HasMessage() bool {
	return e.Message != nil
}
func (e *ServiceException) GetDocs() string {
	return ptr.ToString(e.Docs)
}
func (e *ServiceException) HasDocs() bool {
	return e.Docs != nil
}
func (e *ServiceException) GetTip() string {
	return ptr.ToString(e.Tip)
}
func (e *ServiceException) HasTip() bool {
	return e.Tip != nil
}
func (e *ServiceException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *ServiceException) HasCode() bool {
	return e.Code != nil
}

// Lightsail throws this exception when the user has not been authenticated.
type UnauthenticatedException struct {
	Message *string

	Tip  *string
	Code *string
	Docs *string
}

func (e *UnauthenticatedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnauthenticatedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnauthenticatedException) ErrorCode() string             { return "UnauthenticatedException" }
func (e *UnauthenticatedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *UnauthenticatedException) GetTip() string {
	return ptr.ToString(e.Tip)
}
func (e *UnauthenticatedException) HasTip() bool {
	return e.Tip != nil
}
func (e *UnauthenticatedException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *UnauthenticatedException) HasCode() bool {
	return e.Code != nil
}
func (e *UnauthenticatedException) GetDocs() string {
	return ptr.ToString(e.Docs)
}
func (e *UnauthenticatedException) HasDocs() bool {
	return e.Docs != nil
}
func (e *UnauthenticatedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *UnauthenticatedException) HasMessage() bool {
	return e.Message != nil
}
