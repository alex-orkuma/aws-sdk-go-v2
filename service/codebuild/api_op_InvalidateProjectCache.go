// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resets the cache for a project.
func (c *Client) InvalidateProjectCache(ctx context.Context, params *InvalidateProjectCacheInput, optFns ...func(*Options)) (*InvalidateProjectCacheOutput, error) {
	stack := middleware.NewStack("InvalidateProjectCache", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpInvalidateProjectCacheMiddlewares(stack)
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
	addOpInvalidateProjectCacheValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInvalidateProjectCache(options.Region), middleware.Before)

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
			OperationName: "InvalidateProjectCache",
			Err:           err,
		}
	}
	out := result.(*InvalidateProjectCacheOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InvalidateProjectCacheInput struct {
	// The name of the AWS CodeBuild build project that the cache is reset for.
	ProjectName *string
}

type InvalidateProjectCacheOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpInvalidateProjectCacheMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpInvalidateProjectCache{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpInvalidateProjectCache{}, middleware.After)
}

func newServiceMetadataMiddleware_opInvalidateProjectCache(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "InvalidateProjectCache",
	}
}