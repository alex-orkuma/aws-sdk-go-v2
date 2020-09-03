// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Imports the root certificate of a certificate authority (CA) used to obtain TLS
// certificates used by associated websites within the company network.
func (c *Client) AssociateWebsiteCertificateAuthority(ctx context.Context, params *AssociateWebsiteCertificateAuthorityInput, optFns ...func(*Options)) (*AssociateWebsiteCertificateAuthorityOutput, error) {
	stack := middleware.NewStack("AssociateWebsiteCertificateAuthority", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAssociateWebsiteCertificateAuthorityMiddlewares(stack)
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
	addOpAssociateWebsiteCertificateAuthorityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateWebsiteCertificateAuthority(options.Region), middleware.Before)

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
			OperationName: "AssociateWebsiteCertificateAuthority",
			Err:           err,
		}
	}
	out := result.(*AssociateWebsiteCertificateAuthorityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateWebsiteCertificateAuthorityInput struct {
	// The root certificate of the CA.
	Certificate *string
	// The ARN of the fleet.
	FleetArn *string
	// The certificate name to display.
	DisplayName *string
}

type AssociateWebsiteCertificateAuthorityOutput struct {
	// A unique identifier for the CA.
	WebsiteCaId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAssociateWebsiteCertificateAuthorityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAssociateWebsiteCertificateAuthority{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateWebsiteCertificateAuthority{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateWebsiteCertificateAuthority(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "AssociateWebsiteCertificateAuthority",
	}
}
