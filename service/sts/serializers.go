// Code generated by smithy-go-codegen DO NOT EDIT.

package sts

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/httpbinding"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsAwsquery_serializeOpAssumeRole struct {
}

func (*awsAwsquery_serializeOpAssumeRole) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpAssumeRole) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*AssumeRoleInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("AssumeRole")
	body.Key("Version").String("2011-06-15")

	if err := awsAwsquery_serializeDocumentAssumeRoleInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bodyWriter); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpAssumeRoleWithSAML struct {
}

func (*awsAwsquery_serializeOpAssumeRoleWithSAML) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpAssumeRoleWithSAML) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*AssumeRoleWithSAMLInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("AssumeRoleWithSAML")
	body.Key("Version").String("2011-06-15")

	if err := awsAwsquery_serializeDocumentAssumeRoleWithSAMLInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bodyWriter); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpAssumeRoleWithWebIdentity struct {
}

func (*awsAwsquery_serializeOpAssumeRoleWithWebIdentity) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpAssumeRoleWithWebIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*AssumeRoleWithWebIdentityInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("AssumeRoleWithWebIdentity")
	body.Key("Version").String("2011-06-15")

	if err := awsAwsquery_serializeDocumentAssumeRoleWithWebIdentityInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bodyWriter); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpDecodeAuthorizationMessage struct {
}

func (*awsAwsquery_serializeOpDecodeAuthorizationMessage) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpDecodeAuthorizationMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DecodeAuthorizationMessageInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("DecodeAuthorizationMessage")
	body.Key("Version").String("2011-06-15")

	if err := awsAwsquery_serializeDocumentDecodeAuthorizationMessageInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bodyWriter); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpGetAccessKeyInfo struct {
}

func (*awsAwsquery_serializeOpGetAccessKeyInfo) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpGetAccessKeyInfo) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetAccessKeyInfoInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("GetAccessKeyInfo")
	body.Key("Version").String("2011-06-15")

	if err := awsAwsquery_serializeDocumentGetAccessKeyInfoInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bodyWriter); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpGetCallerIdentity struct {
}

func (*awsAwsquery_serializeOpGetCallerIdentity) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpGetCallerIdentity) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetCallerIdentityInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("GetCallerIdentity")
	body.Key("Version").String("2011-06-15")

	_ = input

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bodyWriter); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpGetFederationToken struct {
}

func (*awsAwsquery_serializeOpGetFederationToken) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpGetFederationToken) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetFederationTokenInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("GetFederationToken")
	body.Key("Version").String("2011-06-15")

	if err := awsAwsquery_serializeDocumentGetFederationTokenInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bodyWriter); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsquery_serializeOpGetSessionToken struct {
}

func (*awsAwsquery_serializeOpGetSessionToken) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsquery_serializeOpGetSessionToken) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetSessionTokenInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-www-form-urlencoded")

	bodyWriter := bytes.NewBuffer(nil)
	bodyEncoder := query.NewEncoder(bodyWriter)
	body := bodyEncoder.Object()
	body.Key("Action").String("GetSessionToken")
	body.Key("Version").String("2011-06-15")

	if err := awsAwsquery_serializeDocumentGetSessionTokenInput(input, bodyEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	err = bodyEncoder.Encode()
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bodyWriter); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsquery_serializeDocumentPolicyDescriptorListType(v []*types.PolicyDescriptorType, value query.Value) error {
	if len(v) == 0 {
		return nil
	}
	array := value.Array("member")

	for i := range v {
		if vv := v[i]; vv == nil {
			continue
		}
		av := array.Value()
		if err := awsAwsquery_serializeDocumentPolicyDescriptorType(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsquery_serializeDocumentPolicyDescriptorType(v *types.PolicyDescriptorType, value query.Value) error {
	object := value.Object()
	_ = object

	if v.Arn != nil {
		objectKey := object.Key("arn")
		objectKey.String(*v.Arn)
	}

	return nil
}

func awsAwsquery_serializeDocumentTag(v *types.Tag, value query.Value) error {
	object := value.Object()
	_ = object

	if v.Key != nil {
		objectKey := object.Key("Key")
		objectKey.String(*v.Key)
	}

	if v.Value != nil {
		objectKey := object.Key("Value")
		objectKey.String(*v.Value)
	}

	return nil
}

func awsAwsquery_serializeDocumentTagKeyListType(v []*string, value query.Value) error {
	if len(v) == 0 {
		return nil
	}
	array := value.Array("member")

	for i := range v {
		if vv := v[i]; vv == nil {
			continue
		}
		av := array.Value()
		av.String(*v[i])
	}
	return nil
}

func awsAwsquery_serializeDocumentTagListType(v []*types.Tag, value query.Value) error {
	if len(v) == 0 {
		return nil
	}
	array := value.Array("member")

	for i := range v {
		if vv := v[i]; vv == nil {
			continue
		}
		av := array.Value()
		if err := awsAwsquery_serializeDocumentTag(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsquery_serializeDocumentAssumeRoleInput(v *AssumeRoleInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.DurationSeconds != nil {
		objectKey := object.Key("DurationSeconds")
		objectKey.Integer(*v.DurationSeconds)
	}

	if v.ExternalId != nil {
		objectKey := object.Key("ExternalId")
		objectKey.String(*v.ExternalId)
	}

	if v.Policy != nil {
		objectKey := object.Key("Policy")
		objectKey.String(*v.Policy)
	}

	if v.PolicyArns != nil {
		objectKey := object.Key("PolicyArns")
		if err := awsAwsquery_serializeDocumentPolicyDescriptorListType(v.PolicyArns, objectKey); err != nil {
			return err
		}
	}

	if v.RoleArn != nil {
		objectKey := object.Key("RoleArn")
		objectKey.String(*v.RoleArn)
	}

	if v.RoleSessionName != nil {
		objectKey := object.Key("RoleSessionName")
		objectKey.String(*v.RoleSessionName)
	}

	if v.SerialNumber != nil {
		objectKey := object.Key("SerialNumber")
		objectKey.String(*v.SerialNumber)
	}

	if v.Tags != nil {
		objectKey := object.Key("Tags")
		if err := awsAwsquery_serializeDocumentTagListType(v.Tags, objectKey); err != nil {
			return err
		}
	}

	if v.TokenCode != nil {
		objectKey := object.Key("TokenCode")
		objectKey.String(*v.TokenCode)
	}

	if v.TransitiveTagKeys != nil {
		objectKey := object.Key("TransitiveTagKeys")
		if err := awsAwsquery_serializeDocumentTagKeyListType(v.TransitiveTagKeys, objectKey); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsquery_serializeDocumentAssumeRoleWithSAMLInput(v *AssumeRoleWithSAMLInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.DurationSeconds != nil {
		objectKey := object.Key("DurationSeconds")
		objectKey.Integer(*v.DurationSeconds)
	}

	if v.Policy != nil {
		objectKey := object.Key("Policy")
		objectKey.String(*v.Policy)
	}

	if v.PolicyArns != nil {
		objectKey := object.Key("PolicyArns")
		if err := awsAwsquery_serializeDocumentPolicyDescriptorListType(v.PolicyArns, objectKey); err != nil {
			return err
		}
	}

	if v.PrincipalArn != nil {
		objectKey := object.Key("PrincipalArn")
		objectKey.String(*v.PrincipalArn)
	}

	if v.RoleArn != nil {
		objectKey := object.Key("RoleArn")
		objectKey.String(*v.RoleArn)
	}

	if v.SAMLAssertion != nil {
		objectKey := object.Key("SAMLAssertion")
		objectKey.String(*v.SAMLAssertion)
	}

	return nil
}

func awsAwsquery_serializeDocumentAssumeRoleWithWebIdentityInput(v *AssumeRoleWithWebIdentityInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.DurationSeconds != nil {
		objectKey := object.Key("DurationSeconds")
		objectKey.Integer(*v.DurationSeconds)
	}

	if v.Policy != nil {
		objectKey := object.Key("Policy")
		objectKey.String(*v.Policy)
	}

	if v.PolicyArns != nil {
		objectKey := object.Key("PolicyArns")
		if err := awsAwsquery_serializeDocumentPolicyDescriptorListType(v.PolicyArns, objectKey); err != nil {
			return err
		}
	}

	if v.ProviderId != nil {
		objectKey := object.Key("ProviderId")
		objectKey.String(*v.ProviderId)
	}

	if v.RoleArn != nil {
		objectKey := object.Key("RoleArn")
		objectKey.String(*v.RoleArn)
	}

	if v.RoleSessionName != nil {
		objectKey := object.Key("RoleSessionName")
		objectKey.String(*v.RoleSessionName)
	}

	if v.WebIdentityToken != nil {
		objectKey := object.Key("WebIdentityToken")
		objectKey.String(*v.WebIdentityToken)
	}

	return nil
}

func awsAwsquery_serializeDocumentDecodeAuthorizationMessageInput(v *DecodeAuthorizationMessageInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.EncodedMessage != nil {
		objectKey := object.Key("EncodedMessage")
		objectKey.String(*v.EncodedMessage)
	}

	return nil
}

func awsAwsquery_serializeDocumentGetAccessKeyInfoInput(v *GetAccessKeyInfoInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.AccessKeyId != nil {
		objectKey := object.Key("AccessKeyId")
		objectKey.String(*v.AccessKeyId)
	}

	return nil
}

func awsAwsquery_serializeDocumentGetCallerIdentityInput(v *GetCallerIdentityInput, value query.Value) error {
	object := value.Object()
	_ = object

	return nil
}

func awsAwsquery_serializeDocumentGetFederationTokenInput(v *GetFederationTokenInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.DurationSeconds != nil {
		objectKey := object.Key("DurationSeconds")
		objectKey.Integer(*v.DurationSeconds)
	}

	if v.Name != nil {
		objectKey := object.Key("Name")
		objectKey.String(*v.Name)
	}

	if v.Policy != nil {
		objectKey := object.Key("Policy")
		objectKey.String(*v.Policy)
	}

	if v.PolicyArns != nil {
		objectKey := object.Key("PolicyArns")
		if err := awsAwsquery_serializeDocumentPolicyDescriptorListType(v.PolicyArns, objectKey); err != nil {
			return err
		}
	}

	if v.Tags != nil {
		objectKey := object.Key("Tags")
		if err := awsAwsquery_serializeDocumentTagListType(v.Tags, objectKey); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsquery_serializeDocumentGetSessionTokenInput(v *GetSessionTokenInput, value query.Value) error {
	object := value.Object()
	_ = object

	if v.DurationSeconds != nil {
		objectKey := object.Key("DurationSeconds")
		objectKey.Integer(*v.DurationSeconds)
	}

	if v.SerialNumber != nil {
		objectKey := object.Key("SerialNumber")
		objectKey.String(*v.SerialNumber)
	}

	if v.TokenCode != nil {
		objectKey := object.Key("TokenCode")
		objectKey.String(*v.TokenCode)
	}

	return nil
}
