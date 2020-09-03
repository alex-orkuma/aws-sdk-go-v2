// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of objects that contain information about events in a flow
// execution.
func (c *Client) ListFlowExecutionMessages(ctx context.Context, params *ListFlowExecutionMessagesInput, optFns ...func(*Options)) (*ListFlowExecutionMessagesOutput, error) {
	stack := middleware.NewStack("ListFlowExecutionMessages", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListFlowExecutionMessagesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListFlowExecutionMessagesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListFlowExecutionMessages(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ListFlowExecutionMessages",
			Err:           err,
		}
	}
	out := result.(*ListFlowExecutionMessagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFlowExecutionMessagesInput struct {
	// The maximum number of results to return in the response.
	MaxResults *int32
	// The ID of the flow execution.
	FlowExecutionId *string
	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string
}

type ListFlowExecutionMessagesOutput struct {
	// A list of objects that contain information about events in the specified flow
	// execution.
	Messages []*types.FlowExecutionMessage
	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListFlowExecutionMessagesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListFlowExecutionMessages{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListFlowExecutionMessages{}, middleware.After)
}

func newServiceMetadataMiddleware_opListFlowExecutionMessages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "ListFlowExecutionMessages",
	}
}
