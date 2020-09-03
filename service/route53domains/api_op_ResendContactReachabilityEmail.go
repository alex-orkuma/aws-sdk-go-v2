// Code generated by smithy-go-codegen DO NOT EDIT.

package route53domains

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// For operations that require confirmation that the email address for the
// registrant contact is valid, such as registering a new domain, this operation
// resends the confirmation email to the current email address for the registrant
// contact.
func (c *Client) ResendContactReachabilityEmail(ctx context.Context, params *ResendContactReachabilityEmailInput, optFns ...func(*Options)) (*ResendContactReachabilityEmailOutput, error) {
	stack := middleware.NewStack("ResendContactReachabilityEmail", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpResendContactReachabilityEmailMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opResendContactReachabilityEmail(options.Region), middleware.Before)

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
			OperationName: "ResendContactReachabilityEmail",
			Err:           err,
		}
	}
	out := result.(*ResendContactReachabilityEmailOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResendContactReachabilityEmailInput struct {
	// The name of the domain for which you want Route 53 to resend a confirmation
	// email to the registrant contact.
	DomainName *string
}

type ResendContactReachabilityEmailOutput struct {
	// True if the email address for the registrant contact has already been verified,
	// and false otherwise. If the email address has already been verified, we don't
	// send another confirmation email.
	IsAlreadyVerified *bool
	// The email address for the registrant contact at the time that we sent the
	// verification email.
	EmailAddress *string
	// The domain name for which you requested a confirmation email.
	DomainName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpResendContactReachabilityEmailMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpResendContactReachabilityEmail{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpResendContactReachabilityEmail{}, middleware.After)
}

func newServiceMetadataMiddleware_opResendContactReachabilityEmail(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53domains",
		OperationName: "ResendContactReachabilityEmail",
	}
}
