// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Represents the failure of a job as returned to the pipeline by a job worker.
// Used for custom actions only.
func (c *Client) PutJobFailureResult(ctx context.Context, params *PutJobFailureResultInput, optFns ...func(*Options)) (*PutJobFailureResultOutput, error) {
	stack := middleware.NewStack("PutJobFailureResult", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutJobFailureResultMiddlewares(stack)
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
	addOpPutJobFailureResultValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutJobFailureResult(options.Region), middleware.Before)

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
			OperationName: "PutJobFailureResult",
			Err:           err,
		}
	}
	out := result.(*PutJobFailureResultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PutJobFailureResult action.
type PutJobFailureResultInput struct {
	// The unique system-generated ID of the job that failed. This is the same ID
	// returned from PollForJobs.
	JobId *string
	// The details about the failure of a job.
	FailureDetails *types.FailureDetails
}

type PutJobFailureResultOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutJobFailureResultMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutJobFailureResult{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutJobFailureResult{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutJobFailureResult(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "PutJobFailureResult",
	}
}