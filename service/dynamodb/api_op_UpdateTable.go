// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the provisioned throughput settings, global secondary indexes, or
// DynamoDB Streams settings for a given table. You can only perform one of the
// following operations at once:
//
//     * Modify the provisioned throughput settings
// of the table.
//
//     * Enable or disable DynamoDB Streams on the table.
//
//     *
// Remove a global secondary index from the table.
//
//     * Create a new global
// secondary index on the table. After the index begins backfilling, you can use
// UpdateTable to perform other operations.
//
// UpdateTable is an asynchronous
// operation; while it is executing, the table status changes from ACTIVE to
// UPDATING. While it is UPDATING, you cannot issue another UpdateTable request.
// When the table returns to the ACTIVE state, the UpdateTable operation is
// complete.
func (c *Client) UpdateTable(ctx context.Context, params *UpdateTableInput, optFns ...func(*Options)) (*UpdateTableOutput, error) {
	stack := middleware.NewStack("UpdateTable", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpUpdateTableMiddlewares(stack)
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
	addOpUpdateTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateTable(options.Region), middleware.Before)

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
			OperationName: "UpdateTable",
			Err:           err,
		}
	}
	out := result.(*UpdateTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of an UpdateTable operation.
type UpdateTableInput struct {
	// The new server-side encryption settings for the specified table.
	SSESpecification *types.SSESpecification
	// The name of the table to be updated.
	TableName *string
	// An array of attributes that describe the key schema for the table and indexes.
	// If you are adding a new global secondary index to the table,
	// AttributeDefinitions must include the key element(s) of the new index.
	AttributeDefinitions []*types.AttributeDefinition
	// An array of one or more global secondary indexes for the table. For each index
	// in the array, you can request one action:
	//
	//     * Create - add a new global
	// secondary index to the table.
	//
	//     * Update - modify the provisioned throughput
	// settings of an existing global secondary index.
	//
	//     * Delete - remove a global
	// secondary index from the table.
	//
	// You can create or delete only one global
	// secondary index per UpdateTable operation. For more information, see Managing
	// Global Secondary Indexes
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.OnlineOps.html)
	// in the Amazon DynamoDB Developer Guide.
	GlobalSecondaryIndexUpdates []*types.GlobalSecondaryIndexUpdate
	// Controls how you are charged for read and write throughput and how you manage
	// capacity. When switching from pay-per-request to provisioned capacity, initial
	// provisioned capacity values must be set. The initial provisioned capacity values
	// are estimated based on the consumed read and write capacity of your table and
	// global secondary indexes over the past 30 minutes.
	//
	//     * PROVISIONED - We
	// recommend using PROVISIONED for predictable workloads. PROVISIONED sets the
	// billing mode to Provisioned Mode
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual).
	//
	//
	// * PAY_PER_REQUEST - We recommend using PAY_PER_REQUEST for unpredictable
	// workloads. PAY_PER_REQUEST sets the billing mode to On-Demand Mode
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand).
	BillingMode types.BillingMode
	// The new provisioned throughput settings for the specified table or index.
	ProvisionedThroughput *types.ProvisionedThroughput
	// A list of replica update actions (create, delete, or update) for the table. This
	// property only applies to Version 2019.11.21
	// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html)
	// of global tables.
	ReplicaUpdates []*types.ReplicationGroupUpdate
	// Represents the DynamoDB Streams configuration for the table. You receive a
	// ResourceInUseException if you try to enable a stream on a table that already has
	// a stream, or if you try to disable a stream on a table that doesn't have a
	// stream.
	StreamSpecification *types.StreamSpecification
}

// Represents the output of an UpdateTable operation.
type UpdateTableOutput struct {
	// Represents the properties of the table.
	TableDescription *types.TableDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpUpdateTableMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateTable{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateTable{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "UpdateTable",
	}
}
