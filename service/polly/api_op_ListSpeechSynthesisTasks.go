// Code generated by smithy-go-codegen DO NOT EDIT.

package polly

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/polly/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of SpeechSynthesisTask objects ordered by their creation date.
// This operation can filter the tasks by their status, for example, allowing users
// to list only tasks that are completed.
func (c *Client) ListSpeechSynthesisTasks(ctx context.Context, params *ListSpeechSynthesisTasksInput, optFns ...func(*Options)) (*ListSpeechSynthesisTasksOutput, error) {
	stack := middleware.NewStack("ListSpeechSynthesisTasks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListSpeechSynthesisTasksMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListSpeechSynthesisTasks(options.Region), middleware.Before)

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
			OperationName: "ListSpeechSynthesisTasks",
			Err:           err,
		}
	}
	out := result.(*ListSpeechSynthesisTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSpeechSynthesisTasksInput struct {
	// Maximum number of speech synthesis tasks returned in a List operation.
	MaxResults *int32
	// The pagination token to use in the next request to continue the listing of
	// speech synthesis tasks.
	NextToken *string
	// Status of the speech synthesis tasks returned in a List operation
	Status types.TaskStatus
}

type ListSpeechSynthesisTasksOutput struct {
	// List of SynthesisTask objects that provides information from the specified task
	// in the list request, including output format, creation time, task status, and so
	// on.
	SynthesisTasks []*types.SynthesisTask
	// An opaque pagination token returned from the previous List operation in this
	// request. If present, this indicates where to continue the listing.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListSpeechSynthesisTasksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListSpeechSynthesisTasks{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListSpeechSynthesisTasks{}, middleware.After)
}

func newServiceMetadataMiddleware_opListSpeechSynthesisTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "polly",
		OperationName: "ListSpeechSynthesisTasks",
	}
}
