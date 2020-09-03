// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a transcription job. To see the status of the job,
// check the TranscriptionJobStatus field. If the status is COMPLETED, the job is
// finished and you can find the results at the location specified in the
// TranscriptFileUri field. If you enable content redaction, the redacted
// transcript appears in RedactedTranscriptFileUri.
func (c *Client) GetTranscriptionJob(ctx context.Context, params *GetTranscriptionJobInput, optFns ...func(*Options)) (*GetTranscriptionJobOutput, error) {
	stack := middleware.NewStack("GetTranscriptionJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetTranscriptionJobMiddlewares(stack)
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
	addOpGetTranscriptionJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTranscriptionJob(options.Region), middleware.Before)

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
			OperationName: "GetTranscriptionJob",
			Err:           err,
		}
	}
	out := result.(*GetTranscriptionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTranscriptionJobInput struct {
	// The name of the job.
	TranscriptionJobName *string
}

type GetTranscriptionJobOutput struct {
	// An object that contains the results of the transcription job.
	TranscriptionJob *types.TranscriptionJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetTranscriptionJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetTranscriptionJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTranscriptionJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTranscriptionJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "GetTranscriptionJob",
	}
}
