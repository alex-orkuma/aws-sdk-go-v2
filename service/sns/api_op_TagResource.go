// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Add tags to the specified Amazon SNS topic. For an overview, see Amazon SNS Tags
// (https://docs.aws.amazon.com/sns/latest/dg/sns-tags.html) in the Amazon SNS
// Developer Guide. When you use topic tags, keep the following guidelines in
// mind:
//
//     * Adding more than 50 tags to a topic isn't recommended.
//
//     * Tags
// don't have any semantic meaning. Amazon SNS interprets tags as character
// strings.
//
//     * Tags are case-sensitive.
//
//     * A new tag with a key identical
// to that of an existing tag overwrites the existing tag.
//
//     * Tagging actions
// are limited to 10 TPS per AWS account, per AWS region. If your application
// requires a higher throughput, file a technical support request
// (https://console.aws.amazon.com/support/home#/case/create?issueType=technical).
func (c *Client) TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) {
	stack := middleware.NewStack("TagResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpTagResourceMiddlewares(stack)
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
	addOpTagResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTagResource(options.Region), middleware.Before)

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
			OperationName: "TagResource",
			Err:           err,
		}
	}
	out := result.(*TagResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TagResourceInput struct {
	// The tags to be added to the specified topic. A tag consists of a required key
	// and an optional value.
	Tags []*types.Tag
	// The ARN of the topic to which to add tags.
	ResourceArn *string
}

type TagResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpTagResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpTagResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpTagResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opTagResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "TagResource",
	}
}
