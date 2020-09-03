// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticloadbalancing

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified the load balancers. If no load balancers are specified,
// the call describes all of your load balancers.
func (c *Client) DescribeLoadBalancers(ctx context.Context, params *DescribeLoadBalancersInput, optFns ...func(*Options)) (*DescribeLoadBalancersOutput, error) {
	stack := middleware.NewStack("DescribeLoadBalancers", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeLoadBalancersMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoadBalancers(options.Region), middleware.Before)

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
			OperationName: "DescribeLoadBalancers",
			Err:           err,
		}
	}
	out := result.(*DescribeLoadBalancersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeLoadBalancers.
type DescribeLoadBalancersInput struct {
	// The names of the load balancers.
	LoadBalancerNames []*string
	// The marker for the next set of results. (You received this marker from a
	// previous call.)
	Marker *string
	// The maximum number of results to return with this call (a number from 1 to 400).
	// The default is 400.
	PageSize *int32
}

// Contains the parameters for DescribeLoadBalancers.
type DescribeLoadBalancersOutput struct {
	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	NextMarker *string
	// Information about the load balancers.
	LoadBalancerDescriptions []*types.LoadBalancerDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeLoadBalancersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLoadBalancers{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLoadBalancers{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeLoadBalancers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticloadbalancing",
		OperationName: "DescribeLoadBalancers",
	}
}
