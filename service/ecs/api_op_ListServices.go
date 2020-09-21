// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the services that are running in a specified cluster.
func (c *Client) ListServices(ctx context.Context, params *ListServicesInput, optFns ...func(*Options)) (*ListServicesOutput, error) {
	stack := middleware.NewStack("ListServices", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListServicesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListServices(options.Region), middleware.Before)

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
			OperationName: "ListServices",
			Err:           err,
		}
	}
	out := result.(*ListServicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServicesInput struct {
	// The maximum number of service results returned by ListServices in paginated
	// output. When this parameter is used, ListServices only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another ListServices
	// request with the returned nextToken value. This value can be between 1 and 100.
	// If this parameter is not used, then ListServices returns up to 10 results and a
	// nextToken value if applicable.
	MaxResults *int32
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// services to list. If you do not specify a cluster, the default cluster is
	// assumed.
	Cluster *string
	// The nextToken value returned from a ListServices request indicating that more
	// results are available to fulfill the request and further calls will be needed.
	// If maxResults was provided, it is possible the number of results to be fewer
	// than maxResults. This token should be treated as an opaque identifier that is
	// only used to retrieve the next items in a list and not for other programmatic
	// purposes.
	NextToken *string
	// The scheduling strategy for services to list.
	SchedulingStrategy types.SchedulingStrategy
	// The launch type for the services to list.
	LaunchType types.LaunchType
}

type ListServicesOutput struct {
	// The nextToken value to include in a future ListServices request. When the
	// results of a ListServices request exceed maxResults, this value can be used to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string
	// The list of full ARN entries for each service associated with the specified
	// cluster.
	ServiceArns []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListServicesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListServices{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListServices{}, middleware.After)
}

func newServiceMetadataMiddleware_opListServices(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "ListServices",
	}
}