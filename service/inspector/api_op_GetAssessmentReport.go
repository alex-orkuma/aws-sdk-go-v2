// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/inspector/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Produces an assessment report that includes detailed and comprehensive results
// of a specified assessment run.
func (c *Client) GetAssessmentReport(ctx context.Context, params *GetAssessmentReportInput, optFns ...func(*Options)) (*GetAssessmentReportOutput, error) {
	stack := middleware.NewStack("GetAssessmentReport", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetAssessmentReportMiddlewares(stack)
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
	addOpGetAssessmentReportValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAssessmentReport(options.Region), middleware.Before)

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
			OperationName: "GetAssessmentReport",
			Err:           err,
		}
	}
	out := result.(*GetAssessmentReportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAssessmentReportInput struct {
	// Specifies the file format (html or pdf) of the assessment report that you want
	// to generate.
	ReportFileFormat types.ReportFileFormat
	// The ARN that specifies the assessment run for which you want to generate a
	// report.
	AssessmentRunArn *string
	// Specifies the type of the assessment report that you want to generate. There are
	// two types of assessment reports: a finding report and a full report. For more
	// information, see Assessment Reports
	// (https://docs.aws.amazon.com/inspector/latest/userguide/inspector_reports.html).
	ReportType types.ReportType
}

type GetAssessmentReportOutput struct {
	// Specifies the status of the request to generate an assessment report.
	Status types.ReportStatus
	// Specifies the URL where you can find the generated assessment report. This
	// parameter is only returned if the report is successfully generated.
	Url *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetAssessmentReportMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetAssessmentReport{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAssessmentReport{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetAssessmentReport(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "GetAssessmentReport",
	}
}
