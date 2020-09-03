// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Create a new pool of dedicated IP addresses. A pool can include one or more
// dedicated IP addresses that are associated with your Amazon Pinpoint account.
// You can associate a pool with a configuration set. When you send an email that
// uses that configuration set, Amazon Pinpoint sends it using only the IP
// addresses in the associated pool.
func (c *Client) CreateDedicatedIpPool(ctx context.Context, params *CreateDedicatedIpPoolInput, optFns ...func(*Options)) (*CreateDedicatedIpPoolOutput, error) {
	stack := middleware.NewStack("CreateDedicatedIpPool", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDedicatedIpPoolMiddlewares(stack)
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
	addOpCreateDedicatedIpPoolValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDedicatedIpPool(options.Region), middleware.Before)

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
			OperationName: "CreateDedicatedIpPool",
			Err:           err,
		}
	}
	out := result.(*CreateDedicatedIpPoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to create a new dedicated IP pool.
type CreateDedicatedIpPoolInput struct {
	// The name of the dedicated IP pool.
	PoolName *string
	// An object that defines the tags (keys and values) that you want to associate
	// with the pool.
	Tags []*types.Tag
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type CreateDedicatedIpPoolOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDedicatedIpPoolMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDedicatedIpPool{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDedicatedIpPool{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDedicatedIpPool(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "CreateDedicatedIpPool",
	}
}
