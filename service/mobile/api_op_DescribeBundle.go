// Code generated by smithy-go-codegen DO NOT EDIT.

package mobile

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mobile/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Get the bundle details for the requested bundle id.
func (c *Client) DescribeBundle(ctx context.Context, params *DescribeBundleInput, optFns ...func(*Options)) (*DescribeBundleOutput, error) {
	stack := middleware.NewStack("DescribeBundle", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeBundleMiddlewares(stack)
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
	addOpDescribeBundleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBundle(options.Region), middleware.Before)

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
			OperationName: "DescribeBundle",
			Err:           err,
		}
	}
	out := result.(*DescribeBundleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Request structure to request the details of a specific bundle.
type DescribeBundleInput struct {
	// Unique bundle identifier.
	BundleId *string
}

// Result structure contains the details of the bundle.
type DescribeBundleOutput struct {
	// The details of the bundle.
	Details *types.BundleDetails

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeBundleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBundle{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBundle{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeBundle(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "awsmobilehubservice",
		OperationName: "DescribeBundle",
	}
}
