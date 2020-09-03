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

// Returns a list of the requester's subscriptions. Each call returns a limited
// list of subscriptions, up to 100. If there are more subscriptions, a NextToken
// is also returned. Use the NextToken parameter in a new ListSubscriptions call to
// get further results. This action is throttled at 30 transactions per second
// (TPS).
func (c *Client) ListSubscriptions(ctx context.Context, params *ListSubscriptionsInput, optFns ...func(*Options)) (*ListSubscriptionsOutput, error) {
	stack := middleware.NewStack("ListSubscriptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListSubscriptionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSubscriptions(options.Region), middleware.Before)

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
			OperationName: "ListSubscriptions",
			Err:           err,
		}
	}
	out := result.(*ListSubscriptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for ListSubscriptions action.
type ListSubscriptionsInput struct {
	// Token returned by the previous ListSubscriptions request.
	NextToken *string
}

// Response for ListSubscriptions action
type ListSubscriptionsOutput struct {
	// Token to pass along to the next ListSubscriptions request. This element is
	// returned if there are more subscriptions to retrieve.
	NextToken *string
	// A list of subscriptions.
	Subscriptions []*types.Subscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListSubscriptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListSubscriptions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListSubscriptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSubscriptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "ListSubscriptions",
	}
}
