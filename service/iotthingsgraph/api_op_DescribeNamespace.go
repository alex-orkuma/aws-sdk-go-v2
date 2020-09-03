// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the latest version of the user's namespace and the public version that it
// is tracking.
func (c *Client) DescribeNamespace(ctx context.Context, params *DescribeNamespaceInput, optFns ...func(*Options)) (*DescribeNamespaceOutput, error) {
	stack := middleware.NewStack("DescribeNamespace", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeNamespaceMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeNamespace(options.Region), middleware.Before)

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
			OperationName: "DescribeNamespace",
			Err:           err,
		}
	}
	out := result.(*DescribeNamespaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeNamespaceInput struct {
	// The name of the user's namespace. Set this to aws to get the public namespace.
	NamespaceName *string
}

type DescribeNamespaceOutput struct {
	// The version of the public namespace that the latest version is tracking.
	TrackingNamespaceVersion *int64
	// The name of the public namespace that the latest namespace version is tracking.
	TrackingNamespaceName *string
	// The version of the user's namespace to describe.
	NamespaceVersion *int64
	// The name of the namespace.
	NamespaceName *string
	// The ARN of the namespace.
	NamespaceArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeNamespaceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeNamespace{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeNamespace{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeNamespace(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "DescribeNamespace",
	}
}
