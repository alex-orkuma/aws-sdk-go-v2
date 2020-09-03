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

// Gets a list of InferRxNorm jobs that you have submitted.
func (c *Client) ListRxNormInferenceJobs(ctx context.Context, params *ListRxNormInferenceJobsInput, optFns ...func(*Options)) (*ListRxNormInferenceJobsOutput, error) {
	stack := middleware.NewStack("ListRxNormInferenceJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListRxNormInferenceJobsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRxNormInferenceJobs(options.Region), middleware.Before)

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
			OperationName: "ListRxNormInferenceJobs",
			Err:           err,
		}
	}
	out := result.(*ListRxNormInferenceJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRxNormInferenceJobsInput struct {
	// Filters the jobs that are returned. You can filter jobs based on their names,
	// status, or the date and time that they were submitted. You can only set one
	// filter at a time.
	Filter *types.ComprehendMedicalAsyncJobFilter
	// Identifies the next page of results to return.
	MaxResults *int32
	// Identifies the next page of results to return.
	NextToken *string
}

type ListRxNormInferenceJobsOutput struct {
	// Identifies the next page of results to return.
	NextToken *string
	// The maximum number of results to return in each page. The default is 100.
	ComprehendMedicalAsyncJobPropertiesList []*types.ComprehendMedicalAsyncJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListRxNormInferenceJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListRxNormInferenceJobs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRxNormInferenceJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListRxNormInferenceJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehendmedical",
		OperationName: "ListRxNormInferenceJobs",
	}
}
