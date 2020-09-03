// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts the running of the version of a model. Starting a model takes a while to
// complete. To check the current state of the model, use DescribeProjectVersions
// (). Once the model is running, you can detect custom labels in new images by
// calling DetectCustomLabels (). You are charged for the amount of time that the
// model is running. To stop a running model, call StopProjectVersion (). This
// operation requires permissions to perform the rekognition:StartProjectVersion
// action.
func (c *Client) StartProjectVersion(ctx context.Context, params *StartProjectVersionInput, optFns ...func(*Options)) (*StartProjectVersionOutput, error) {
	stack := middleware.NewStack("StartProjectVersion", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartProjectVersionMiddlewares(stack)
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
	addOpStartProjectVersionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartProjectVersion(options.Region), middleware.Before)

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
			OperationName: "StartProjectVersion",
			Err:           err,
		}
	}
	out := result.(*StartProjectVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartProjectVersionInput struct {
	// The Amazon Resource Name(ARN) of the model version that you want to start.
	ProjectVersionArn *string
	// The minimum number of inference units to use. A single inference unit represents
	// 1 hour of processing and can support up to 5 Transaction Pers Second (TPS). Use
	// a higher number to increase the TPS throughput of your model. You are charged
	// for the number of inference units that you use.
	MinInferenceUnits *int32
}

type StartProjectVersionOutput struct {
	// The current running status of the model.
	Status types.ProjectVersionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartProjectVersionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartProjectVersion{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartProjectVersion{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartProjectVersion(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "StartProjectVersion",
	}
}