// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves table statistics of columns.
func (c *Client) DeleteColumnStatisticsForTable(ctx context.Context, params *DeleteColumnStatisticsForTableInput, optFns ...func(*Options)) (*DeleteColumnStatisticsForTableOutput, error) {
	stack := middleware.NewStack("DeleteColumnStatisticsForTable", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteColumnStatisticsForTableMiddlewares(stack)
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
	addOpDeleteColumnStatisticsForTableValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteColumnStatisticsForTable(options.Region), middleware.Before)

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
			OperationName: "DeleteColumnStatisticsForTable",
			Err:           err,
		}
	}
	out := result.(*DeleteColumnStatisticsForTableOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteColumnStatisticsForTableInput struct {
	// The name of the partitions' table.
	TableName *string
	// The name of the column.
	ColumnName *string
	// The ID of the Data Catalog where the partitions in question reside. If none is
	// supplied, the AWS account ID is used by default.
	CatalogId *string
	// The name of the catalog database where the partitions reside.
	DatabaseName *string
}

type DeleteColumnStatisticsForTableOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteColumnStatisticsForTableMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteColumnStatisticsForTable{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteColumnStatisticsForTable{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteColumnStatisticsForTable(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "DeleteColumnStatisticsForTable",
	}
}
