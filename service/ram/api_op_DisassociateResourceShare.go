// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates the specified principals or resources from the specified resource
// share.
func (c *Client) DisassociateResourceShare(ctx context.Context, params *DisassociateResourceShareInput, optFns ...func(*Options)) (*DisassociateResourceShareOutput, error) {
	stack := middleware.NewStack("DisassociateResourceShare", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDisassociateResourceShareMiddlewares(stack)
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
	addOpDisassociateResourceShareValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateResourceShare(options.Region), middleware.Before)

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
			OperationName: "DisassociateResourceShare",
			Err:           err,
		}
	}
	out := result.(*DisassociateResourceShareOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateResourceShareInput struct {
	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string
	// The principals.
	Principals []*string
	// The Amazon Resource Names (ARNs) of the resources.
	ResourceArns []*string
	// The Amazon Resource Name (ARN) of the resource share.
	ResourceShareArn *string
}

type DisassociateResourceShareOutput struct {
	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request.
	ClientToken *string
	// Information about the associations.
	ResourceShareAssociations []*types.ResourceShareAssociation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDisassociateResourceShareMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateResourceShare{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateResourceShare{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateResourceShare(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "DisassociateResourceShare",
	}
}
