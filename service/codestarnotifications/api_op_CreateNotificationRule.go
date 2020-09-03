// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarnotifications

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a notification rule for a resource. The rule specifies the events you
// want notifications about and the targets (such as SNS topics) where you want to
// receive them.
func (c *Client) CreateNotificationRule(ctx context.Context, params *CreateNotificationRuleInput, optFns ...func(*Options)) (*CreateNotificationRuleOutput, error) {
	stack := middleware.NewStack("CreateNotificationRule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateNotificationRuleMiddlewares(stack)
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
	addIdempotencyToken_opCreateNotificationRuleMiddleware(stack, options)
	addOpCreateNotificationRuleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateNotificationRule(options.Region), middleware.Before)

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
			OperationName: "CreateNotificationRule",
			Err:           err,
		}
	}
	out := result.(*CreateNotificationRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateNotificationRuleInput struct {
	// A list of tags to apply to this notification rule. Key names cannot start with
	// "aws".
	Tags map[string]*string
	// A list of event types associated with this notification rule. For a list of
	// allowed events, see EventTypeSummary ().
	EventTypeIds []*string
	// A unique, client-generated idempotency token that, when provided in a request,
	// ensures the request cannot be repeated with a changed parameter. If a request
	// with the same parameters is received and a token is included, the request
	// returns information about the initial request that used that token. The AWS SDKs
	// prepopulate client request tokens. If you are using an AWS SDK, an idempotency
	// token is created for you.
	ClientRequestToken *string
	// The level of detail to include in the notifications for this resource. BASIC
	// will include only the contents of the event as it would appear in AWS
	// CloudWatch. FULL will include any supplemental information provided by AWS
	// CodeStar Notifications and/or the service for the resource for which the
	// notification is created.
	DetailType types.DetailType
	// The Amazon Resource Name (ARN) of the resource to associate with the
	// notification rule. Supported resources include pipelines in AWS CodePipeline,
	// repositories in AWS CodeCommit, and build projects in AWS CodeBuild.
	Resource *string
	// The status of the notification rule. The default value is ENABLED. If the status
	// is set to DISABLED, notifications aren't sent for the notification rule.
	Status types.NotificationRuleStatus
	// A list of Amazon Resource Names (ARNs) of SNS topics to associate with the
	// notification rule.
	Targets []*types.Target
	// The name for the notification rule. Notifictaion rule names must be unique in
	// your AWS account.
	Name *string
}

type CreateNotificationRuleOutput struct {
	// The Amazon Resource Name (ARN) of the notification rule.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateNotificationRuleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateNotificationRule{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateNotificationRule{}, middleware.After)
}

type idempotencyToken_initializeOpCreateNotificationRule struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateNotificationRule) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateNotificationRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateNotificationRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateNotificationRuleInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateNotificationRuleMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateNotificationRule{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateNotificationRule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-notifications",
		OperationName: "CreateNotificationRule",
	}
}