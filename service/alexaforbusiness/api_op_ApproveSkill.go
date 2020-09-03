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

// Associates a skill with the organization under the customer's AWS account. If a
// skill is private, the user implicitly accepts access to this skill during
// enablement.
func (c *Client) ApproveSkill(ctx context.Context, params *ApproveSkillInput, optFns ...func(*Options)) (*ApproveSkillOutput, error) {
	stack := middleware.NewStack("ApproveSkill", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpApproveSkillMiddlewares(stack)
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
	addOpApproveSkillValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opApproveSkill(options.Region), middleware.Before)

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
			OperationName: "ApproveSkill",
			Err:           err,
		}
	}
	out := result.(*ApproveSkillOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ApproveSkillInput struct {
	// The unique identifier of the skill.
	SkillId *string
}

type ApproveSkillOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpApproveSkillMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpApproveSkill{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpApproveSkill{}, middleware.After)
}

func newServiceMetadataMiddleware_opApproveSkill(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "ApproveSkill",
	}
}
