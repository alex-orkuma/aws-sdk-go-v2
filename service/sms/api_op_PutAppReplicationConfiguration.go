// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates or updates a replication configuration for an application.
func (c *Client) PutAppReplicationConfiguration(ctx context.Context, params *PutAppReplicationConfigurationInput, optFns ...func(*Options)) (*PutAppReplicationConfigurationOutput, error) {
	stack := middleware.NewStack("PutAppReplicationConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutAppReplicationConfigurationMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutAppReplicationConfiguration(options.Region), middleware.Before)

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
			OperationName: "PutAppReplicationConfiguration",
			Err:           err,
		}
	}
	out := result.(*PutAppReplicationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAppReplicationConfigurationInput struct {
	// Replication configurations for server groups in the application.
	ServerGroupReplicationConfigurations []*types.ServerGroupReplicationConfiguration
	// ID of the application tassociated with the replication configuration.
	AppId *string
}

type PutAppReplicationConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutAppReplicationConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutAppReplicationConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAppReplicationConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutAppReplicationConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "PutAppReplicationConfiguration",
	}
}
