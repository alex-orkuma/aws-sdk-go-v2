// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves the configuration for asynchronous invocation for a function, version,
// or alias. To configure options for asynchronous invocation, use
// PutFunctionEventInvokeConfig ().
func (c *Client) GetFunctionEventInvokeConfig(ctx context.Context, params *GetFunctionEventInvokeConfigInput, optFns ...func(*Options)) (*GetFunctionEventInvokeConfigOutput, error) {
	stack := middleware.NewStack("GetFunctionEventInvokeConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetFunctionEventInvokeConfigMiddlewares(stack)
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
	addOpGetFunctionEventInvokeConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFunctionEventInvokeConfig(options.Region), middleware.Before)

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
			OperationName: "GetFunctionEventInvokeConfig",
			Err:           err,
		}
	}
	out := result.(*GetFunctionEventInvokeConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFunctionEventInvokeConfigInput struct {
	// A version number or alias name.
	Qualifier *string
	// The name of the Lambda function, version, or alias. Name formats
	//
	//     * Function
	// name - my-function (name-only), my-function:v1 (with alias).
	//
	//     * Function ARN
	// - arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//     * Partial ARN
	// - 123456789012:function:my-function.
	//
	// You can append a version number or alias
	// to any of the formats. The length constraint applies only to the full ARN. If
	// you specify only the function name, it is limited to 64 characters in length.
	FunctionName *string
}

type GetFunctionEventInvokeConfigOutput struct {
	// The maximum age of a request that Lambda sends to a function for processing.
	MaximumEventAgeInSeconds *int32
	// A destination for events after they have been sent to a function for processing.
	// Destinations
	//
	//     * Function - The Amazon Resource Name (ARN) of a Lambda
	// function.
	//
	//     * Queue - The ARN of an SQS queue.
	//
	//     * Topic - The ARN of an
	// SNS topic.
	//
	//     * Event Bus - The ARN of an Amazon EventBridge event bus.
	DestinationConfig *types.DestinationConfig
	// The date and time that the configuration was last updated.
	LastModified *time.Time
	// The maximum number of times to retry when the function returns an error.
	MaximumRetryAttempts *int32
	// The Amazon Resource Name (ARN) of the function.
	FunctionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetFunctionEventInvokeConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetFunctionEventInvokeConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFunctionEventInvokeConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFunctionEventInvokeConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetFunctionEventInvokeConfig",
	}
}
