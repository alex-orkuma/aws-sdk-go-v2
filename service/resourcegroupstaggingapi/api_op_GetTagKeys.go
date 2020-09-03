// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns all tag keys in the specified Region for the AWS account.
func (c *Client) GetTagKeys(ctx context.Context, params *GetTagKeysInput, optFns ...func(*Options)) (*GetTagKeysOutput, error) {
	stack := middleware.NewStack("GetTagKeys", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetTagKeysMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetTagKeys(options.Region), middleware.Before)

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
			OperationName: "GetTagKeys",
			Err:           err,
		}
	}
	out := result.(*GetTagKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTagKeysInput struct {
	MaxResults *int32
	// A string that indicates that additional data is available. Leave this value
	// empty for your initial request. If the response includes a PaginationToken, use
	// that string for this value to request an additional page of data.
	PaginationToken *string
}

type GetTagKeysOutput struct {
	// A string that indicates that the response contains more data than can be
	// returned in a single response. To receive additional data, specify this string
	// for the PaginationToken value in a subsequent request.
	PaginationToken *string
	// A list of all tag keys in the AWS account.
	TagKeys []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetTagKeysMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetTagKeys{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTagKeys{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetTagKeys(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tagging",
		OperationName: "GetTagKeys",
	}
}
