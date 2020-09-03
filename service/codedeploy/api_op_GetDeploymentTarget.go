// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a deployment target.
func (c *Client) GetDeploymentTarget(ctx context.Context, params *GetDeploymentTargetInput, optFns ...func(*Options)) (*GetDeploymentTargetOutput, error) {
	stack := middleware.NewStack("GetDeploymentTarget", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetDeploymentTargetMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeploymentTarget(options.Region), middleware.Before)

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
			OperationName: "GetDeploymentTarget",
			Err:           err,
		}
	}
	out := result.(*GetDeploymentTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeploymentTargetInput struct {
	// The unique ID of a deployment target.
	TargetId *string
	// The unique ID of a deployment.
	DeploymentId *string
}

type GetDeploymentTargetOutput struct {
	// A deployment target that contains information about a deployment such as its
	// status, lifecycle events, and when it was last updated. It also contains
	// metadata about the deployment target. The deployment target metadata depends on
	// the deployment target's type (instanceTarget, lambdaTarget, or ecsTarget).
	DeploymentTarget *types.DeploymentTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetDeploymentTargetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetDeploymentTarget{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDeploymentTarget{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDeploymentTarget(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "GetDeploymentTarget",
	}
}
