// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the device policy configuration for the specified fleet.
func (c *Client) DescribeDevicePolicyConfiguration(ctx context.Context, params *DescribeDevicePolicyConfigurationInput, optFns ...func(*Options)) (*DescribeDevicePolicyConfigurationOutput, error) {
	stack := middleware.NewStack("DescribeDevicePolicyConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeDevicePolicyConfigurationMiddlewares(stack)
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
	addOpDescribeDevicePolicyConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDevicePolicyConfiguration(options.Region), middleware.Before)

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
			OperationName: "DescribeDevicePolicyConfiguration",
			Err:           err,
		}
	}
	out := result.(*DescribeDevicePolicyConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDevicePolicyConfigurationInput struct {
	// The ARN of the fleet.
	FleetArn *string
}

type DescribeDevicePolicyConfigurationOutput struct {
	// The certificate chain, including intermediate certificates and the root
	// certificate authority certificate used to issue device certificates.
	DeviceCaCertificate *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeDevicePolicyConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeDevicePolicyConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeDevicePolicyConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDevicePolicyConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "DescribeDevicePolicyConfiguration",
	}
}
