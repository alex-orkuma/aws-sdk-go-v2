// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a fleet's instances, including instance IDs. Use
// this action to get details on all instances in the fleet or get details on one
// specific instance. To get a specific instance, specify fleet ID and instance ID.
// To get all instances in a fleet, specify a fleet ID only. Use the pagination
// parameters to retrieve results as a set of sequential pages. If successful, an
// Instance () object is returned for each result. Learn more Remotely Access Fleet
// Instances
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-remote-access.html)Debug
// Fleet Issues
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-debug.html)
// Related operations
//
//     * DescribeInstances ()
//
//     * GetInstanceAccess ()
func (c *Client) DescribeInstances(ctx context.Context, params *DescribeInstancesInput, optFns ...func(*Options)) (*DescribeInstancesOutput, error) {
	stack := middleware.NewStack("DescribeInstances", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeInstancesMiddlewares(stack)
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
	addOpDescribeInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstances(options.Region), middleware.Before)

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
			OperationName: "DescribeInstances",
			Err:           err,
		}
	}
	out := result.(*DescribeInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type DescribeInstancesInput struct {
	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this action. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string
	// A unique identifier for a fleet to retrieve instance information for. You can
	// use either the fleet ID or ARN value.
	FleetId *string
	// A unique identifier for an instance to retrieve. Specify an instance ID or leave
	// blank to retrieve all instances in the fleet.
	InstanceId *string
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32
}

// Represents the returned data in response to a request action.
type DescribeInstancesOutput struct {
	// A collection of objects containing properties for each instance returned.
	Instances []*types.Instance
	// Token that indicates where to resume retrieving results on the next call to this
	// action. If no token is returned, these results represent the end of the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeInstancesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeInstances{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeInstances{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeInstances",
	}
}