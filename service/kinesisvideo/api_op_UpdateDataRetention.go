// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Increases or decreases the stream's data retention period by the value that you
// specify. To indicate whether you want to increase or decrease the data retention
// period, specify the Operation parameter in the request body. In the request, you
// must specify either the StreamName or the StreamARN. The retention period that
// you specify replaces the current value.  <p>This operation requires permission
// for the <code>KinesisVideo:UpdateDataRetention</code> action.</p> <p>Changing
// the data retention period affects the data in the stream as follows:</p> <ul>
// <li> <p>If the data retention period is increased, existing data is retained for
// the new retention period. For example, if the data retention period is increased
// from one hour to seven hours, all existing data is retained for seven hours.</p>
// </li> <li> <p>If the data retention period is decreased, existing data is
// retained for the new retention period. For example, if the data retention period
// is decreased from seven hours to one hour, all existing data is retained for one
// hour, and any data older than one hour is deleted immediately.</p> </li> </ul>
func (c *Client) UpdateDataRetention(ctx context.Context, params *UpdateDataRetentionInput, optFns ...func(*Options)) (*UpdateDataRetentionOutput, error) {
	stack := middleware.NewStack("UpdateDataRetention", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateDataRetentionMiddlewares(stack)
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
	addOpUpdateDataRetentionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDataRetention(options.Region), middleware.Before)

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
			OperationName: "UpdateDataRetention",
			Err:           err,
		}
	}
	out := result.(*UpdateDataRetentionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDataRetentionInput struct {
	// The name of the stream whose retention period you want to change.
	StreamName *string
	// The version of the stream whose retention period you want to change. To get the
	// version, call either the DescribeStream or the ListStreams API.
	CurrentVersion *string
	// The retention period, in hours. The value you specify replaces the current
	// value. The maximum value for this parameter is 87600 (ten years).
	DataRetentionChangeInHours *int32
	// Indicates whether you want to increase or decrease the retention period.
	Operation types.UpdateDataRetentionOperation
	// The Amazon Resource Name (ARN) of the stream whose retention period you want to
	// change.
	StreamARN *string
}

type UpdateDataRetentionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateDataRetentionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDataRetention{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDataRetention{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDataRetention(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "UpdateDataRetention",
	}
}