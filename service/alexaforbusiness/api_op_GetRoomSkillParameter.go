// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets room skill parameter details by room, skill, and parameter key ARN.
func (c *Client) GetRoomSkillParameter(ctx context.Context, params *GetRoomSkillParameterInput, optFns ...func(*Options)) (*GetRoomSkillParameterOutput, error) {
	stack := middleware.NewStack("GetRoomSkillParameter", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetRoomSkillParameterMiddlewares(stack)
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
	addOpGetRoomSkillParameterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRoomSkillParameter(options.Region), middleware.Before)

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
			OperationName: "GetRoomSkillParameter",
			Err:           err,
		}
	}
	out := result.(*GetRoomSkillParameterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRoomSkillParameterInput struct {
	// The room skill parameter key for which to get details. Required.
	ParameterKey *string
	// The ARN of the skill from which to get the room skill parameter details.
	// Required.
	SkillId *string
	// The ARN of the room from which to get the room skill parameter details.
	RoomArn *string
}

type GetRoomSkillParameterOutput struct {
	// The details of the room skill parameter requested. Required.
	RoomSkillParameter *types.RoomSkillParameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetRoomSkillParameterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetRoomSkillParameter{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRoomSkillParameter{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRoomSkillParameter(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "GetRoomSkillParameter",
	}
}
