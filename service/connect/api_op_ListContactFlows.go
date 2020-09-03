// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides information about the contact flows for the specified Amazon Connect
// instance.
func (c *Client) ListContactFlows(ctx context.Context, params *ListContactFlowsInput, optFns ...func(*Options)) (*ListContactFlowsOutput, error) {
	stack := middleware.NewStack("ListContactFlows", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListContactFlowsMiddlewares(stack)
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
	addOpListContactFlowsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListContactFlows(options.Region), middleware.Before)

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
			OperationName: "ListContactFlows",
			Err:           err,
		}
	}
	out := result.(*ListContactFlowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListContactFlowsInput struct {
	// The type of contact flow.
	ContactFlowTypes []types.ContactFlowType
	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string
	// The identifier of the Amazon Connect instance.
	InstanceId *string
	// The maximimum number of results to return per page.
	MaxResults *int32
}

type ListContactFlowsOutput struct {
	// Information about the contact flows.
	ContactFlowSummaryList []*types.ContactFlowSummary
	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListContactFlowsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListContactFlows{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListContactFlows{}, middleware.After)
}

func newServiceMetadataMiddleware_opListContactFlows(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListContactFlows",
	}
}