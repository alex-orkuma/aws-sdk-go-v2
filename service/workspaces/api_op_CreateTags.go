// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates the specified tags for the specified WorkSpaces resource.
func (c *Client) CreateTags(ctx context.Context, params *CreateTagsInput, optFns ...func(*Options)) (*CreateTagsOutput, error) {
	stack := middleware.NewStack("CreateTags", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateTagsMiddlewares(stack)
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
	addOpCreateTagsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTags(options.Region), middleware.Before)

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
			OperationName: "CreateTags",
			Err:           err,
		}
	}
	out := result.(*CreateTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTagsInput struct {
	// The tags. Each WorkSpaces resource can have a maximum of 50 tags. If you want to
	// add new tags to a set of existing tags, you must submit all of the existing tags
	// along with the new ones.
	Tags []*types.Tag
	// The identifier of the WorkSpaces resource. The supported resource types are
	// WorkSpaces, registered directories, images, custom bundles, and IP access
	// control groups.
	ResourceId *string
}

type CreateTagsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateTagsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTags{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTags{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTags(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "CreateTags",
	}
}
