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

// Updates port settings for a fleet. To update settings, specify the fleet ID to
// be updated and list the permissions you want to update. List the permissions you
// want to add in InboundPermissionAuthorizations, and permissions you want to
// remove in InboundPermissionRevocations. Permissions to be removed must match
// existing fleet permissions. If successful, the fleet ID for the updated fleet is
// returned. Learn more Setting up GameLift Fleets
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html)
// Related operations
//
//     * CreateFleet ()
//
//     * ListFleets ()
//
//     * DeleteFleet
// ()
//
//     * DescribeFleetAttributes ()
//
//     * Update fleets:
//
//         *
// UpdateFleetAttributes ()
//
//         * UpdateFleetCapacity ()
//
//         *
// UpdateFleetPortSettings ()
//
//         * UpdateRuntimeConfiguration ()
//
//     *
// StartFleetActions () or StopFleetActions ()
func (c *Client) UpdateFleetPortSettings(ctx context.Context, params *UpdateFleetPortSettingsInput, optFns ...func(*Options)) (*UpdateFleetPortSettingsOutput, error) {
	stack := middleware.NewStack("UpdateFleetPortSettings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateFleetPortSettingsMiddlewares(stack)
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
	addOpUpdateFleetPortSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFleetPortSettings(options.Region), middleware.Before)

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
			OperationName: "UpdateFleetPortSettings",
			Err:           err,
		}
	}
	out := result.(*UpdateFleetPortSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type UpdateFleetPortSettingsInput struct {
	// A collection of port settings to be added to the fleet resource.
	InboundPermissionAuthorizations []*types.IpPermission
	// A collection of port settings to be removed from the fleet resource.
	InboundPermissionRevocations []*types.IpPermission
	// A unique identifier for a fleet to update port settings for. You can use either
	// the fleet ID or ARN value.
	FleetId *string
}

// Represents the returned data in response to a request action.
type UpdateFleetPortSettingsOutput struct {
	// A unique identifier for a fleet that was updated.
	FleetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateFleetPortSettingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFleetPortSettings{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFleetPortSettings{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateFleetPortSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateFleetPortSettings",
	}
}
