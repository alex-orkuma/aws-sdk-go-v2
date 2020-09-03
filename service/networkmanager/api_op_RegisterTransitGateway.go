// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a transit gateway in your global network. The transit gateway can be
// in any AWS Region, but it must be owned by the same AWS account that owns the
// global network. You cannot register a transit gateway in more than one global
// network.
func (c *Client) RegisterTransitGateway(ctx context.Context, params *RegisterTransitGatewayInput, optFns ...func(*Options)) (*RegisterTransitGatewayOutput, error) {
	stack := middleware.NewStack("RegisterTransitGateway", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRegisterTransitGatewayMiddlewares(stack)
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
	addOpRegisterTransitGatewayValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterTransitGateway(options.Region), middleware.Before)

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
			OperationName: "RegisterTransitGateway",
			Err:           err,
		}
	}
	out := result.(*RegisterTransitGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterTransitGatewayInput struct {
	// The Amazon Resource Name (ARN) of the transit gateway. For more information, see
	// Resources Defined by Amazon EC2
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazonec2.html#amazonec2-resources-for-iam-policies).
	TransitGatewayArn *string
	// The ID of the global network.
	GlobalNetworkId *string
}

type RegisterTransitGatewayOutput struct {
	// Information about the transit gateway registration.
	TransitGatewayRegistration *types.TransitGatewayRegistration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRegisterTransitGatewayMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRegisterTransitGateway{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterTransitGateway{}, middleware.After)
}

func newServiceMetadataMiddleware_opRegisterTransitGateway(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "RegisterTransitGateway",
	}
}
