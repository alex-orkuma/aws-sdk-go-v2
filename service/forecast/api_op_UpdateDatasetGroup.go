// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Replaces the datasets in a dataset group with the specified datasets. The Status
// of the dataset group must be ACTIVE before you can use the dataset group to
// create a predictor. Use the DescribeDatasetGroup () operation to get the status.
func (c *Client) UpdateDatasetGroup(ctx context.Context, params *UpdateDatasetGroupInput, optFns ...func(*Options)) (*UpdateDatasetGroupOutput, error) {
	stack := middleware.NewStack("UpdateDatasetGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateDatasetGroupMiddlewares(stack)
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
	addOpUpdateDatasetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDatasetGroup(options.Region), middleware.Before)

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
			OperationName: "UpdateDatasetGroup",
			Err:           err,
		}
	}
	out := result.(*UpdateDatasetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDatasetGroupInput struct {
	// An array of the Amazon Resource Names (ARNs) of the datasets to add to the
	// dataset group.
	DatasetArns []*string
	// The ARN of the dataset group.
	DatasetGroupArn *string
}

type UpdateDatasetGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateDatasetGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateDatasetGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateDatasetGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDatasetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "UpdateDatasetGroup",
	}
}
