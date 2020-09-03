// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the license configurations for the specified resource.
func (c *Client) ListLicenseSpecificationsForResource(ctx context.Context, params *ListLicenseSpecificationsForResourceInput, optFns ...func(*Options)) (*ListLicenseSpecificationsForResourceOutput, error) {
	stack := middleware.NewStack("ListLicenseSpecificationsForResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListLicenseSpecificationsForResourceMiddlewares(stack)
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
	addOpListLicenseSpecificationsForResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListLicenseSpecificationsForResource(options.Region), middleware.Before)

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
			OperationName: "ListLicenseSpecificationsForResource",
			Err:           err,
		}
	}
	out := result.(*ListLicenseSpecificationsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListLicenseSpecificationsForResourceInput struct {
	// Amazon Resource Name (ARN) of a resource that has an associated license
	// configuration.
	ResourceArn *string
	// Maximum number of results to return in a single call.
	MaxResults *int32
	// Token for the next set of results.
	NextToken *string
}

type ListLicenseSpecificationsForResourceOutput struct {
	// Token for the next set of results.
	NextToken *string
	// License configurations associated with a resource.
	LicenseSpecifications []*types.LicenseSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListLicenseSpecificationsForResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListLicenseSpecificationsForResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListLicenseSpecificationsForResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opListLicenseSpecificationsForResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "ListLicenseSpecificationsForResource",
	}
}