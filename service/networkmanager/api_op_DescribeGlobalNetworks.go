// Code generated by smithy-go-codegen DO NOT EDIT.

package networkmanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkmanager/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more global networks. By default, all global networks are
// described. To describe the objects in your global network, you must use the
// appropriate Get* action. For example, to list the transit gateways in your
// global network, use GetTransitGatewayRegistrations.
func (c *Client) DescribeGlobalNetworks(ctx context.Context, params *DescribeGlobalNetworksInput, optFns ...func(*Options)) (*DescribeGlobalNetworksOutput, error) {
	if params == nil {
		params = &DescribeGlobalNetworksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGlobalNetworks", params, optFns, addOperationDescribeGlobalNetworksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGlobalNetworksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGlobalNetworksInput struct {

	// The IDs of one or more global networks. The maximum is 10.
	GlobalNetworkIds []*string

	// The maximum number of results to return.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type DescribeGlobalNetworksOutput struct {

	// Information about the global networks.
	GlobalNetworks []*types.GlobalNetwork

	// The token for the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeGlobalNetworksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeGlobalNetworks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeGlobalNetworks{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGlobalNetworks(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeGlobalNetworks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "networkmanager",
		OperationName: "DescribeGlobalNetworks",
	}
}
