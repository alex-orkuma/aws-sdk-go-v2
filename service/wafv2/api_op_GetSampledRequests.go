// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from the
// prior release, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). Gets
// detailed information about a specified number of requests--a sample--that AWS
// WAF randomly selects from among the first 5,000 requests that your AWS resource
// received during a time range that you choose. You can specify a sample size of
// up to 500 requests, and you can specify any time range in the previous three
// hours. GetSampledRequests returns a time range, which is usually the time range
// that you specified. However, if your resource (such as a CloudFront
// distribution) received 5,000 requests before the specified time range elapsed,
// GetSampledRequests returns an updated time range. This new time range indicates
// the actual period during which AWS WAF selected the requests in the sample.
func (c *Client) GetSampledRequests(ctx context.Context, params *GetSampledRequestsInput, optFns ...func(*Options)) (*GetSampledRequestsOutput, error) {
	stack := middleware.NewStack("GetSampledRequests", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetSampledRequestsMiddlewares(stack)
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
	addOpGetSampledRequestsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSampledRequests(options.Region), middleware.Before)

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
			OperationName: "GetSampledRequests",
			Err:           err,
		}
	}
	out := result.(*GetSampledRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSampledRequestsInput struct {
	// The Amazon resource name (ARN) of the WebACL for which you want a sample of
	// requests.
	WebAclArn *string
	// The start date and time and the end date and time of the range for which you
	// want GetSampledRequests to return a sample of requests. You must specify the
	// times in Coordinated Universal Time (UTC) format. UTC format includes the
	// special designator, Z. For example, "2016-09-27T14:50Z". You can specify any
	// time range in the previous three hours.
	TimeWindow *types.TimeWindow
	// The number of requests that you want AWS WAF to return from among the first
	// 5,000 requests that your AWS resource received during the time range. If your
	// resource received fewer requests than the value of MaxItems, GetSampledRequests
	// returns information about all of them.
	MaxItems *int64
	// The metric name assigned to the Rule or RuleGroup for which you want a sample of
	// requests.
	RuleMetricName *string
	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB) or
	// an API Gateway stage. To work with CloudFront, you must also specify the Region
	// US East (N. Virginia) as follows:
	//
	//     * CLI - Specify the Region when you use
	// the CloudFront scope: --scope=CLOUDFRONT --region=us-east-1.
	//
	//     * API and SDKs
	// - For all calls, use the Region endpoint us-east-1.
	Scope types.Scope
}

type GetSampledRequestsOutput struct {
	// Usually, TimeWindow is the time range that you specified in the
	// GetSampledRequests request. However, if your AWS resource received more than
	// 5,000 requests during the time range that you specified in the request,
	// GetSampledRequests returns the time range for the first 5,000 requests. Times
	// are in Coordinated Universal Time (UTC) format.
	TimeWindow *types.TimeWindow
	// The total number of requests from which GetSampledRequests got a sample of
	// MaxItems requests. If PopulationSize is less than MaxItems, the sample includes
	// every request that your AWS resource received during the specified time range.
	PopulationSize *int64
	// A complex type that contains detailed information about each of the requests in
	// the sample.
	SampledRequests []*types.SampledHTTPRequest

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetSampledRequestsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetSampledRequests{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSampledRequests{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSampledRequests(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "GetSampledRequests",
	}
}
