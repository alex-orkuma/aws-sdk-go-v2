// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpoint

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes one or more attributes, of the same attribute type, from all the
// endpoints that are associated with an application.
func (c *Client) RemoveAttributes(ctx context.Context, params *RemoveAttributesInput, optFns ...func(*Options)) (*RemoveAttributesOutput, error) {
	stack := middleware.NewStack("RemoveAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRemoveAttributesMiddlewares(stack)
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
	addOpRemoveAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveAttributes(options.Region), middleware.Before)

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
			OperationName: "RemoveAttributes",
			Err:           err,
		}
	}
	out := result.(*RemoveAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveAttributesInput struct {
	// The type of attribute or attributes to remove. Valid values are:
	//
	//     *
	// endpoint-custom-attributes - Custom attributes that describe endpoints, such as
	// the date when an associated user opted in or out of receiving communications
	// from you through a specific type of channel.
	//
	//     * endpoint-metric-attributes -
	// Custom metrics that your app reports to Amazon Pinpoint for endpoints, such as
	// the number of app sessions or the number of items left in a cart.
	//
	//     *
	// endpoint-user-attributes - Custom attributes that describe users, such as first
	// name, last name, and age.
	AttributeType *string
	// The unique identifier for the application. This identifier is displayed as the
	// Project ID on the Amazon Pinpoint console.
	ApplicationId *string
	// Specifies one or more attributes to remove from all the endpoints that are
	// associated with an application.
	UpdateAttributesRequest *types.UpdateAttributesRequest
}

type RemoveAttributesOutput struct {
	// Provides information about the type and the names of attributes that were
	// removed from all the endpoints that are associated with an application.
	AttributesResource *types.AttributesResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRemoveAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRemoveAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRemoveAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mobiletargeting",
		OperationName: "RemoveAttributes",
	}
}
