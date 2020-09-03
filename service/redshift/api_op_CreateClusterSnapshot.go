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

// Creates a manual snapshot of the specified cluster. The cluster must be in the
// available state. For more information about working with snapshots, go to Amazon
// Redshift Snapshots
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-snapshots.html)
// in the Amazon Redshift Cluster Management Guide.
func (c *Client) CreateClusterSnapshot(ctx context.Context, params *CreateClusterSnapshotInput, optFns ...func(*Options)) (*CreateClusterSnapshotOutput, error) {
	stack := middleware.NewStack("CreateClusterSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpCreateClusterSnapshotMiddlewares(stack)
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
	addOpCreateClusterSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateClusterSnapshot(options.Region), middleware.Before)

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
			OperationName: "CreateClusterSnapshot",
			Err:           err,
		}
	}
	out := result.(*CreateClusterSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type CreateClusterSnapshotInput struct {
	// The number of days that a manual snapshot is retained. If the value is -1, the
	// manual snapshot is retained indefinitely.  <p>The value must be either -1 or an
	// integer between 1 and 3,653.</p> <p>The default value is -1.</p>
	ManualSnapshotRetentionPeriod *int32
	// A list of tag instances.
	Tags []*types.Tag
	// The cluster identifier for which you want a snapshot.
	ClusterIdentifier *string
	// A unique identifier for the snapshot that you are requesting. This identifier
	// must be unique for all snapshots within the AWS account. Constraints:
	//
	//     *
	// Cannot be null, empty, or blank
	//
	//     * Must contain from 1 to 255 alphanumeric
	// characters or hyphens
	//
	//     * First character must be a letter
	//
	//     * Cannot end
	// with a hyphen or contain two consecutive hyphens
	//
	// Example: my-snapshot-id
	SnapshotIdentifier *string
}

type CreateClusterSnapshotOutput struct {
	// Describes a snapshot.
	Snapshot *types.Snapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpCreateClusterSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpCreateClusterSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpCreateClusterSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateClusterSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "CreateClusterSnapshot",
	}
}
