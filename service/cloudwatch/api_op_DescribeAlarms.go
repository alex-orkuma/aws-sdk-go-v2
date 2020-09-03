// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the specified alarms. You can filter the results by specifying a a
// prefix for the alarm name, the alarm state, or a prefix for any action.
func (c *Client) DescribeAlarms(ctx context.Context, params *DescribeAlarmsInput, optFns ...func(*Options)) (*DescribeAlarmsOutput, error) {
	stack := middleware.NewStack("DescribeAlarms", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeAlarmsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAlarms(options.Region), middleware.Before)

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
			OperationName: "DescribeAlarms",
			Err:           err,
		}
	}
	out := result.(*DescribeAlarmsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAlarmsInput struct {
	// Use this parameter to specify whether you want the operation to return metric
	// alarms or composite alarms. If you omit this parameter, only metric alarms are
	// returned.
	AlarmTypes []types.AlarmType
	// The token returned by a previous call to indicate that there is more data
	// available.
	NextToken *string
	// Specify this parameter to receive information only about alarms that are
	// currently in the state that you specify.
	StateValue types.StateValue
	// The names of the alarms to retrieve information about.
	AlarmNames []*string
	// If you use this parameter and specify the name of a metric or composite alarm,
	// the operation returns information about the "parent" alarms of the alarm you
	// specify. These are the composite alarms that have AlarmRule parameters that
	// reference the alarm named in ParentsOfAlarmName. Information about the alarm
	// that you specify in ParentsOfAlarmName is not returned. If you specify
	// ParentsOfAlarmName, you cannot specify any other parameters in the request
	// except for MaxRecords and NextToken. If you do so, you receive a validation
	// error. Only the Alarm Name and ARN are returned by this operation when you use
	// this parameter. To get complete information about these alarms, perform another
	// DescribeAlarms operation and specify the parent alarm names in the AlarmNames
	// parameter.
	ParentsOfAlarmName *string
	// Use this parameter to filter the results of the operation to only those alarms
	// that use a certain alarm action. For example, you could specify the ARN of an
	// SNS topic to find all alarms that send notifications to that topic.
	ActionPrefix *string
	// The maximum number of alarm descriptions to retrieve.
	MaxRecords *int32
	// If you use this parameter and specify the name of a composite alarm, the
	// operation returns information about the "children" alarms of the alarm you
	// specify. These are the metric alarms and composite alarms referenced in the
	// AlarmRule field of the composite alarm that you specify in ChildrenOfAlarmName.
	// Information about the composite alarm that you name in ChildrenOfAlarmName is
	// not returned. If you specify ChildrenOfAlarmName, you cannot specify any other
	// parameters in the request except for MaxRecords and NextToken. If you do so, you
	// receive a validation error. Only the Alarm Name, ARN, StateValue
	// (OK/ALARM/INSUFFICIENT_DATA), and StateUpdatedTimestamp information are returned
	// by this operation when you use this parameter. To get complete information about
	// these alarms, perform another DescribeAlarms operation and specify the parent
	// alarm names in the AlarmNames parameter.
	ChildrenOfAlarmName *string
	// An alarm name prefix. If you specify this parameter, you receive information
	// about all alarms that have names that start with this prefix. If this parameter
	// is specified, you cannot specify AlarmNames.
	AlarmNamePrefix *string
}

type DescribeAlarmsOutput struct {
	// The token that marks the start of the next batch of returned results.
	NextToken *string
	// The information about any metric alarms returned by the operation.
	MetricAlarms []*types.MetricAlarm
	// The information about any composite alarms returned by the operation.
	CompositeAlarms []*types.CompositeAlarm

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeAlarmsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeAlarms{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeAlarms{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeAlarms(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "monitoring",
		OperationName: "DescribeAlarms",
	}
}
