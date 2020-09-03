// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehendmedical

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehendmedical/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a list of medical entity detection jobs that you have submitted.
func (c *Client) ListEntitiesDetectionV2Jobs(ctx context.Context, params *ListEntitiesDetectionV2JobsInput, optFns ...func(*Options)) (*ListEntitiesDetectionV2JobsOutput, error) {
	stack := middleware.NewStack("ListEntitiesDetectionV2Jobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListEntitiesDetectionV2JobsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListEntitiesDetectionV2Jobs(options.Region), middleware.Before)

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
			OperationName: "ListEntitiesDetectionV2Jobs",
			Err:           err,
		}
	}
	out := result.(*ListEntitiesDetectionV2JobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEntitiesDetectionV2JobsInput struct {
	// Filters the jobs that are returned. You can filter jobs based on their names,
	// status, or the date and time that they were submitted. You can only set one
	// filter at a time.
	Filter *types.ComprehendMedicalAsyncJobFilter
	// Identifies the next page of results to return.
	NextToken *string
	// The maximum number of results to return in each page. The default is 100.
	MaxResults *int32
}

type ListEntitiesDetectionV2JobsOutput struct {
	// A list containing the properties of each job returned.
	ComprehendMedicalAsyncJobPropertiesList []*types.ComprehendMedicalAsyncJobProperties
	// Identifies the next page of results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListEntitiesDetectionV2JobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListEntitiesDetectionV2Jobs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEntitiesDetectionV2Jobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListEntitiesDetectionV2Jobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "ListEntitiesDetectionV2Jobs",
	}
}
