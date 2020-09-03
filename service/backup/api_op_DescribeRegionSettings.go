// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the current service opt-in settings for the Region. If the service has a
// value set to true, AWS Backup attempts to protect that service's resources in
// this Region, when included in an on-demand backup or scheduled backup plan. If
// the value is set to false for a service, AWS Backup does not attempt to protect
// that service's resources in this Region.
func (c *Client) DescribeRegionSettings(ctx context.Context, params *DescribeRegionSettingsInput, optFns ...func(*Options)) (*DescribeRegionSettingsOutput, error) {
	stack := middleware.NewStack("DescribeRegionSettings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeRegionSettingsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRegionSettings(options.Region), middleware.Before)

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
			OperationName: "DescribeRegionSettings",
			Err:           err,
		}
	}
	out := result.(*DescribeRegionSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRegionSettingsInput struct {
}

type DescribeRegionSettingsOutput struct {
	// Returns a list of all services along with the opt-in preferences in the region.
	ResourceTypeOptInPreference map[string]*bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeRegionSettingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRegionSettings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRegionSettings{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeRegionSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "DescribeRegionSettings",
	}
}
