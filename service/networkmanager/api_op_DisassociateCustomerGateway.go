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

// Disassociates a customer gateway from a device and a link.
func (c *Client) DisassociateCustomerGateway(ctx context.Context, params *DisassociateCustomerGatewayInput, optFns ...func(*Options)) (*DisassociateCustomerGatewayOutput, error) {
	stack := middleware.NewStack("DisassociateCustomerGateway", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDisassociateCustomerGatewayMiddlewares(stack)
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
	addOpDisassociateCustomerGatewayValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateCustomerGateway(options.Region), middleware.Before)

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
			OperationName: "DisassociateCustomerGateway",
			Err:           err,
		}
	}
	out := result.(*DisassociateCustomerGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateCustomerGatewayInput struct {
	// The Amazon Resource Name (ARN) of the customer gateway. For more information,
	// see Resources Defined by Amazon EC2
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/list_amazonec2.html#amazonec2-resources-for-iam-policies).
	CustomerGatewayArn *string
	// The ID of the global network.
	GlobalNetworkId *string
}

type DisassociateCustomerGatewayOutput struct {
	// Information about the customer gateway association.
	CustomerGatewayAssociation *types.CustomerGatewayAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDisassociateCustomerGatewayMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateCustomerGateway{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateCustomerGateway{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateCustomerGateway(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "DisassociateCustomerGateway",
	}
}
