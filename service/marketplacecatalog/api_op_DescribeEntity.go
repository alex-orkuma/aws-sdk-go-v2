// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacecatalog

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the metadata and content of the entity.
func (c *Client) DescribeEntity(ctx context.Context, params *DescribeEntityInput, optFns ...func(*Options)) (*DescribeEntityOutput, error) {
	stack := middleware.NewStack("DescribeEntity", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeEntityMiddlewares(stack)
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
	addOpDescribeEntityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEntity(options.Region), middleware.Before)

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
			OperationName: "DescribeEntity",
			Err:           err,
		}
	}
	out := result.(*DescribeEntityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEntityInput struct {
	// Required. The catalog related to the request. Fixed value: AWSMarketplace
	Catalog *string
	// Required. The unique ID of the entity to describe.
	EntityId *string
}

type DescribeEntityOutput struct {
	// The last modified date of the entity, in ISO 8601 format (2018-02-27T13:45:22Z).
	LastModifiedDate *string
	// The named type of the entity, in the format of EntityType@Version.
	EntityType *string
	// The identifier of the entity, in the format of EntityId@RevisionId.
	EntityIdentifier *string
	// This stringified JSON object includes the details of the entity.
	Details *string
	// The ARN associated to the unique identifier for the change set referenced in
	// this request.
	EntityArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeEntityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeEntity{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeEntity{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeEntity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "aws-marketplace",
		OperationName: "DescribeEntity",
	}
}
