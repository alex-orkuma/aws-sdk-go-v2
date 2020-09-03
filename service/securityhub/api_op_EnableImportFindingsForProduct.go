// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables the integration of a partner product with Security Hub. Integrated
// products send findings to Security Hub. When you enable a product integration, a
// permissions policy that grants permission for the product to send findings to
// Security Hub is applied.
func (c *Client) EnableImportFindingsForProduct(ctx context.Context, params *EnableImportFindingsForProductInput, optFns ...func(*Options)) (*EnableImportFindingsForProductOutput, error) {
	stack := middleware.NewStack("EnableImportFindingsForProduct", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpEnableImportFindingsForProductMiddlewares(stack)
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
	addOpEnableImportFindingsForProductValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableImportFindingsForProduct(options.Region), middleware.Before)

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
			OperationName: "EnableImportFindingsForProduct",
			Err:           err,
		}
	}
	out := result.(*EnableImportFindingsForProductOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableImportFindingsForProductInput struct {
	// The ARN of the product to enable the integration for.
	ProductArn *string
}

type EnableImportFindingsForProductOutput struct {
	// The ARN of your subscription to the product to enable integrations for.
	ProductSubscriptionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpEnableImportFindingsForProductMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpEnableImportFindingsForProduct{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpEnableImportFindingsForProduct{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableImportFindingsForProduct(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "EnableImportFindingsForProduct",
	}
}
