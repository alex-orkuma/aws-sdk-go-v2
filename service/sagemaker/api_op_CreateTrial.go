// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an Amazon SageMaker trial. A trial is a set of steps called trial
// components that produce a machine learning model. A trial is part of a single
// Amazon SageMaker experiment. When you use Amazon SageMaker Studio or the Amazon
// SageMaker Python SDK, all experiments, trials, and trial components are
// automatically tracked, logged, and indexed. When you use the AWS SDK for Python
// (Boto), you must use the logging APIs provided by the SDK. You can add tags to a
// trial and then use the Search () API to search for the tags. To get a list of
// all your trials, call the ListTrials () API. To view a trial's properties, call
// the DescribeTrial () API. To create a trial component, call the
// CreateTrialComponent () API.
func (c *Client) CreateTrial(ctx context.Context, params *CreateTrialInput, optFns ...func(*Options)) (*CreateTrialOutput, error) {
	stack := middleware.NewStack("CreateTrial", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateTrialMiddlewares(stack)
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
	addOpCreateTrialValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTrial(options.Region), middleware.Before)

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
			OperationName: "CreateTrial",
			Err:           err,
		}
	}
	out := result.(*CreateTrialOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTrialInput struct {
	// The name of the trial as displayed. The name doesn't need to be unique. If
	// DisplayName isn't specified, TrialName is displayed.
	DisplayName *string
	// The name of the experiment to associate the trial with.
	ExperimentName *string
	// The name of the trial. The name must be unique in your AWS account and is not
	// case-sensitive.
	TrialName *string
	// A list of tags to associate with the trial. You can use Search () API to search
	// on the tags.
	Tags []*types.Tag
}

type CreateTrialOutput struct {
	// The Amazon Resource Name (ARN) of the trial.
	TrialArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateTrialMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateTrial{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateTrial{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTrial(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "CreateTrial",
	}
}