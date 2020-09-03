// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about comments made on the comparison between two commits.
// Reaction counts might include numbers from user identities who were deleted
// after the reaction was made. For a count of reactions from active identities,
// use GetCommentReactions.
func (c *Client) GetCommentsForComparedCommit(ctx context.Context, params *GetCommentsForComparedCommitInput, optFns ...func(*Options)) (*GetCommentsForComparedCommitOutput, error) {
	stack := middleware.NewStack("GetCommentsForComparedCommit", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetCommentsForComparedCommitMiddlewares(stack)
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
	addOpGetCommentsForComparedCommitValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCommentsForComparedCommit(options.Region), middleware.Before)

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
			OperationName: "GetCommentsForComparedCommit",
			Err:           err,
		}
	}
	out := result.(*GetCommentsForComparedCommitOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCommentsForComparedCommitInput struct {
	// To establish the directionality of the comparison, the full commit ID of the
	// before commit.
	BeforeCommitId *string
	// An enumeration token that when provided in a request, returns the next batch of
	// the results.
	NextToken *string
	// A non-zero, non-negative integer used to limit the number of returned results.
	// The default is 100 comments, but you can configure up to 500.
	MaxResults *int32
	// To establish the directionality of the comparison, the full commit ID of the
	// after commit.
	AfterCommitId *string
	// The name of the repository where you want to compare commits.
	RepositoryName *string
}

type GetCommentsForComparedCommitOutput struct {
	// A list of comment objects on the compared commit.
	CommentsForComparedCommitData []*types.CommentsForComparedCommit
	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetCommentsForComparedCommitMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetCommentsForComparedCommit{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCommentsForComparedCommit{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetCommentsForComparedCommit(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "GetCommentsForComparedCommit",
	}
}
