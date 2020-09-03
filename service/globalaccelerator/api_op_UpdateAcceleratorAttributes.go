// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Update the attributes for an accelerator. To see an AWS CLI example of updating
// an accelerator to enable flow logs, scroll down to Example.
func (c *Client) UpdateAcceleratorAttributes(ctx context.Context, params *UpdateAcceleratorAttributesInput, optFns ...func(*Options)) (*UpdateAcceleratorAttributesOutput, error) {
	stack := middleware.NewStack("UpdateAcceleratorAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateAcceleratorAttributesMiddlewares(stack)
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
	addOpUpdateAcceleratorAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAcceleratorAttributes(options.Region), middleware.Before)

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
			OperationName: "UpdateAcceleratorAttributes",
			Err:           err,
		}
	}
	out := result.(*UpdateAcceleratorAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAcceleratorAttributesInput struct {
	// Update the prefix for the location in the Amazon S3 bucket for the flow logs.
	// Attribute is required if FlowLogsEnabled is true. If you don’t specify a prefix,
	// the flow logs are stored in the root of the bucket. If you specify slash (/) for
	// the S3 bucket prefix, the log file bucket folder structure will include a double
	// slash (//), like the following: s3-bucket_name//AWSLogs/aws_account_id
	FlowLogsS3Prefix *string
	// The name of the Amazon S3 bucket for the flow logs. Attribute is required if
	// FlowLogsEnabled is true. The bucket must exist and have a bucket policy that
	// grants AWS Global Accelerator permission to write to the bucket.
	FlowLogsS3Bucket *string
	// The Amazon Resource Name (ARN) of the accelerator that you want to update.
	AcceleratorArn *string
	// Update whether flow logs are enabled. The default value is false. If the value
	// is true, FlowLogsS3Bucket and FlowLogsS3Prefix must be specified. For more
	// information, see Flow Logs
	// (https://docs.aws.amazon.com/global-accelerator/latest/dg/monitoring-global-accelerator.flow-logs.html)
	// in the AWS Global Accelerator Developer Guide.
	FlowLogsEnabled *bool
}

type UpdateAcceleratorAttributesOutput struct {
	// Updated attributes for the accelerator.
	AcceleratorAttributes *types.AcceleratorAttributes

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateAcceleratorAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAcceleratorAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAcceleratorAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateAcceleratorAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "UpdateAcceleratorAttributes",
	}
}
