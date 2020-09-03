// Code generated by smithy-go-codegen DO NOT EDIT.

package codestar

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a project, including project resources. Does not delete users associated
// with the project, but does delete the IAM roles that allowed access to the
// project.
func (c *Client) DeleteProject(ctx context.Context, params *DeleteProjectInput, optFns ...func(*Options)) (*DeleteProjectOutput, error) {
	stack := middleware.NewStack("DeleteProject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteProjectMiddlewares(stack)
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
	addOpDeleteProjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteProject(options.Region), middleware.Before)

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
			OperationName: "DeleteProject",
			Err:           err,
		}
	}
	out := result.(*DeleteProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteProjectInput struct {
	// Whether to send a delete request for the primary stack in AWS CloudFormation
	// originally used to generate the project and its resources. This option will
	// delete all AWS resources for the project (except for any buckets in Amazon S3)
	// as well as deleting the project itself. Recommended for most use cases.
	DeleteStack *bool
	// The ID of the project to be deleted in AWS CodeStar.
	Id *string
	// A user- or system-generated token that identifies the entity that requested
	// project deletion. This token can be used to repeat the request.
	ClientRequestToken *string
}

type DeleteProjectOutput struct {
	// The ID of the primary stack in AWS CloudFormation that will be deleted as part
	// of deleting the project and its resources.
	StackId *string
	// The Amazon Resource Name (ARN) of the deleted project.
	ProjectArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteProjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteProject{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteProject{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteProject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar",
		OperationName: "DeleteProject",
	}
}
