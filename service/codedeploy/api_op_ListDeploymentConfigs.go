// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the deployment configurations with the IAM user or AWS account.
func (c *Client) ListDeploymentConfigs(ctx context.Context, params *ListDeploymentConfigsInput, optFns ...func(*Options)) (*ListDeploymentConfigsOutput, error) {
	stack := middleware.NewStack("ListDeploymentConfigs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListDeploymentConfigsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDeploymentConfigs(options.Region), middleware.Before)

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
			OperationName: "ListDeploymentConfigs",
			Err:           err,
		}
	}
	out := result.(*ListDeploymentConfigsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListDeploymentConfigs operation.
type ListDeploymentConfigsInput struct {
	// An identifier returned from the previous ListDeploymentConfigs call. It can be
	// used to return the next set of deployment configurations in the list.
	NextToken *string
}

// Represents the output of a ListDeploymentConfigs operation.
type ListDeploymentConfigsOutput struct {
	// A list of deployment configurations, including built-in configurations such as
	// CodeDeployDefault.OneAtATime.
	DeploymentConfigsList []*string
	// If a large amount of information is returned, an identifier is also returned. It
	// can be used in a subsequent list deployment configurations call to return the
	// next set of deployment configurations in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListDeploymentConfigsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListDeploymentConfigs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDeploymentConfigs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListDeploymentConfigs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "ListDeploymentConfigs",
	}
}
