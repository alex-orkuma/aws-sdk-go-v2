// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Specifies which devices and operating systems users can use to access their
// WorkSpaces. For more information, see  Control Device Access
// (https://docs.aws.amazon.com/workspaces/latest/adminguide/update-directory-details.html#control-device-access).
func (c *Client) ModifyWorkspaceAccessProperties(ctx context.Context, params *ModifyWorkspaceAccessPropertiesInput, optFns ...func(*Options)) (*ModifyWorkspaceAccessPropertiesOutput, error) {
	stack := middleware.NewStack("ModifyWorkspaceAccessProperties", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpModifyWorkspaceAccessPropertiesMiddlewares(stack)
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
	addOpModifyWorkspaceAccessPropertiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyWorkspaceAccessProperties(options.Region), middleware.Before)

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
			OperationName: "ModifyWorkspaceAccessProperties",
			Err:           err,
		}
	}
	out := result.(*ModifyWorkspaceAccessPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyWorkspaceAccessPropertiesInput struct {
	// The identifier of the directory.
	ResourceId *string
	// The device types and operating systems to enable or disable for access.
	WorkspaceAccessProperties *types.WorkspaceAccessProperties
}

type ModifyWorkspaceAccessPropertiesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpModifyWorkspaceAccessPropertiesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpModifyWorkspaceAccessProperties{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyWorkspaceAccessProperties{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyWorkspaceAccessProperties(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "ModifyWorkspaceAccessProperties",
	}
}
