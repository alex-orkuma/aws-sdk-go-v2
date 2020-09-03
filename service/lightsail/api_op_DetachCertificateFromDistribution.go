// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detaches an SSL/TLS certificate from your Amazon Lightsail content delivery
// network (CDN) distribution.  <p>After the certificate is detached, your
// distribution stops accepting traffic for all of the domains that are associated
// with the certificate.</p>
func (c *Client) DetachCertificateFromDistribution(ctx context.Context, params *DetachCertificateFromDistributionInput, optFns ...func(*Options)) (*DetachCertificateFromDistributionOutput, error) {
	stack := middleware.NewStack("DetachCertificateFromDistribution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDetachCertificateFromDistributionMiddlewares(stack)
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
	addOpDetachCertificateFromDistributionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachCertificateFromDistribution(options.Region), middleware.Before)

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
			OperationName: "DetachCertificateFromDistribution",
			Err:           err,
		}
	}
	out := result.(*DetachCertificateFromDistributionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachCertificateFromDistributionInput struct {
	// The name of the distribution from which to detach the certificate.  <p>Use the
	// <code>GetDistributions</code> action to get a list of distribution names that
	// you can specify.</p>
	DistributionName *string
}

type DetachCertificateFromDistributionOutput struct {
	// An object that describes the result of the action, such as the status of the
	// request, the timestamp of the request, and the resources affected by the
	// request.
	Operation *types.Operation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDetachCertificateFromDistributionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDetachCertificateFromDistribution{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetachCertificateFromDistribution{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetachCertificateFromDistribution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "DetachCertificateFromDistribution",
	}
}
