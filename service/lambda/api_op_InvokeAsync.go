// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
)

// For asynchronous function invocation, use Invoke (). Invokes a function
// asynchronously.
func (c *Client) InvokeAsync(ctx context.Context, params *InvokeAsyncInput, optFns ...func(*Options)) (*InvokeAsyncOutput, error) {
	stack := middleware.NewStack("InvokeAsync", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpInvokeAsyncMiddlewares(stack)
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
	addOpInvokeAsyncValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInvokeAsync(options.Region), middleware.Before)

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
			OperationName: "InvokeAsync",
			Err:           err,
		}
	}
	out := result.(*InvokeAsyncOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvokeAsyncInput struct {
	// The JSON that you want to provide to your Lambda function as input.
	InvokeArgs io.Reader
	// The name of the Lambda function. Name formats
	//
	//     * Function name -
	// my-function.
	//
	//     * Function ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:my-function.
	//
	//     * Partial ARN -
	// 123456789012:function:my-function.
	//
	// The length constraint applies only to the
	// full ARN. If you specify only the function name, it is limited to 64 characters
	// in length.
	FunctionName *string
}

// A success response (202 Accepted) indicates that the request is queued for
// invocation.
type InvokeAsyncOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpInvokeAsyncMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpInvokeAsync{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpInvokeAsync{}, middleware.After)
}

func newServiceMetadataMiddleware_opInvokeAsync(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "InvokeAsync",
	}
}
