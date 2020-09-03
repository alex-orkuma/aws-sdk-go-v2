// Code generated by smithy-go-codegen DO NOT EDIT.

package support

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/support/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsAwsjson11_serializeOpAddAttachmentsToSet struct {
}

func (*awsAwsjson11_serializeOpAddAttachmentsToSet) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpAddAttachmentsToSet) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*AddAttachmentsToSetInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSSupport_20130415.AddAttachmentsToSet")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentAddAttachmentsToSetInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpAddCommunicationToCase struct {
}

func (*awsAwsjson11_serializeOpAddCommunicationToCase) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpAddCommunicationToCase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*AddCommunicationToCaseInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSSupport_20130415.AddCommunicationToCase")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentAddCommunicationToCaseInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpCreateCase struct {
}

func (*awsAwsjson11_serializeOpCreateCase) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpCreateCase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateCaseInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSSupport_20130415.CreateCase")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentCreateCaseInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeAttachment struct {
}

func (*awsAwsjson11_serializeOpDescribeAttachment) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeAttachment) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeAttachmentInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSSupport_20130415.DescribeAttachment")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentDescribeAttachmentInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeCases struct {
}

func (*awsAwsjson11_serializeOpDescribeCases) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeCases) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeCasesInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSSupport_20130415.DescribeCases")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentDescribeCasesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeCommunications struct {
}

func (*awsAwsjson11_serializeOpDescribeCommunications) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeCommunications) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeCommunicationsInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSSupport_20130415.DescribeCommunications")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentDescribeCommunicationsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeServices struct {
}

func (*awsAwsjson11_serializeOpDescribeServices) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeServices) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeServicesInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSSupport_20130415.DescribeServices")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentDescribeServicesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeSeverityLevels struct {
}

func (*awsAwsjson11_serializeOpDescribeSeverityLevels) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeSeverityLevels) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeSeverityLevelsInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSSupport_20130415.DescribeSeverityLevels")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentDescribeSeverityLevelsInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeTrustedAdvisorCheckRefreshStatuses struct {
}

func (*awsAwsjson11_serializeOpDescribeTrustedAdvisorCheckRefreshStatuses) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeTrustedAdvisorCheckRefreshStatuses) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeTrustedAdvisorCheckRefreshStatusesInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSSupport_20130415.DescribeTrustedAdvisorCheckRefreshStatuses")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentDescribeTrustedAdvisorCheckRefreshStatusesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeTrustedAdvisorCheckResult struct {
}

func (*awsAwsjson11_serializeOpDescribeTrustedAdvisorCheckResult) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeTrustedAdvisorCheckResult) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeTrustedAdvisorCheckResultInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSSupport_20130415.DescribeTrustedAdvisorCheckResult")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentDescribeTrustedAdvisorCheckResultInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeTrustedAdvisorChecks struct {
}

func (*awsAwsjson11_serializeOpDescribeTrustedAdvisorChecks) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeTrustedAdvisorChecks) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeTrustedAdvisorChecksInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSSupport_20130415.DescribeTrustedAdvisorChecks")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentDescribeTrustedAdvisorChecksInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpDescribeTrustedAdvisorCheckSummaries struct {
}

func (*awsAwsjson11_serializeOpDescribeTrustedAdvisorCheckSummaries) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpDescribeTrustedAdvisorCheckSummaries) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DescribeTrustedAdvisorCheckSummariesInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSSupport_20130415.DescribeTrustedAdvisorCheckSummaries")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentDescribeTrustedAdvisorCheckSummariesInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpRefreshTrustedAdvisorCheck struct {
}

func (*awsAwsjson11_serializeOpRefreshTrustedAdvisorCheck) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpRefreshTrustedAdvisorCheck) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*RefreshTrustedAdvisorCheckInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSSupport_20130415.RefreshTrustedAdvisorCheck")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentRefreshTrustedAdvisorCheckInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpResolveCase struct {
}

func (*awsAwsjson11_serializeOpResolveCase) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpResolveCase) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ResolveCaseInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSSupport_20130415.ResolveCase")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentResolveCaseInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentAttachment(v *types.Attachment, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Data != nil {
		ok := object.Key("data")
		ok.Base64EncodeBytes(v.Data)
	}

	if v.FileName != nil {
		ok := object.Key("fileName")
		ok.String(*v.FileName)
	}

	return nil
}

func awsAwsjson11_serializeDocumentAttachments(v []*types.Attachment, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			av.Null()
			continue
		}
		if err := awsAwsjson11_serializeDocumentAttachment(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentCaseIdList(v []*string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			av.Null()
			continue
		}
		av.String(*v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentCcEmailAddressList(v []*string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			av.Null()
			continue
		}
		av.String(*v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentServiceCodeList(v []*string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			av.Null()
			continue
		}
		av.String(*v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentStringList(v []*string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			av.Null()
			continue
		}
		av.String(*v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentAddAttachmentsToSetInput(v *AddAttachmentsToSetInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Attachments != nil {
		ok := object.Key("attachments")
		if err := awsAwsjson11_serializeDocumentAttachments(v.Attachments, ok); err != nil {
			return err
		}
	}

	if v.AttachmentSetId != nil {
		ok := object.Key("attachmentSetId")
		ok.String(*v.AttachmentSetId)
	}

	return nil
}

func awsAwsjson11_serializeDocumentAddCommunicationToCaseInput(v *AddCommunicationToCaseInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AttachmentSetId != nil {
		ok := object.Key("attachmentSetId")
		ok.String(*v.AttachmentSetId)
	}

	if v.CaseId != nil {
		ok := object.Key("caseId")
		ok.String(*v.CaseId)
	}

	if v.CcEmailAddresses != nil {
		ok := object.Key("ccEmailAddresses")
		if err := awsAwsjson11_serializeDocumentCcEmailAddressList(v.CcEmailAddresses, ok); err != nil {
			return err
		}
	}

	if v.CommunicationBody != nil {
		ok := object.Key("communicationBody")
		ok.String(*v.CommunicationBody)
	}

	return nil
}

func awsAwsjson11_serializeDocumentCreateCaseInput(v *CreateCaseInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AttachmentSetId != nil {
		ok := object.Key("attachmentSetId")
		ok.String(*v.AttachmentSetId)
	}

	if v.CategoryCode != nil {
		ok := object.Key("categoryCode")
		ok.String(*v.CategoryCode)
	}

	if v.CcEmailAddresses != nil {
		ok := object.Key("ccEmailAddresses")
		if err := awsAwsjson11_serializeDocumentCcEmailAddressList(v.CcEmailAddresses, ok); err != nil {
			return err
		}
	}

	if v.CommunicationBody != nil {
		ok := object.Key("communicationBody")
		ok.String(*v.CommunicationBody)
	}

	if v.IssueType != nil {
		ok := object.Key("issueType")
		ok.String(*v.IssueType)
	}

	if v.Language != nil {
		ok := object.Key("language")
		ok.String(*v.Language)
	}

	if v.ServiceCode != nil {
		ok := object.Key("serviceCode")
		ok.String(*v.ServiceCode)
	}

	if v.SeverityCode != nil {
		ok := object.Key("severityCode")
		ok.String(*v.SeverityCode)
	}

	if v.Subject != nil {
		ok := object.Key("subject")
		ok.String(*v.Subject)
	}

	return nil
}

func awsAwsjson11_serializeDocumentDescribeAttachmentInput(v *DescribeAttachmentInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AttachmentId != nil {
		ok := object.Key("attachmentId")
		ok.String(*v.AttachmentId)
	}

	return nil
}

func awsAwsjson11_serializeDocumentDescribeCasesInput(v *DescribeCasesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AfterTime != nil {
		ok := object.Key("afterTime")
		ok.String(*v.AfterTime)
	}

	if v.BeforeTime != nil {
		ok := object.Key("beforeTime")
		ok.String(*v.BeforeTime)
	}

	if v.CaseIdList != nil {
		ok := object.Key("caseIdList")
		if err := awsAwsjson11_serializeDocumentCaseIdList(v.CaseIdList, ok); err != nil {
			return err
		}
	}

	if v.DisplayId != nil {
		ok := object.Key("displayId")
		ok.String(*v.DisplayId)
	}

	if v.IncludeCommunications != nil {
		ok := object.Key("includeCommunications")
		ok.Boolean(*v.IncludeCommunications)
	}

	if v.IncludeResolvedCases != nil {
		ok := object.Key("includeResolvedCases")
		ok.Boolean(*v.IncludeResolvedCases)
	}

	if v.Language != nil {
		ok := object.Key("language")
		ok.String(*v.Language)
	}

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsAwsjson11_serializeDocumentDescribeCommunicationsInput(v *DescribeCommunicationsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AfterTime != nil {
		ok := object.Key("afterTime")
		ok.String(*v.AfterTime)
	}

	if v.BeforeTime != nil {
		ok := object.Key("beforeTime")
		ok.String(*v.BeforeTime)
	}

	if v.CaseId != nil {
		ok := object.Key("caseId")
		ok.String(*v.CaseId)
	}

	if v.MaxResults != nil {
		ok := object.Key("maxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("nextToken")
		ok.String(*v.NextToken)
	}

	return nil
}

func awsAwsjson11_serializeDocumentDescribeServicesInput(v *DescribeServicesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Language != nil {
		ok := object.Key("language")
		ok.String(*v.Language)
	}

	if v.ServiceCodeList != nil {
		ok := object.Key("serviceCodeList")
		if err := awsAwsjson11_serializeDocumentServiceCodeList(v.ServiceCodeList, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentDescribeSeverityLevelsInput(v *DescribeSeverityLevelsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Language != nil {
		ok := object.Key("language")
		ok.String(*v.Language)
	}

	return nil
}

func awsAwsjson11_serializeDocumentDescribeTrustedAdvisorCheckRefreshStatusesInput(v *DescribeTrustedAdvisorCheckRefreshStatusesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CheckIds != nil {
		ok := object.Key("checkIds")
		if err := awsAwsjson11_serializeDocumentStringList(v.CheckIds, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentDescribeTrustedAdvisorCheckResultInput(v *DescribeTrustedAdvisorCheckResultInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CheckId != nil {
		ok := object.Key("checkId")
		ok.String(*v.CheckId)
	}

	if v.Language != nil {
		ok := object.Key("language")
		ok.String(*v.Language)
	}

	return nil
}

func awsAwsjson11_serializeDocumentDescribeTrustedAdvisorChecksInput(v *DescribeTrustedAdvisorChecksInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Language != nil {
		ok := object.Key("language")
		ok.String(*v.Language)
	}

	return nil
}

func awsAwsjson11_serializeDocumentDescribeTrustedAdvisorCheckSummariesInput(v *DescribeTrustedAdvisorCheckSummariesInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CheckIds != nil {
		ok := object.Key("checkIds")
		if err := awsAwsjson11_serializeDocumentStringList(v.CheckIds, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentRefreshTrustedAdvisorCheckInput(v *RefreshTrustedAdvisorCheckInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CheckId != nil {
		ok := object.Key("checkId")
		ok.String(*v.CheckId)
	}

	return nil
}

func awsAwsjson11_serializeDocumentResolveCaseInput(v *ResolveCaseInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CaseId != nil {
		ok := object.Key("caseId")
		ok.String(*v.CaseId)
	}

	return nil
}
