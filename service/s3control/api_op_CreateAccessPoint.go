// Code generated by smithy-go-codegen DO NOT EDIT.

package s3control

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3control/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an access point and associates it with the specified bucket.
func (c *Client) CreateAccessPoint(ctx context.Context, params *CreateAccessPointInput, optFns ...func(*Options)) (*CreateAccessPointOutput, error) {
	stack := middleware.NewStack("CreateAccessPoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpCreateAccessPointMiddlewares(stack)
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
	addOpCreateAccessPointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccessPoint(options.Region), middleware.Before)

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
			OperationName: "CreateAccessPoint",
			Err:           err,
		}
	}
	out := result.(*CreateAccessPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessPointInput struct {
	// The name of the bucket that you want to associate this access point with.
	Bucket *string
	// The name you want to assign to this access point.
	Name *string
	// The PublicAccessBlock configuration that you want to apply to this Amazon S3
	// bucket. You can enable the configuration options in any combination. For more
	// information about when Amazon S3 considers a bucket or object public, see The
	// Meaning of "Public"
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status)
	// in the Amazon Simple Storage Service Developer Guide.
	PublicAccessBlockConfiguration *types.PublicAccessBlockConfiguration
	// The AWS account ID for the owner of the bucket for which you want to create an
	// access point.
	AccountId *string
	// If you include this field, Amazon S3 restricts access to this access point to
	// requests from the specified virtual private cloud (VPC).
	VpcConfiguration *types.VpcConfiguration
}

type CreateAccessPointOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpCreateAccessPointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpCreateAccessPoint{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpCreateAccessPoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateAccessPoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "CreateAccessPoint",
	}
}
