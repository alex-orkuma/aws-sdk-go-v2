// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about a specified partition.
func (c *Client) GetPartition(ctx context.Context, params *GetPartitionInput, optFns ...func(*Options)) (*GetPartitionOutput, error) {
	stack := middleware.NewStack("GetPartition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetPartitionMiddlewares(stack)
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
	addOpGetPartitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPartition(options.Region), middleware.Before)

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
			OperationName: "GetPartition",
			Err:           err,
		}
	}
	out := result.(*GetPartitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPartitionInput struct {
	// The values that define the partition.
	PartitionValues []*string
	// The ID of the Data Catalog where the partition in question resides. If none is
	// provided, the AWS account ID is used by default.
	CatalogId *string
	// The name of the catalog database where the partition resides.
	DatabaseName *string
	// The name of the partition's table.
	TableName *string
}

type GetPartitionOutput struct {
	// The requested information, in the form of a Partition object.
	Partition *types.Partition

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetPartitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetPartition{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPartition{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetPartition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "GetPartition",
	}
}
