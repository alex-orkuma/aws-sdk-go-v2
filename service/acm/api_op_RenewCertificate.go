// Code generated by smithy-go-codegen DO NOT EDIT.

package acm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Renews an eligable ACM certificate. At this time, only exported private
// certificates can be renewed with this operation. In order to renew your ACM PCA
// certificates with ACM, you must first grant the ACM service principal permission
// to do so
// (https://docs.aws.amazon.com/acm-pca/latest/userguide/PcaPermissions.html). For
// more information, see Testing Managed Renewal
// (https://docs.aws.amazon.com/acm/latest/userguide/manual-renewal.html) in the
// ACM User Guide.
func (c *Client) RenewCertificate(ctx context.Context, params *RenewCertificateInput, optFns ...func(*Options)) (*RenewCertificateOutput, error) {
	stack := middleware.NewStack("RenewCertificate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRenewCertificateMiddlewares(stack)
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
	addOpRenewCertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRenewCertificate(options.Region), middleware.Before)

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
			OperationName: "RenewCertificate",
			Err:           err,
		}
	}
	out := result.(*RenewCertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RenewCertificateInput struct {
	// String that contains the ARN of the ACM certificate to be renewed. This must be
	// of the form:
	// arn:aws:acm:region:123456789012:certificate/12345678-1234-1234-1234-123456789012
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html).
	CertificateArn *string
}

type RenewCertificateOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRenewCertificateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRenewCertificate{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRenewCertificate{}, middleware.After)
}

func newServiceMetadataMiddleware_opRenewCertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm",
		OperationName: "RenewCertificate",
	}
}