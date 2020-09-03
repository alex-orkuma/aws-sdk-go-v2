// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a custom insight in Security Hub. An insight is a consolidation of
// findings that relate to a security issue that requires attention or remediation.
// To group the related findings in the insight, use the GroupByAttribute.
func (c *Client) CreateInsight(ctx context.Context, params *CreateInsightInput, optFns ...func(*Options)) (*CreateInsightOutput, error) {
	stack := middleware.NewStack("CreateInsight", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateInsightMiddlewares(stack)
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
	addOpCreateInsightValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateInsight(options.Region), middleware.Before)

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
			OperationName: "CreateInsight",
			Err:           err,
		}
	}
	out := result.(*CreateInsightOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateInsightInput struct {
	// The name of the custom insight to create.
	Name *string
	// The attribute used to group the findings for the insight. The grouping attribute
	// identifies the type of item that the insight applies to. For example, if an
	// insight is grouped by resource identifier, then the insight produces a list of
	// resource identifiers.
	GroupByAttribute *string
	// One or more attributes used to filter the findings included in the insight. The
	// insight only includes findings that match the criteria defined in the filters.
	Filters *types.AwsSecurityFindingFilters
}

type CreateInsightOutput struct {
	// The ARN of the insight created.
	InsightArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateInsightMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateInsight{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateInsight{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateInsight(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "CreateInsight",
	}
}
