// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Replaces the contents of the description of a pull request.
func (c *Client) UpdatePullRequestDescription(ctx context.Context, params *UpdatePullRequestDescriptionInput, optFns ...func(*Options)) (*UpdatePullRequestDescriptionOutput, error) {
	stack := middleware.NewStack("UpdatePullRequestDescription", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdatePullRequestDescriptionMiddlewares(stack)
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
	addOpUpdatePullRequestDescriptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePullRequestDescription(options.Region), middleware.Before)

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
			OperationName: "UpdatePullRequestDescription",
			Err:           err,
		}
	}
	out := result.(*UpdatePullRequestDescriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePullRequestDescriptionInput struct {
	// The updated content of the description for the pull request. This content
	// replaces the existing description.
	Description *string
	// The system-generated ID of the pull request. To get this ID, use
	// ListPullRequests ().
	PullRequestId *string
}

type UpdatePullRequestDescriptionOutput struct {
	// Information about the updated pull request.
	PullRequest *types.PullRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdatePullRequestDescriptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdatePullRequestDescription{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdatePullRequestDescription{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdatePullRequestDescription(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "UpdatePullRequestDescription",
	}
}
