// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the phone configuration settings for the specified user.
func (c *Client) UpdateUserPhoneConfig(ctx context.Context, params *UpdateUserPhoneConfigInput, optFns ...func(*Options)) (*UpdateUserPhoneConfigOutput, error) {
	stack := middleware.NewStack("UpdateUserPhoneConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateUserPhoneConfigMiddlewares(stack)
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
	addOpUpdateUserPhoneConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateUserPhoneConfig(options.Region), middleware.Before)

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
			OperationName: "UpdateUserPhoneConfig",
			Err:           err,
		}
	}
	out := result.(*UpdateUserPhoneConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateUserPhoneConfigInput struct {
	// The identifier of the user account.
	UserId *string
	// Information about phone configuration settings for the user.
	PhoneConfig *types.UserPhoneConfig
	// The identifier of the Amazon Connect instance.
	InstanceId *string
}

type UpdateUserPhoneConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateUserPhoneConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateUserPhoneConfig{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateUserPhoneConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateUserPhoneConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateUserPhoneConfig",
	}
}
