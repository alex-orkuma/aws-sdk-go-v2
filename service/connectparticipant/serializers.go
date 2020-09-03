// Code generated by smithy-go-codegen DO NOT EDIT.

package connectparticipant

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/connectparticipant/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsRestjson1_serializeOpCreateParticipantConnection struct {
}

func (*awsRestjson1_serializeOpCreateParticipantConnection) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpCreateParticipantConnection) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*CreateParticipantConnectionInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/participant/connection")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeHttpBindingsCreateParticipantConnectionInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeDocumentCreateParticipantConnectionInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeHttpBindingsCreateParticipantConnectionInput(v *CreateParticipantConnectionInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ParticipantToken != nil {
		locationName := "X-Amz-Bearer"
		if len(*v.ParticipantToken) > 0 {
			encoder.SetHeader(locationName).String(*v.ParticipantToken)
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentCreateParticipantConnectionInput(v *CreateParticipantConnectionInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Type != nil {
		ok := object.Key("Type")
		if err := awsRestjson1_serializeDocumentConnectionTypeList(v.Type, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpDisconnectParticipant struct {
}

func (*awsRestjson1_serializeOpDisconnectParticipant) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpDisconnectParticipant) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*DisconnectParticipantInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/participant/disconnect")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeHttpBindingsDisconnectParticipantInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeDocumentDisconnectParticipantInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeHttpBindingsDisconnectParticipantInput(v *DisconnectParticipantInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ConnectionToken != nil {
		locationName := "X-Amz-Bearer"
		if len(*v.ConnectionToken) > 0 {
			encoder.SetHeader(locationName).String(*v.ConnectionToken)
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentDisconnectParticipantInput(v *DisconnectParticipantInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientToken != nil {
		ok := object.Key("ClientToken")
		ok.String(*v.ClientToken)
	}

	return nil
}

type awsRestjson1_serializeOpGetTranscript struct {
}

func (*awsRestjson1_serializeOpGetTranscript) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpGetTranscript) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GetTranscriptInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/participant/transcript")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeHttpBindingsGetTranscriptInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeDocumentGetTranscriptInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeHttpBindingsGetTranscriptInput(v *GetTranscriptInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ConnectionToken != nil {
		locationName := "X-Amz-Bearer"
		if len(*v.ConnectionToken) > 0 {
			encoder.SetHeader(locationName).String(*v.ConnectionToken)
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentGetTranscriptInput(v *GetTranscriptInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ContactId != nil {
		ok := object.Key("ContactId")
		ok.String(*v.ContactId)
	}

	if v.MaxResults != nil {
		ok := object.Key("MaxResults")
		ok.Integer(*v.MaxResults)
	}

	if v.NextToken != nil {
		ok := object.Key("NextToken")
		ok.String(*v.NextToken)
	}

	if len(v.ScanDirection) > 0 {
		ok := object.Key("ScanDirection")
		ok.String(string(v.ScanDirection))
	}

	if len(v.SortOrder) > 0 {
		ok := object.Key("SortOrder")
		ok.String(string(v.SortOrder))
	}

	if v.StartPosition != nil {
		ok := object.Key("StartPosition")
		if err := awsRestjson1_serializeDocumentStartPosition(v.StartPosition, ok); err != nil {
			return err
		}
	}

	return nil
}

type awsRestjson1_serializeOpSendEvent struct {
}

func (*awsRestjson1_serializeOpSendEvent) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpSendEvent) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*SendEventInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/participant/event")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeHttpBindingsSendEventInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeDocumentSendEventInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeHttpBindingsSendEventInput(v *SendEventInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ConnectionToken != nil {
		locationName := "X-Amz-Bearer"
		if len(*v.ConnectionToken) > 0 {
			encoder.SetHeader(locationName).String(*v.ConnectionToken)
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentSendEventInput(v *SendEventInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientToken != nil {
		ok := object.Key("ClientToken")
		ok.String(*v.ClientToken)
	}

	if v.Content != nil {
		ok := object.Key("Content")
		ok.String(*v.Content)
	}

	if v.ContentType != nil {
		ok := object.Key("ContentType")
		ok.String(*v.ContentType)
	}

	return nil
}

type awsRestjson1_serializeOpSendMessage struct {
}

func (*awsRestjson1_serializeOpSendMessage) ID() string {
	return "OperationSerializer"
}

func (m *awsRestjson1_serializeOpSendMessage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*SendMessageInput)
	_ = input
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	opPath, opQuery := httpbinding.SplitURI("/participant/message")
	request.URL.Path = opPath
	if len(request.URL.RawQuery) > 0 {
		request.URL.RawQuery = "&" + opQuery
	} else {
		request.URL.RawQuery = opQuery
	}

	request.Method = "POST"
	restEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if err := awsRestjson1_serializeHttpBindingsSendMessageInput(input, restEncoder); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	restEncoder.SetHeader("Content-Type").String("application/json")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsRestjson1_serializeDocumentSendMessageInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = restEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsRestjson1_serializeHttpBindingsSendMessageInput(v *SendMessageInput, encoder *httpbinding.Encoder) error {
	if v == nil {
		return fmt.Errorf("unsupported serialization of nil %T", v)
	}

	if v.ConnectionToken != nil {
		locationName := "X-Amz-Bearer"
		if len(*v.ConnectionToken) > 0 {
			encoder.SetHeader(locationName).String(*v.ConnectionToken)
		}
	}

	return nil
}

func awsRestjson1_serializeDocumentSendMessageInput(v *SendMessageInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ClientToken != nil {
		ok := object.Key("ClientToken")
		ok.String(*v.ClientToken)
	}

	if v.Content != nil {
		ok := object.Key("Content")
		ok.String(*v.Content)
	}

	if v.ContentType != nil {
		ok := object.Key("ContentType")
		ok.String(*v.ContentType)
	}

	return nil
}

func awsRestjson1_serializeDocumentConnectionTypeList(v []types.ConnectionType, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsRestjson1_serializeDocumentStartPosition(v *types.StartPosition, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.AbsoluteTime != nil {
		ok := object.Key("AbsoluteTime")
		ok.String(*v.AbsoluteTime)
	}

	if v.Id != nil {
		ok := object.Key("Id")
		ok.String(*v.Id)
	}

	if v.MostRecent != nil {
		ok := object.Key("MostRecent")
		ok.Integer(*v.MostRecent)
	}

	return nil
}
