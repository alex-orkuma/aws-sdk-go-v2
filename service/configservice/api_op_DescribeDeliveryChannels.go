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

// Returns details about the specified delivery channel. If a delivery channel is
// not specified, this action returns the details of all delivery channels
// associated with the account. Currently, you can specify only one delivery
// channel per region in your account.
func (c *Client) DescribeDeliveryChannels(ctx context.Context, params *DescribeDeliveryChannelsInput, optFns ...func(*Options)) (*DescribeDeliveryChannelsOutput, error) {
	stack := middleware.NewStack("DescribeDeliveryChannels", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeDeliveryChannelsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDeliveryChannels(options.Region), middleware.Before)

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
			OperationName: "DescribeDeliveryChannels",
			Err:           err,
		}
	}
	out := result.(*DescribeDeliveryChannelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeDeliveryChannels () action.
type DescribeDeliveryChannelsInput struct {
	// A list of delivery channel names.
	DeliveryChannelNames []*string
}

// The output for the DescribeDeliveryChannels () action.
type DescribeDeliveryChannelsOutput struct {
	// A list that contains the descriptions of the specified delivery channel.
	DeliveryChannels []*types.DeliveryChannel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeDeliveryChannelsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDeliveryChannels{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDeliveryChannels{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDeliveryChannels(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeDeliveryChannels",
	}
}
