// Code generated by smithy-go-codegen DO NOT EDIT.

package rdsdata

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rdsdata/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Runs a batch SQL statement over an array of data. You can run bulk update and
// insert operations for multiple records using a DML statement with different
// parameter sets. Bulk operations can provide a significant performance
// improvement over individual insert and update operations. If a call isn't part
// of a transaction because it doesn't include the transactionID parameter, changes
// that result from the call are committed automatically.
func (c *Client) BatchExecuteStatement(ctx context.Context, params *BatchExecuteStatementInput, optFns ...func(*Options)) (*BatchExecuteStatementOutput, error) {
	stack := middleware.NewStack("BatchExecuteStatement", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpBatchExecuteStatementMiddlewares(stack)
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
	addOpBatchExecuteStatementValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchExecuteStatement(options.Region), middleware.Before)

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
			OperationName: "BatchExecuteStatement",
			Err:           err,
		}
	}
	out := result.(*BatchExecuteStatementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request parameters represent the input of a SQL statement over an array of
// data.
type BatchExecuteStatementInput struct {
	// The name of the database.
	Database *string
	// The parameter set for the batch operation. The SQL statement is executed as many
	// times as the number of parameter sets provided. To execute a SQL statement with
	// no parameters, use one of the following options:
	//
	//     * Specify one or more
	// empty parameter sets.
	//
	//     * Use the ExecuteStatement operation instead of the
	// BatchExecuteStatement operation.
	//
	// Array parameters are not supported.
	ParameterSets [][]*types.SqlParameter
	// The Amazon Resource Name (ARN) of the Aurora Serverless DB cluster.
	ResourceArn *string
	// The name of the database schema.
	Schema *string
	// The name or ARN of the secret that enables access to the DB cluster.
	SecretArn *string
	// The SQL statement to run.
	Sql *string
	// The identifier of a transaction that was started by using the BeginTransaction
	// operation. Specify the transaction ID of the transaction that you want to
	// include the SQL statement in. If the SQL statement is not part of a transaction,
	// don't set this parameter.
	TransactionId *string
}

// The response elements represent the output of a SQL statement over an array of
// data.
type BatchExecuteStatementOutput struct {
	// The execution results of each batch entry.
	UpdateResults []*types.UpdateResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpBatchExecuteStatementMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpBatchExecuteStatement{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchExecuteStatement{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchExecuteStatement(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "",
		OperationName: "BatchExecuteStatement",
	}
}
