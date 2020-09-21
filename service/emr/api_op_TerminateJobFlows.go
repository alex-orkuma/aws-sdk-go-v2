// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// TerminateJobFlows shuts a list of clusters (job flows) down. When a job flow is
// shut down, any step not yet completed is canceled and the EC2 instances on which
// the cluster is running are stopped. Any log files not already saved are uploaded
// to Amazon S3 if a LogUri was specified when the cluster was created. The maximum
// number of clusters allowed is 10. The call to TerminateJobFlows is asynchronous.
// Depending on the configuration of the cluster, it may take up to 1-5 minutes for
// the cluster to completely terminate and release allocated resources, such as
// Amazon EC2 instances.
func (c *Client) TerminateJobFlows(ctx context.Context, params *TerminateJobFlowsInput, optFns ...func(*Options)) (*TerminateJobFlowsOutput, error) {
	stack := middleware.NewStack("TerminateJobFlows", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpTerminateJobFlowsMiddlewares(stack)
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
	addOpTerminateJobFlowsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTerminateJobFlows(options.Region), middleware.Before)

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
			OperationName: "TerminateJobFlows",
			Err:           err,
		}
	}
	out := result.(*TerminateJobFlowsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the TerminateJobFlows () operation.
type TerminateJobFlowsInput struct {
	// A list of job flows to be shutdown.
	JobFlowIds []*string
}

type TerminateJobFlowsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpTerminateJobFlowsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpTerminateJobFlows{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpTerminateJobFlows{}, middleware.After)
}

func newServiceMetadataMiddleware_opTerminateJobFlows(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "TerminateJobFlows",
	}
}