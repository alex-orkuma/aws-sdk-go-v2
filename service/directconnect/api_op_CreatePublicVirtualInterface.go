// Code generated by smithy-go-codegen DO NOT EDIT.

package directconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a public virtual interface. A virtual interface is the VLAN that
// transports AWS Direct Connect traffic. A public virtual interface supports
// sending traffic to public services of AWS such as Amazon S3. When creating an
// IPv6 public virtual interface (addressFamily is ipv6), leave the customer and
// amazon address fields blank to use auto-assigned IPv6 space. Custom IPv6
// addresses are not supported.
func (c *Client) CreatePublicVirtualInterface(ctx context.Context, params *CreatePublicVirtualInterfaceInput, optFns ...func(*Options)) (*CreatePublicVirtualInterfaceOutput, error) {
	stack := middleware.NewStack("CreatePublicVirtualInterface", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreatePublicVirtualInterfaceMiddlewares(stack)
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
	addOpCreatePublicVirtualInterfaceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePublicVirtualInterface(options.Region), middleware.Before)

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
			OperationName: "CreatePublicVirtualInterface",
			Err:           err,
		}
	}
	out := result.(*CreatePublicVirtualInterfaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePublicVirtualInterfaceInput struct {
	// The ID of the connection.
	ConnectionId *string
	// Information about the public virtual interface.
	NewPublicVirtualInterface *types.NewPublicVirtualInterface
}

// Information about a virtual interface.
type CreatePublicVirtualInterfaceOutput struct {
	// The ID of the connection.
	ConnectionId *string
	// The BGP peers configured on this virtual interface.
	BgpPeers []*types.BGPPeer
	// The state of the virtual interface. The following are the possible values:
	//
	//
	// * confirming: The creation of the virtual interface is pending confirmation from
	// the virtual interface owner. If the owner of the virtual interface is different
	// from the owner of the connection on which it is provisioned, then the virtual
	// interface will remain in this state until it is confirmed by the virtual
	// interface owner.
	//
	//     * verifying: This state only applies to public virtual
	// interfaces. Each public virtual interface needs validation before the virtual
	// interface can be created.
	//
	//     * pending: A virtual interface is in this state
	// from the time that it is created until the virtual interface is ready to forward
	// traffic.
	//
	//     * available: A virtual interface that is able to forward
	// traffic.
	//
	//     * down: A virtual interface that is BGP down.
	//
	//     * deleting: A
	// virtual interface is in this state immediately after calling
	// DeleteVirtualInterface () until it can no longer forward traffic.
	//
	//     *
	// deleted: A virtual interface that cannot forward traffic.
	//
	//     * rejected: The
	// virtual interface owner has declined creation of the virtual interface. If a
	// virtual interface in the Confirming state is deleted by the virtual interface
	// owner, the virtual interface enters the Rejected state.
	//
	//     * unknown: The
	// state of the virtual interface is not available.
	VirtualInterfaceState types.VirtualInterfaceState
	// The type of virtual interface. The possible values are private and public.
	VirtualInterfaceType *string
	// The routes to be advertised to the AWS network in this Region. Applies to public
	// virtual interfaces.
	RouteFilterPrefixes []*types.RouteFilterPrefix
	// The authentication key for BGP configuration. This string has a minimum length
	// of 6 characters and and a maximun lenth of 80 characters.
	AuthKey *string
	// The Direct Connect endpoint on which the virtual interface terminates.
	AwsDeviceV2 *string
	// The tags associated with the virtual interface.
	Tags []*types.Tag
	// Indicates whether jumbo frames (9001 MTU) are supported.
	JumboFrameCapable *bool
	// The ID of the AWS account that owns the virtual interface.
	OwnerAccount *string
	// The maximum transmission unit (MTU), in bytes. The supported values are 1500 and
	// 9001. The default value is 1500.
	Mtu *int32
	// The IP address assigned to the customer interface.
	CustomerAddress *string
	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string
	// The customer router configuration.
	CustomerRouterConfig *string
	// The location of the connection.
	Location *string
	// The ID of the virtual interface.
	VirtualInterfaceId *string
	// The autonomous system number (ASN) for the Amazon side of the connection.
	AmazonSideAsn *int64
	// The IP address assigned to the Amazon interface.
	AmazonAddress *string
	// The ID of the virtual private gateway. Applies only to private virtual
	// interfaces.
	VirtualGatewayId *string
	// The name of the virtual interface assigned by the customer network. The name has
	// a maximum of 100 characters. The following are valid characters: a-z, 0-9 and a
	// hyphen (-).
	VirtualInterfaceName *string
	// The ID of the VLAN.
	Vlan *int32
	// The autonomous system (AS) number for Border Gateway Protocol (BGP)
	// configuration. The valid values are 1-2147483647.
	Asn *int32
	// The address family for the BGP peer.
	AddressFamily types.AddressFamily
	// The AWS Region where the virtual interface is located.
	Region *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreatePublicVirtualInterfaceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePublicVirtualInterface{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePublicVirtualInterface{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreatePublicVirtualInterface(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "directconnect",
		OperationName: "CreatePublicVirtualInterface",
	}
}
