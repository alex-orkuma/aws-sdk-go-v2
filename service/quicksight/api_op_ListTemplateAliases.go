// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the aliases of a template.
func (c *Client) ListTemplateAliases(ctx context.Context, params *ListTemplateAliasesInput, optFns ...func(*Options)) (*ListTemplateAliasesOutput, error) {
	if params == nil {
		params = &ListTemplateAliasesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTemplateAliases", params, optFns, addOperationListTemplateAliasesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTemplateAliasesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTemplateAliasesInput struct {

	// The ID of the AWS account that contains the template aliases that you're
	// listing.
	//
	// This member is required.
	AwsAccountId *string

	// The ID for the template.
	//
	// This member is required.
	TemplateId *string

	// The maximum number of results to be returned per request.
	MaxResults *int32

	// The token for the next set of results, or null if there are no more results.
	NextToken *string
}

type ListTemplateAliasesOutput struct {

	// The token for the next set of results, or null if there are no more results.
	NextToken *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// A structure containing the list of the template's aliases.
	TemplateAliasList []*types.TemplateAlias

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListTemplateAliasesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTemplateAliases{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTemplateAliases{}, middleware.After)
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
	if err = addOpListTemplateAliasesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTemplateAliases(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListTemplateAliases(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "ListTemplateAliases",
	}
}
