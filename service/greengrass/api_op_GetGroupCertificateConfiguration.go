// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the current configuration for the CA used by the group.
func (c *Client) GetGroupCertificateConfiguration(ctx context.Context, params *GetGroupCertificateConfigurationInput, optFns ...func(*Options)) (*GetGroupCertificateConfigurationOutput, error) {
	stack := middleware.NewStack("GetGroupCertificateConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetGroupCertificateConfigurationMiddlewares(stack)
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
	addOpGetGroupCertificateConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetGroupCertificateConfiguration(options.Region), middleware.Before)

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
			OperationName: "GetGroupCertificateConfiguration",
			Err:           err,
		}
	}
	out := result.(*GetGroupCertificateConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGroupCertificateConfigurationInput struct {
	// The ID of the Greengrass group.
	GroupId *string
}

type GetGroupCertificateConfigurationOutput struct {
	// The amount of time remaining before the certificate expires, in milliseconds.
	CertificateExpiryInMilliseconds *string
	// The amount of time remaining before the certificate authority expires, in
	// milliseconds.
	CertificateAuthorityExpiryInMilliseconds *string
	// The ID of the group certificate configuration.
	GroupId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetGroupCertificateConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetGroupCertificateConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetGroupCertificateConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetGroupCertificateConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "GetGroupCertificateConfiguration",
	}
}
