// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointemail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sends an email message. You can use the Amazon Pinpoint Email API to send two
// types of messages:
//
//     * Simple – A standard email message. When you create
// this type of message, you specify the sender, the recipient, and the message
// body, and Amazon Pinpoint assembles the message for you.
//
//     * Raw – A raw,
// MIME-formatted email message. When you send this type of email, you have to
// specify all of the message headers, as well as the message body. You can use
// this message type to send messages that contain attachments. The message that
// you specify has to be a valid MIME message.
func (c *Client) SendEmail(ctx context.Context, params *SendEmailInput, optFns ...func(*Options)) (*SendEmailOutput, error) {
	stack := middleware.NewStack("SendEmail", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSendEmailMiddlewares(stack)
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
	addOpSendEmailValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSendEmail(options.Region), middleware.Before)

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
			OperationName: "SendEmail",
			Err:           err,
		}
	}
	out := result.(*SendEmailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to send an email message.
type SendEmailInput struct {
	// The "Reply-to" email addresses for the message. When the recipient replies to
	// the message, each Reply-to address receives the reply.
	ReplyToAddresses []*string
	// The name of the configuration set that you want to use when sending the email.
	ConfigurationSetName *string
	// The address that Amazon Pinpoint should send bounce and complaint notifications
	// to.
	FeedbackForwardingEmailAddress *string
	// The email address that you want to use as the "From" address for the email. The
	// address that you specify has to be verified.
	FromEmailAddress *string
	// An object that contains the recipients of the email message.
	Destination *types.Destination
	// An object that contains the body of the message. You can send either a Simple
	// message or a Raw message.
	Content *types.EmailContent
	// A list of tags, in the form of name/value pairs, to apply to an email that you
	// send using the SendEmail operation. Tags correspond to characteristics of the
	// email that you define, so that you can publish email sending events.
	EmailTags []*types.MessageTag
}

// A unique message ID that you receive when Amazon Pinpoint accepts an email for
// sending.
type SendEmailOutput struct {
	// A unique identifier for the message that is generated when Amazon Pinpoint
	// accepts the message. It is possible for Amazon Pinpoint to accept a message
	// without sending it. This can happen when the message you're trying to send has
	// an attachment doesn't pass a virus check, or when you send a templated email
	// that contains invalid personalization content, for example.
	MessageId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSendEmailMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSendEmail{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSendEmail{}, middleware.After)
}

func newServiceMetadataMiddleware_opSendEmail(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "SendEmail",
	}
}
