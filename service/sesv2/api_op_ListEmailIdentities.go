// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of all of the email identities that are associated with your AWS
// account. An identity can be either an email address or a domain. This operation
// returns identities that are verified as well as those that aren't. This
// operation returns identities that are associated with Amazon SES and Amazon
// Pinpoint.
func (c *Client) ListEmailIdentities(ctx context.Context, params *ListEmailIdentitiesInput, optFns ...func(*Options)) (*ListEmailIdentitiesOutput, error) {
	stack := middleware.NewStack("ListEmailIdentities", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListEmailIdentitiesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListEmailIdentities(options.Region), middleware.Before)

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
			OperationName: "ListEmailIdentities",
			Err:           err,
		}
	}
	out := result.(*ListEmailIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to list all of the email identities associated with your AWS account.
// This list includes identities that you've already verified, identities that are
// unverified, and identities that were verified in the past, but are no longer
// verified.
type ListEmailIdentitiesInput struct {
	// A token returned from a previous call to ListEmailIdentities to indicate the
	// position in the list of identities.
	NextToken *string
	// The number of results to show in a single call to ListEmailIdentities. If the
	// number of results is larger than the number you specified in this parameter,
	// then the response includes a NextToken element, which you can use to obtain
	// additional results. The value you specify has to be at least 0, and can be no
	// more than 1000.
	PageSize *int32
}

// A list of all of the identities that you've attempted to verify, regardless of
// whether or not those identities were successfully verified.
type ListEmailIdentitiesOutput struct {
	// An array that includes all of the email identities associated with your AWS
	// account.
	EmailIdentities []*types.IdentityInfo
	// A token that indicates that there are additional configuration sets to list. To
	// view additional configuration sets, issue another request to
	// ListEmailIdentities, and pass this token in the NextToken parameter.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListEmailIdentitiesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListEmailIdentities{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListEmailIdentities{}, middleware.After)
}

func newServiceMetadataMiddleware_opListEmailIdentities(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListEmailIdentities",
	}
}
