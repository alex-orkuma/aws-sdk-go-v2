// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a contact by the contact ARN.
func (c *Client) DeleteContact(ctx context.Context, params *DeleteContactInput, optFns ...func(*Options)) (*DeleteContactOutput, error) {
	stack := middleware.NewStack("DeleteContact", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteContactMiddlewares(stack)
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
	addOpDeleteContactValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteContact(options.Region), middleware.Before)

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
			OperationName: "DeleteContact",
			Err:           err,
		}
	}
	out := result.(*DeleteContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteContactInput struct {
	// The ARN of the contact to delete.
	ContactArn *string
}

type DeleteContactOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteContactMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteContact{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteContact{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteContact(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "DeleteContact",
	}
}
