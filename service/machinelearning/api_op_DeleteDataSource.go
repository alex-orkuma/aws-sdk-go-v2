// Code generated by smithy-go-codegen DO NOT EDIT.

package machinelearning

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Assigns the DELETED status to a DataSource, rendering it unusable. After using
// the DeleteDataSource operation, you can use the GetDataSource () operation to
// verify that the status of the DataSource changed to DELETED. Caution: The
// results of the DeleteDataSource operation are irreversible.
func (c *Client) DeleteDataSource(ctx context.Context, params *DeleteDataSourceInput, optFns ...func(*Options)) (*DeleteDataSourceOutput, error) {
	stack := middleware.NewStack("DeleteDataSource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteDataSourceMiddlewares(stack)
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
	addOpDeleteDataSourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteDataSource(options.Region), middleware.Before)

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
			OperationName: "DeleteDataSource",
			Err:           err,
		}
	}
	out := result.(*DeleteDataSourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteDataSourceInput struct {
	// A user-supplied ID that uniquely identifies the DataSource.
	DataSourceId *string
}

// Represents the output of a DeleteDataSource operation.
type DeleteDataSourceOutput struct {
	// A user-supplied ID that uniquely identifies the DataSource. This value should be
	// identical to the value of the DataSourceID in the request.
	DataSourceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteDataSourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteDataSource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteDataSource{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteDataSource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "machinelearning",
		OperationName: "DeleteDataSource",
	}
}
