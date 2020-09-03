// Code generated by smithy-go-codegen DO NOT EDIT.

package mq

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mq/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describe available engine types and versions.
func (c *Client) DescribeBrokerEngineTypes(ctx context.Context, params *DescribeBrokerEngineTypesInput, optFns ...func(*Options)) (*DescribeBrokerEngineTypesOutput, error) {
	stack := middleware.NewStack("DescribeBrokerEngineTypes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeBrokerEngineTypesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBrokerEngineTypes(options.Region), middleware.Before)

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
			OperationName: "DescribeBrokerEngineTypes",
			Err:           err,
		}
	}
	out := result.(*DescribeBrokerEngineTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBrokerEngineTypesInput struct {
	// The maximum number of engine types that Amazon MQ can return per page (20 by
	// default). This value must be an integer from 5 to 100.
	MaxResults *int32
	// The token that specifies the next page of results Amazon MQ should return. To
	// request the first page, leave nextToken empty.
	NextToken *string
	// Filter response by engine type.
	EngineType *string
}

type DescribeBrokerEngineTypesOutput struct {
	// The token that specifies the next page of results Amazon MQ should return. To
	// request the first page, leave nextToken empty.
	NextToken *string
	// List of available engine types and versions.
	BrokerEngineTypes []*types.BrokerEngineType
	// Required. The maximum number of engine types that can be returned per page (20
	// by default). This value must be an integer from 5 to 100.
	MaxResults *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeBrokerEngineTypesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeBrokerEngineTypes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeBrokerEngineTypes{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeBrokerEngineTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mq",
		OperationName: "DescribeBrokerEngineTypes",
	}
}
