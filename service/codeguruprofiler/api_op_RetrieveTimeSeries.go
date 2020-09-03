// Code generated by smithy-go-codegen DO NOT EDIT.

package codeguruprofiler

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeguruprofiler/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

func (c *Client) RetrieveTimeSeries(ctx context.Context, params *RetrieveTimeSeriesInput, optFns ...func(*Options)) (*RetrieveTimeSeriesOutput, error) {
	stack := middleware.NewStack("RetrieveTimeSeries", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRetrieveTimeSeriesMiddlewares(stack)
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
	addOpRetrieveTimeSeriesValidationMiddleware(stack)

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
			OperationName: "RetrieveTimeSeries",
			Err:           err,
		}
	}
	out := result.(*RetrieveTimeSeriesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RetrieveTimeSeriesInput struct {
	EndTime            *time.Time
	FrameMetrics       []*types.FrameMetric
	Period             *string
	ProfilingGroupName *string
	StartTime          *time.Time
	TargetResolution   types.AggregationPeriod
}

type RetrieveTimeSeriesOutput struct {
	Data                [][]*float64
	EndTime             *time.Time
	EndTimes            []*time.Time
	FrameMetrics        []*types.FrameMetric
	Resolution          types.AggregationPeriod
	StartTime           *time.Time
	UnprocessedEndTimes map[string][]*time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRetrieveTimeSeriesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRetrieveTimeSeries{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRetrieveTimeSeries{}, middleware.After)
}
