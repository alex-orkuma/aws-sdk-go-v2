// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a delivery channel object to deliver configuration information to an
// Amazon S3 bucket and Amazon SNS topic. Before you can create a delivery channel,
// you must create a configuration recorder. You can use this action to change the
// Amazon S3 bucket or an Amazon SNS topic of the existing delivery channel. To
// change the Amazon S3 bucket or an Amazon SNS topic, call this action and specify
// the changed values for the S3 bucket and the SNS topic. If you specify a
// different value for either the S3 bucket or the SNS topic, this action will keep
// the existing value for the parameter that is not changed. You can have only one
// delivery channel per region in your account.  </note>
func (c *Client) PutDeliveryChannel(ctx context.Context, params *PutDeliveryChannelInput, optFns ...func(*Options)) (*PutDeliveryChannelOutput, error) {
	stack := middleware.NewStack("PutDeliveryChannel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutDeliveryChannelMiddlewares(stack)
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
	addOpPutDeliveryChannelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutDeliveryChannel(options.Region), middleware.Before)

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
			OperationName: "PutDeliveryChannel",
			Err:           err,
		}
	}
	out := result.(*PutDeliveryChannelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the PutDeliveryChannel () action.
type PutDeliveryChannelInput struct {
	// The configuration delivery channel object that delivers the configuration
	// information to an Amazon S3 bucket and to an Amazon SNS topic.
	DeliveryChannel *types.DeliveryChannel
}

type PutDeliveryChannelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutDeliveryChannelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutDeliveryChannel{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutDeliveryChannel{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutDeliveryChannel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutDeliveryChannel",
	}
}
