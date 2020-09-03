// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information on VPC peering connections. Use this operation to get
// peering information for all fleets or for one specific fleet ID. To retrieve
// connection information, call this operation from the AWS account that is used to
// manage the Amazon GameLift fleets. Specify a fleet ID or leave the parameter
// empty to retrieve all connection records. If successful, the retrieved
// information includes both active and pending connections. Active connections
// identify the IpV4 CIDR block that the VPC uses to connect.
//
//     *
// CreateVpcPeeringAuthorization ()
//
//     * DescribeVpcPeeringAuthorizations ()
//
//
// * DeleteVpcPeeringAuthorization ()
//
//     * CreateVpcPeeringConnection ()
//
//     *
// DescribeVpcPeeringConnections ()
//
//     * DeleteVpcPeeringConnection ()
func (c *Client) DescribeVpcPeeringConnections(ctx context.Context, params *DescribeVpcPeeringConnectionsInput, optFns ...func(*Options)) (*DescribeVpcPeeringConnectionsOutput, error) {
	stack := middleware.NewStack("DescribeVpcPeeringConnections", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeVpcPeeringConnectionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcPeeringConnections(options.Region), middleware.Before)

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
			OperationName: "DescribeVpcPeeringConnections",
			Err:           err,
		}
	}
	out := result.(*DescribeVpcPeeringConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type DescribeVpcPeeringConnectionsInput struct {
	// A unique identifier for a fleet. You can use either the fleet ID or ARN value.
	FleetId *string
}

// Represents the returned data in response to a request action.
type DescribeVpcPeeringConnectionsOutput struct {
	// A collection of VPC peering connection records that match the request.
	VpcPeeringConnections []*types.VpcPeeringConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeVpcPeeringConnectionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeVpcPeeringConnections{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeVpcPeeringConnections{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeVpcPeeringConnections(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeVpcPeeringConnections",
	}
}
