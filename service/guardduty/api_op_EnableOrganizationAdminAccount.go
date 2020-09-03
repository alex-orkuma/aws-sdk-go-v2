// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables an AWS account within the organization as the GuardDuty delegated
// administrator.
func (c *Client) EnableOrganizationAdminAccount(ctx context.Context, params *EnableOrganizationAdminAccountInput, optFns ...func(*Options)) (*EnableOrganizationAdminAccountOutput, error) {
	stack := middleware.NewStack("EnableOrganizationAdminAccount", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpEnableOrganizationAdminAccountMiddlewares(stack)
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
	addOpEnableOrganizationAdminAccountValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableOrganizationAdminAccount(options.Region), middleware.Before)

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
			OperationName: "EnableOrganizationAdminAccount",
			Err:           err,
		}
	}
	out := result.(*EnableOrganizationAdminAccountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableOrganizationAdminAccountInput struct {
	// The AWS Account ID for the organization account to be enabled as a GuardDuty
	// delegated administrator.
	AdminAccountId *string
}

type EnableOrganizationAdminAccountOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpEnableOrganizationAdminAccountMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpEnableOrganizationAdminAccount{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpEnableOrganizationAdminAccount{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableOrganizationAdminAccount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "EnableOrganizationAdminAccount",
	}
}
