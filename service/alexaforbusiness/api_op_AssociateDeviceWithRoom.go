// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a device with a given room. This applies all the settings from the
// room profile to the device, and all the skills in any skill groups added to that
// room. This operation requires the device to be online, or else a manual sync is
// required.
func (c *Client) AssociateDeviceWithRoom(ctx context.Context, params *AssociateDeviceWithRoomInput, optFns ...func(*Options)) (*AssociateDeviceWithRoomOutput, error) {
	stack := middleware.NewStack("AssociateDeviceWithRoom", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateDeviceWithRoomMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateDeviceWithRoom(options.Region), middleware.Before)

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
			OperationName: "AssociateDeviceWithRoom",
			Err:           err,
		}
	}
	out := result.(*AssociateDeviceWithRoomOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateDeviceWithRoomInput struct {
	// The ARN of the room with which to associate the device. Required.
	RoomArn *string
	// The ARN of the device to associate to a room. Required.
	DeviceArn *string
}

type AssociateDeviceWithRoomOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateDeviceWithRoomMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateDeviceWithRoom{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateDeviceWithRoom{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateDeviceWithRoom(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "AssociateDeviceWithRoom",
	}
}
