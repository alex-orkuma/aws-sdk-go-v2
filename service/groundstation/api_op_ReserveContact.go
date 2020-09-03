// Code generated by smithy-go-codegen DO NOT EDIT.

package groundstation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Reserves a contact using specified parameters.
func (c *Client) ReserveContact(ctx context.Context, params *ReserveContactInput, optFns ...func(*Options)) (*ReserveContactOutput, error) {
	stack := middleware.NewStack("ReserveContact", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpReserveContactMiddlewares(stack)
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
	addOpReserveContactValidationMiddleware(stack)

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
			OperationName: "ReserveContact",
			Err:           err,
		}
	}
	out := result.(*ReserveContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ReserveContactInput struct {
	// End time of a contact.
	EndTime *time.Time
	// Name of a ground station.
	GroundStation *string
	// ARN of a mission profile.
	MissionProfileArn *string
	// ARN of a satellite
	SatelliteArn *string
	// Start time of a contact.
	StartTime *time.Time
	// Tags assigned to a contact.
	Tags map[string]*string
}

//
type ReserveContactOutput struct {
	// UUID of a contact.
	ContactId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpReserveContactMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpReserveContact{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpReserveContact{}, middleware.After)
}
