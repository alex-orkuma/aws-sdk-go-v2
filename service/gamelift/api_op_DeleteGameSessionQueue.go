// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a game session queue. This action means that any
// StartGameSessionPlacement () requests that reference this queue will fail. To
// delete a queue, specify the queue name. Learn more  Using Multi-Region Queues
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/queues-intro.html)
// Related operations
//
//     * CreateGameSessionQueue ()
//
//     *
// DescribeGameSessionQueues ()
//
//     * UpdateGameSessionQueue ()
//
//     *
// DeleteGameSessionQueue ()
func (c *Client) DeleteGameSessionQueue(ctx context.Context, params *DeleteGameSessionQueueInput, optFns ...func(*Options)) (*DeleteGameSessionQueueOutput, error) {
	stack := middleware.NewStack("DeleteGameSessionQueue", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteGameSessionQueueMiddlewares(stack)
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
	addOpDeleteGameSessionQueueValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteGameSessionQueue(options.Region), middleware.Before)

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
			OperationName: "DeleteGameSessionQueue",
			Err:           err,
		}
	}
	out := result.(*DeleteGameSessionQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type DeleteGameSessionQueueInput struct {
	// A descriptive label that is associated with game session queue. Queue names must
	// be unique within each Region. You can use either the queue ID or ARN value.
	Name *string
}

type DeleteGameSessionQueueOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteGameSessionQueueMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteGameSessionQueue{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteGameSessionQueue{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteGameSessionQueue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DeleteGameSessionQueue",
	}
}
