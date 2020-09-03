// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of tags. You can return tags from a specific resource by
// specifying an ARN, or you can return all tags for a given type of resource, such
// as clusters, snapshots, and so on. The following are limitations for
// DescribeTags:
//
//     * You cannot specify an ARN and a resource-type value
// together in the same request.
//
//     * You cannot use the MaxRecords and Marker
// parameters together with the ARN parameter.
//
//     * The MaxRecords parameter can
// be a range from 10 to 50 results to return in a request.
//
// If you specify both
// tag keys and tag values in the same request, Amazon Redshift returns all
// resources that match any combination of the specified keys and values. For
// example, if you have owner and environment for tag keys, and admin and test for
// tag values, all resources that have any combination of those values are
// returned. If both tag keys and values are omitted from the request, resources
// are returned regardless of whether they have tag keys or values associated with
// them.
func (c *Client) DescribeTags(ctx context.Context, params *DescribeTagsInput, optFns ...func(*Options)) (*DescribeTagsOutput, error) {
	stack := middleware.NewStack("DescribeTags", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeTagsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTags(options.Region), middleware.Before)

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
			OperationName: "DescribeTags",
			Err:           err,
		}
	}
	out := result.(*DescribeTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeTagsInput struct {
	// A tag value or values for which you want to return all matching resources that
	// are associated with the specified value or values. For example, suppose that you
	// have resources tagged with values called admin and test. If you specify both of
	// these tag values in the request, Amazon Redshift returns a response with all
	// resources that have either or both of these tag values associated with them.
	TagValues []*string
	// The Amazon Resource Name (ARN) for which you want to describe the tag or tags.
	// For example, arn:aws:redshift:us-east-2:123456789:cluster:t1.
	ResourceName *string
	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the marker
	// parameter and retrying the command. If the marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string
	// The type of resource with which you want to view tags. Valid resource types
	// are:
	//
	//     * Cluster
	//
	//     * CIDR/IP
	//
	//     * EC2 security group
	//
	//     * Snapshot
	//
	//
	// * Cluster security group
	//
	//     * Subnet group
	//
	//     * HSM connection
	//
	//     * HSM
	// certificate
	//
	//     * Parameter group
	//
	//     * Snapshot copy grant
	//
	// For more
	// information about Amazon Redshift resource types and constructing ARNs, go to
	// Specifying Policy Elements: Actions, Effects, Resources, and Principals
	// (https://docs.aws.amazon.com/redshift/latest/mgmt/redshift-iam-access-control-overview.html#redshift-iam-access-control-specify-actions)
	// in the Amazon Redshift Cluster Management Guide.
	ResourceType *string
	// The maximum number or response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	MaxRecords *int32
	// A tag key or keys for which you want to return all matching resources that are
	// associated with the specified key or keys. For example, suppose that you have
	// resources tagged with keys called owner and environment. If you specify both of
	// these tag keys in the request, Amazon Redshift returns a response with all
	// resources that have either or both of these tag keys associated with them.
	TagKeys []*string
}

//
type DescribeTagsOutput struct {
	// A list of tags with their associated resources.
	TaggedResources []*types.TaggedResource
	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeTagsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeTags{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeTags{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTags(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeTags",
	}
}