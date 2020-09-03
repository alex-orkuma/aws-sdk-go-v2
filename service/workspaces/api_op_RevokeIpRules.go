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

// Removes one or more rules from the specified IP access control group.
func (c *Client) RevokeIpRules(ctx context.Context, params *RevokeIpRulesInput, optFns ...func(*Options)) (*RevokeIpRulesOutput, error) {
	stack := middleware.NewStack("RevokeIpRules", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRevokeIpRulesMiddlewares(stack)
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
	addOpRevokeIpRulesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRevokeIpRules(options.Region), middleware.Before)

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
			OperationName: "RevokeIpRules",
			Err:           err,
		}
	}
	out := result.(*RevokeIpRulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RevokeIpRulesInput struct {
	// The identifier of the group.
	GroupId *string
	// The rules to remove from the group.
	UserRules []*string
}

type RevokeIpRulesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRevokeIpRulesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRevokeIpRules{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRevokeIpRules{}, middleware.After)
}

func newServiceMetadataMiddleware_opRevokeIpRules(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "RevokeIpRules",
	}
}
