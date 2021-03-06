// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The ListWorkersBlocks operation retrieves a list of Workers who are blocked from
// working on your HITs.
func (c *Client) ListWorkerBlocks(ctx context.Context, params *ListWorkerBlocksInput, optFns ...func(*Options)) (*ListWorkerBlocksOutput, error) {
	if params == nil {
		params = &ListWorkerBlocksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkerBlocks", params, optFns, addOperationListWorkerBlocksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkerBlocksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkerBlocksInput struct {
	MaxResults *int32

	// Pagination token
	NextToken *string
}

type ListWorkerBlocksOutput struct {

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The number of assignments on the page in the filtered results list, equivalent
	// to the number of assignments returned by this call.
	NumResults *int32

	// The list of WorkerBlocks, containing the collection of Worker IDs and reasons
	// for blocking.
	WorkerBlocks []*types.WorkerBlock

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListWorkerBlocksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWorkerBlocks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWorkerBlocks{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkerBlocks(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opListWorkerBlocks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ListWorkerBlocks",
	}
}
