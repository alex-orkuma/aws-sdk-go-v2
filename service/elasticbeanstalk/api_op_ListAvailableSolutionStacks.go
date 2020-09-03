// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the available solution stack names, with the public version
// first and then in reverse chronological order.
func (c *Client) ListAvailableSolutionStacks(ctx context.Context, params *ListAvailableSolutionStacksInput, optFns ...func(*Options)) (*ListAvailableSolutionStacksOutput, error) {
	stack := middleware.NewStack("ListAvailableSolutionStacks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListAvailableSolutionStacksMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAvailableSolutionStacks(options.Region), middleware.Before)

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
			OperationName: "ListAvailableSolutionStacks",
			Err:           err,
		}
	}
	out := result.(*ListAvailableSolutionStacksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAvailableSolutionStacksInput struct {
}

// A list of available AWS Elastic Beanstalk solution stacks.
type ListAvailableSolutionStacksOutput struct {
	// A list of available solution stacks and their SolutionStackDescription ().
	SolutionStackDetails []*types.SolutionStackDescription
	// A list of available solution stacks.
	SolutionStacks []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListAvailableSolutionStacksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListAvailableSolutionStacks{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListAvailableSolutionStacks{}, middleware.After)
}

func newServiceMetadataMiddleware_opListAvailableSolutionStacks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "ListAvailableSolutionStacks",
	}
}