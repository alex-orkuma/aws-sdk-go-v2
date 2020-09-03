// Code generated by smithy-go-codegen DO NOT EDIT.

package sso

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/service/sso/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all AWS accounts assigned to the user. These AWS accounts are assigned by
// the administrator of the account. For more information, see Assign User Access
// (https://docs.aws.amazon.com/singlesignon/latest/userguide/useraccess.html#assignusers)
// in the AWS SSO User Guide. This operation returns a paginated response.
func (c *Client) ListAccounts(ctx context.Context, params *ListAccountsInput, optFns ...func(*Options)) (*ListAccountsOutput, error) {
	stack := middleware.NewStack("ListAccounts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListAccountsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListAccountsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAccounts(options.Region), middleware.Before)

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
			OperationName: "ListAccounts",
			Err:           err,
		}
	}
	out := result.(*ListAccountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAccountsInput struct {
	// The token issued by the CreateToken API call. For more information, see
	// CreateToken
	// (https://docs.aws.amazon.com/singlesignon/latest/OIDCAPIReference/API_CreateToken.html)
	// in the AWS SSO OIDC API Reference Guide.
	AccessToken *string
	// This is the number of items clients can request per page.
	MaxResults *int32
	// (Optional) When requesting subsequent pages, this is the page token from the
	// previous response output.
	NextToken *string
}

type ListAccountsOutput struct {
	// A paginated response with the list of account information and the next token if
	// more results are available.
	AccountList []*types.AccountInfo
	// The page token client that is used to retrieve the list of accounts.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListAccountsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListAccounts{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListAccounts{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAccounts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "awsssoportal",
		OperationName: "ListAccounts",
	}
}
