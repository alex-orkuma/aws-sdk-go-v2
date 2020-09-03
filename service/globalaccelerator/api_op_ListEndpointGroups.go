// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/globalaccelerator/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List the endpoint groups that are associated with a listener. To see an AWS CLI
// example of listing the endpoint groups for listener, scroll down to Example.
func (c *Client) ListEndpointGroups(ctx context.Context, params *ListEndpointGroupsInput, optFns ...func(*Options)) (*ListEndpointGroupsOutput, error) {
	stack := middleware.NewStack("ListEndpointGroups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListEndpointGroupsMiddlewares(stack)
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
	addOpListEndpointGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListEndpointGroups(options.Region), middleware.Before)

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
			OperationName: "ListEndpointGroups",
			Err:           err,
		}
	}
	out := result.(*ListEndpointGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEndpointGroupsInput struct {
	// The Amazon Resource Name (ARN) of the listener.
	ListenerArn *string
	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string
	// The number of endpoint group objects that you want to return with this call. The
	// default value is 10.
	MaxResults *int32
}

type ListEndpointGroupsOutput struct {
	// The list of the endpoint groups associated with a listener.
	EndpointGroups []*types.EndpointGroup
	// The token for the next set of results. You receive this token from a previous
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListEndpointGroupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListEndpointGroups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListEndpointGroups{}, middleware.After)
}

func newServiceMetadataMiddleware_opListEndpointGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "ListEndpointGroups",
	}
}
