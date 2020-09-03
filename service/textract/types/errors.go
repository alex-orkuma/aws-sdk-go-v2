// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// You aren't authorized to perform the action.
type AccessDeniedException struct {
	Message *string

	Code *string
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

// Amazon Textract isn't able to read the document.
type BadDocumentException struct {
	Message *string

	Code *string
}

func (e *BadDocumentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadDocumentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadDocumentException) ErrorCode() string             { return "BadDocumentException" }
func (e *BadDocumentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *BadDocumentException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *BadDocumentException) HasMessage() bool {
	return e.Message != nil
}
func (e *BadDocumentException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *BadDocumentException) HasCode() bool {
	return e.Code != nil
}

// The document can't be processed because it's too large. The maximum document
// size for synchronous operations 5 MB. The maximum document size for asynchronous
// operations is 500 MB for PDF files.
type DocumentTooLargeException struct {
	Message *string

	Code *string
}

func (e *DocumentTooLargeException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *DocumentTooLargeException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *DocumentTooLargeException) ErrorCode() string             { return "DocumentTooLargeException" }
func (e *DocumentTooLargeException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *DocumentTooLargeException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *DocumentTooLargeException) HasCode() bool {
	return e.Code != nil
}
func (e *DocumentTooLargeException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *DocumentTooLargeException) HasMessage() bool {
	return e.Message != nil
}

// Indicates you have exceeded the maximum number of active human in the loop
// workflows available
type HumanLoopQuotaExceededException struct {
	Message *string

	Code         *string
	ResourceType *string
	QuotaCode    *string
	ServiceCode  *string
}

func (e *HumanLoopQuotaExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *HumanLoopQuotaExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *HumanLoopQuotaExceededException) ErrorCode() string {
	return "HumanLoopQuotaExceededException"
}
func (e *HumanLoopQuotaExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *HumanLoopQuotaExceededException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *HumanLoopQuotaExceededException) HasCode() bool {
	return e.Code != nil
}
func (e *HumanLoopQuotaExceededException) GetResourceType() string {
	return ptr.ToString(e.ResourceType)
}
func (e *HumanLoopQuotaExceededException) HasResourceType() bool {
	return e.ResourceType != nil
}
func (e *HumanLoopQuotaExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *HumanLoopQuotaExceededException) HasMessage() bool {
	return e.Message != nil
}
func (e *HumanLoopQuotaExceededException) GetQuotaCode() string {
	return ptr.ToString(e.QuotaCode)
}
func (e *HumanLoopQuotaExceededException) HasQuotaCode() bool {
	return e.QuotaCode != nil
}
func (e *HumanLoopQuotaExceededException) GetServiceCode() string {
	return ptr.ToString(e.ServiceCode)
}
func (e *HumanLoopQuotaExceededException) HasServiceCode() bool {
	return e.ServiceCode != nil
}

// A ClientRequestToken input parameter was reused with an operation, but at least
// one of the other input parameters is different from the previous call to the
// operation.
type IdempotentParameterMismatchException struct {
	Message *string

	Code *string
}

func (e *IdempotentParameterMismatchException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *IdempotentParameterMismatchException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *IdempotentParameterMismatchException) ErrorCode() string {
	return "IdempotentParameterMismatchException"
}
func (e *IdempotentParameterMismatchException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *IdempotentParameterMismatchException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *IdempotentParameterMismatchException) HasMessage() bool {
	return e.Message != nil
}
func (e *IdempotentParameterMismatchException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *IdempotentParameterMismatchException) HasCode() bool {
	return e.Code != nil
}

// Amazon Textract experienced a service issue. Try your call again.
type InternalServerError struct {
	Message *string

	Code *string
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerError) ErrorCode() string             { return "InternalServerError" }
func (e *InternalServerError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *InternalServerError) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InternalServerError) HasMessage() bool {
	return e.Message != nil
}
func (e *InternalServerError) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *InternalServerError) HasCode() bool {
	return e.Code != nil
}

// An invalid job identifier was passed to GetDocumentAnalysis () or to
// GetDocumentAnalysis ().
type InvalidJobIdException struct {
	Message *string

	Code *string
}

func (e *InvalidJobIdException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidJobIdException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidJobIdException) ErrorCode() string             { return "InvalidJobIdException" }
func (e *InvalidJobIdException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidJobIdException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidJobIdException) HasMessage() bool {
	return e.Message != nil
}
func (e *InvalidJobIdException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *InvalidJobIdException) HasCode() bool {
	return e.Code != nil
}

// An input parameter violated a constraint. For example, in synchronous
// operations, an InvalidParameterException exception occurs when neither of the
// S3Object or Bytes values are supplied in the Document request parameter.
// Validate your parameter before calling the API operation again.
type InvalidParameterException struct {
	Message *string

	Code *string
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidParameterException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *InvalidParameterException) HasCode() bool {
	return e.Code != nil
}
func (e *InvalidParameterException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidParameterException) HasMessage() bool {
	return e.Message != nil
}

// Amazon Textract is unable to access the S3 object that's specified in the
// request.
type InvalidS3ObjectException struct {
	Message *string

	Code *string
}

func (e *InvalidS3ObjectException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidS3ObjectException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidS3ObjectException) ErrorCode() string             { return "InvalidS3ObjectException" }
func (e *InvalidS3ObjectException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidS3ObjectException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *InvalidS3ObjectException) HasCode() bool {
	return e.Code != nil
}
func (e *InvalidS3ObjectException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidS3ObjectException) HasMessage() bool {
	return e.Message != nil
}

// An Amazon Textract service limit was exceeded. For example, if you start too
// many asynchronous jobs concurrently, calls to start operations
// (StartDocumentTextDetection, for example) raise a LimitExceededException
// exception (HTTP status code: 400) until the number of concurrently running jobs
// is below the Amazon Textract service limit.
type LimitExceededException struct {
	Message *string

	Code *string
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
func (e *LimitExceededException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *LimitExceededException) HasCode() bool {
	return e.Code != nil
}
func (e *LimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *LimitExceededException) HasMessage() bool {
	return e.Message != nil
}

// The number of requests exceeded your throughput limit. If you want to increase
// this limit, contact Amazon Textract.
type ProvisionedThroughputExceededException struct {
	Message *string

	Code *string
}

func (e *ProvisionedThroughputExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ProvisionedThroughputExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ProvisionedThroughputExceededException) ErrorCode() string {
	return "ProvisionedThroughputExceededException"
}
func (e *ProvisionedThroughputExceededException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *ProvisionedThroughputExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ProvisionedThroughputExceededException) HasMessage() bool {
	return e.Message != nil
}
func (e *ProvisionedThroughputExceededException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *ProvisionedThroughputExceededException) HasCode() bool {
	return e.Code != nil
}

// Amazon Textract is temporarily unable to process the request. Try your call
// again.
type ThrottlingException struct {
	Message *string

	Code *string
}

func (e *ThrottlingException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ThrottlingException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ThrottlingException) ErrorCode() string             { return "ThrottlingException" }
func (e *ThrottlingException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *ThrottlingException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *ThrottlingException) HasCode() bool {
	return e.Code != nil
}
func (e *ThrottlingException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ThrottlingException) HasMessage() bool {
	return e.Message != nil
}

// The format of the input document isn't supported. Documents for synchronous
// operations can be in PNG or JPEG format. Documents for asynchronous operations
// can also be in PDF format.
type UnsupportedDocumentException struct {
	Message *string

	Code *string
}

func (e *UnsupportedDocumentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedDocumentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedDocumentException) ErrorCode() string             { return "UnsupportedDocumentException" }
func (e *UnsupportedDocumentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *UnsupportedDocumentException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *UnsupportedDocumentException) HasMessage() bool {
	return e.Message != nil
}
func (e *UnsupportedDocumentException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *UnsupportedDocumentException) HasCode() bool {
	return e.Code != nil
}
