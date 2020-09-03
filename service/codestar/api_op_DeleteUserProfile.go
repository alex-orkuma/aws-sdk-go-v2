// Code generated by smithy-go-codegen DO NOT EDIT.

package codestar

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a user profile in AWS CodeStar, including all personal preference data
// associated with that profile, such as display name and email address. It does
// not delete the history of that user, for example the history of commits made by
// that user.
func (c *Client) DeleteUserProfile(ctx context.Context, params *DeleteUserProfileInput, optFns ...func(*Options)) (*DeleteUserProfileOutput, error) {
	stack := middleware.NewStack("DeleteUserProfile", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteUserProfileMiddlewares(stack)
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
	addOpDeleteUserProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUserProfile(options.Region), middleware.Before)

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
			OperationName: "DeleteUserProfile",
			Err:           err,
		}
	}
	out := result.(*DeleteUserProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteUserProfileInput struct {
	// The Amazon Resource Name (ARN) of the user to delete from AWS CodeStar.
	UserArn *string
}

type DeleteUserProfileOutput struct {
	// The Amazon Resource Name (ARN) of the user deleted from AWS CodeStar.
	UserArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteUserProfileMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteUserProfile{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteUserProfile{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteUserProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar",
		OperationName: "DeleteUserProfile",
	}
}