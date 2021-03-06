// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the security policies that are attached to your file transfer
// protocol-enabled servers.
func (c *Client) ListSecurityPolicies(ctx context.Context, params *ListSecurityPoliciesInput, optFns ...func(*Options)) (*ListSecurityPoliciesOutput, error) {
	if params == nil {
		params = &ListSecurityPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSecurityPolicies", params, optFns, addOperationListSecurityPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSecurityPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSecurityPoliciesInput struct {

	// Specifies the number of security policies to return as a response to the
	// ListSecurityPolicies query.
	MaxResults *int32

	// When additional results are obtained from the ListSecurityPolicies command, a
	// NextToken parameter is returned in the output. You can then pass the NextToken
	// parameter in a subsequent command to continue listing additional security
	// policies.
	NextToken *string
}

type ListSecurityPoliciesOutput struct {

	// An array of security policies that were listed.
	//
	// This member is required.
	SecurityPolicyNames []*string

	// When you can get additional results from the ListSecurityPolicies operation, a
	// NextToken parameter is returned in the output. In a following command, you can
	// pass in the NextToken parameter to continue listing security policies.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSecurityPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSecurityPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSecurityPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSecurityPolicies(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListSecurityPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "ListSecurityPolicies",
	}
}
