// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a previously created webhook by name. Deleting the webhook stops AWS
// CodePipeline from starting a pipeline every time an external event occurs. The
// API returns successfully when trying to delete a webhook that is already
// deleted. If a deleted webhook is re-created by calling PutWebhook with the same
// name, it will have a different URL.
func (c *Client) DeleteWebhook(ctx context.Context, params *DeleteWebhookInput, optFns ...func(*Options)) (*DeleteWebhookOutput, error) {
	stack := middleware.NewStack("DeleteWebhook", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteWebhookMiddlewares(stack)
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
	addOpDeleteWebhookValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteWebhook(options.Region), middleware.Before)

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
			OperationName: "DeleteWebhook",
			Err:           err,
		}
	}
	out := result.(*DeleteWebhookOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteWebhookInput struct {
	// The name of the webhook you want to delete.
	Name *string
}

type DeleteWebhookOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteWebhookMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteWebhook{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteWebhook{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteWebhook(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "DeleteWebhook",
	}
}
