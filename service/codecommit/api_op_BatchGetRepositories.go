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

// Returns information about one or more repositories. The description field for a
// repository accepts all HTML characters and all valid Unicode characters.
// Applications that do not HTML-encode the description and display it in a webpage
// can expose users to potentially malicious code. Make sure that you HTML-encode
// the description field in any application that uses this API to display the
// repository description on a webpage.
func (c *Client) BatchGetRepositories(ctx context.Context, params *BatchGetRepositoriesInput, optFns ...func(*Options)) (*BatchGetRepositoriesOutput, error) {
	stack := middleware.NewStack("BatchGetRepositories", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpBatchGetRepositoriesMiddlewares(stack)
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
	addOpBatchGetRepositoriesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchGetRepositories(options.Region), middleware.Before)

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
			OperationName: "BatchGetRepositories",
			Err:           err,
		}
	}
	out := result.(*BatchGetRepositoriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a batch get repositories operation.
type BatchGetRepositoriesInput struct {
	// The names of the repositories to get information about. The length constraint
	// limit is for each string in the array. The array itself can be empty.
	RepositoryNames []*string
}

// Represents the output of a batch get repositories operation.
type BatchGetRepositoriesOutput struct {
	// Returns a list of repository names for which information could not be found.
	RepositoriesNotFound []*string
	// A list of repositories returned by the batch get repositories operation.
	Repositories []*types.RepositoryMetadata

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpBatchGetRepositoriesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpBatchGetRepositories{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchGetRepositories{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchGetRepositories(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "BatchGetRepositories",
	}
}
