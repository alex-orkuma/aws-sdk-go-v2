// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays the details of a flow. The response includes the flow ARN, name, and
// Availability Zone, as well as details about the source, outputs, and
// entitlements.
func (c *Client) DescribeFlow(ctx context.Context, params *DescribeFlowInput, optFns ...func(*Options)) (*DescribeFlowOutput, error) {
	stack := middleware.NewStack("DescribeFlow", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeFlowMiddlewares(stack)
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
	addOpDescribeFlowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFlow(options.Region), middleware.Before)

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
			OperationName: "DescribeFlow",
			Err:           err,
		}
	}
	out := result.(*DescribeFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeFlowInput struct {
	// The ARN of the flow that you want to describe.
	FlowArn *string
}

type DescribeFlowOutput struct {
	// Messages that provide the state of the flow.
	Messages *types.Messages
	// The settings for a flow, including its source, outputs, and entitlements.
	Flow *types.Flow

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeFlowMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFlow{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFlow{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeFlow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "DescribeFlow",
	}
}
