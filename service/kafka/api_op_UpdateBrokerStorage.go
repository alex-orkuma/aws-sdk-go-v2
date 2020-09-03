// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kafka/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the EBS storage associated with MSK brokers.
func (c *Client) UpdateBrokerStorage(ctx context.Context, params *UpdateBrokerStorageInput, optFns ...func(*Options)) (*UpdateBrokerStorageOutput, error) {
	stack := middleware.NewStack("UpdateBrokerStorage", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateBrokerStorageMiddlewares(stack)
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
	addOpUpdateBrokerStorageValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateBrokerStorage(options.Region), middleware.Before)

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
			OperationName: "UpdateBrokerStorage",
			Err:           err,
		}
	}
	out := result.(*UpdateBrokerStorageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateBrokerStorageInput struct {
	// Describes the target volume size and the ID of the broker to apply the update
	// to.
	TargetBrokerEBSVolumeInfo []*types.BrokerEBSVolumeInfo
	// The Amazon Resource Name (ARN) that uniquely identifies the cluster.
	ClusterArn *string
	// The version of cluster to update from. A successful operation will then generate
	// a new version.
	CurrentVersion *string
}

type UpdateBrokerStorageOutput struct {
	// The Amazon Resource Name (ARN) of the cluster operation.
	ClusterOperationArn *string
	// The Amazon Resource Name (ARN) of the cluster.
	ClusterArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateBrokerStorageMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateBrokerStorage{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateBrokerStorage{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateBrokerStorage(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "UpdateBrokerStorage",
	}
}
