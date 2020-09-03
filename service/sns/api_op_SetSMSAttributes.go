// Code generated by smithy-go-codegen DO NOT EDIT.

package sns

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Use this request to set the default settings for sending SMS messages and
// receiving daily SMS usage reports. You can override some of these settings for a
// single message when you use the Publish action with the
// MessageAttributes.entry.N parameter. For more information, see Sending an SMS
// Message (https://docs.aws.amazon.com/sns/latest/dg/sms_publish-to-phone.html) in
// the Amazon SNS Developer Guide.
func (c *Client) SetSMSAttributes(ctx context.Context, params *SetSMSAttributesInput, optFns ...func(*Options)) (*SetSMSAttributesOutput, error) {
	stack := middleware.NewStack("SetSMSAttributes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSetSMSAttributesMiddlewares(stack)
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
	addOpSetSMSAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetSMSAttributes(options.Region), middleware.Before)

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
			OperationName: "SetSMSAttributes",
			Err:           err,
		}
	}
	out := result.(*SetSMSAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the SetSMSAttributes action.
type SetSMSAttributesInput struct {
	// The default settings for sending SMS messages from your account. You can set
	// values for the following attribute names: MonthlySpendLimit – The maximum amount
	// in USD that you are willing to spend each month to send SMS messages. When
	// Amazon SNS determines that sending an SMS message would incur a cost that
	// exceeds this limit, it stops sending SMS messages within minutes. Amazon SNS
	// stops sending SMS messages within minutes of the limit being crossed. During
	// that interval, if you continue to send SMS messages, you will incur costs that
	// exceed your limit. By default, the spend limit is set to the maximum allowed by
	// Amazon SNS. If you want to raise the limit, submit an SNS Limit Increase case
	// (https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-sns).
	// For New limit value, enter your desired monthly spend limit. In the Use Case
	// Description field, explain that you are requesting an SMS monthly spend limit
	// increase. DeliveryStatusIAMRole – The ARN of the IAM role that allows Amazon SNS
	// to write logs about SMS deliveries in CloudWatch Logs. For each SMS message that
	// you send, Amazon SNS writes a log that includes the message price, the success
	// or failure status, the reason for failure (if the message failed), the message
	// dwell time, and other information. DeliveryStatusSuccessSamplingRate – The
	// percentage of successful SMS deliveries for which Amazon SNS will write logs in
	// CloudWatch Logs. The value can be an integer from 0 - 100. For example, to write
	// logs only for failed deliveries, set this value to 0. To write logs for 10% of
	// your successful deliveries, set it to 10. DefaultSenderID – A string, such as
	// your business brand, that is displayed as the sender on the receiving device.
	// Support for sender IDs varies by country. The sender ID can be 1 - 11
	// alphanumeric characters, and it must contain at least one letter. DefaultSMSType
	// – The type of SMS message that you will send by default. You can assign the
	// following values:
	//
	//     * Promotional – (Default) Noncritical messages, such as
	// marketing messages. Amazon SNS optimizes the message delivery to incur the
	// lowest cost.
	//
	//     * Transactional – Critical messages that support customer
	// transactions, such as one-time passcodes for multi-factor authentication. Amazon
	// SNS optimizes the message delivery to achieve the highest
	// reliability.
	//
	// UsageReportS3Bucket – The name of the Amazon S3 bucket to receive
	// daily SMS usage reports from Amazon SNS. Each day, Amazon SNS will deliver a
	// usage report as a CSV file to the bucket. The report includes the following
	// information for each SMS message that was successfully delivered by your
	// account:
	//
	//     * Time that the message was published (in UTC)
	//
	//     * Message ID
	//
	//
	// * Destination phone number
	//
	//     * Message type
	//
	//     * Delivery status
	//
	//     *
	// Message price (in USD)
	//
	//     * Part number (a message is split into multiple
	// parts if it is too long for a single message)
	//
	//     * Total number of parts
	//
	// To
	// receive the report, the bucket must have a policy that allows the Amazon SNS
	// service principle to perform the s3:PutObject and s3:GetBucketLocation actions.
	// For an example bucket policy and usage report, see Monitoring SMS Activity
	// (https://docs.aws.amazon.com/sns/latest/dg/sms_stats.html) in the Amazon SNS
	// Developer Guide.
	Attributes map[string]*string
}

// The response for the SetSMSAttributes action.
type SetSMSAttributesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSetSMSAttributesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSetSMSAttributes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSetSMSAttributes{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetSMSAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sns",
		OperationName: "SetSMSAttributes",
	}
}
