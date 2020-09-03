// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This operation creates a job.
func (c *Client) CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) {
	stack := middleware.NewStack("CreateJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateJobMiddlewares(stack)
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
	addOpCreateJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateJob(options.Region), middleware.Before)

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
			OperationName: "CreateJob",
			Err:           err,
		}
	}
	out := result.(*CreateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request body for CreateJob.
type CreateJobInput struct {
	// The details for the CreateJob request.
	Details *types.RequestDetails
	// The type of job to be created.
	Type types.Type
}

type CreateJobOutput struct {
	// The job type.
	Type types.Type
	// Details about the job.
	Details *types.ResponseDetails
	// The errors associated with jobs.
	Errors []*types.JobError
	// The date and time that the job was created, in ISO 8601 format.
	CreatedAt *time.Time
	// The state of the job.
	State types.State
	// The unique identifier for the job.
	Id *string
	// The ARN for the job.
	Arn *string
	// The date and time that the job was last updated, in ISO 8601 format.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateJob{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dataexchange",
		OperationName: "CreateJob",
	}
}
