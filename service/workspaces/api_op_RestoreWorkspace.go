// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Restores the specified WorkSpace to its last known healthy state. You cannot
// restore a WorkSpace unless its state is  AVAILABLE, ERROR, UNHEALTHY, or
// STOPPED. Restoring a WorkSpace is a potentially destructive action that can
// result in the loss of data. For more information, see Restore a WorkSpace
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/restore-workspace.html).
// This operation is asynchronous and returns before the WorkSpace is completely
// restored.
func (c *Client) RestoreWorkspace(ctx context.Context, params *RestoreWorkspaceInput, optFns ...func(*Options)) (*RestoreWorkspaceOutput, error) {
	stack := middleware.NewStack("RestoreWorkspace", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRestoreWorkspaceMiddlewares(stack)
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
	addOpRestoreWorkspaceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRestoreWorkspace(options.Region), middleware.Before)

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
			OperationName: "RestoreWorkspace",
			Err:           err,
		}
	}
	out := result.(*RestoreWorkspaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RestoreWorkspaceInput struct {
	// The identifier of the WorkSpace.
	WorkspaceId *string
}

type RestoreWorkspaceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRestoreWorkspaceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRestoreWorkspace{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRestoreWorkspace{}, middleware.After)
}

func newServiceMetadataMiddleware_opRestoreWorkspace(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "RestoreWorkspace",
	}
}
