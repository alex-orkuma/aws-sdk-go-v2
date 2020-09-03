// Code generated by smithy-go-codegen DO NOT EDIT.

package guardduty

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves an Amazon GuardDuty detector specified by the detectorId.
func (c *Client) GetDetector(ctx context.Context, params *GetDetectorInput, optFns ...func(*Options)) (*GetDetectorOutput, error) {
	stack := middleware.NewStack("GetDetector", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDetectorMiddlewares(stack)
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
	addOpGetDetectorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDetector(options.Region), middleware.Before)

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
			OperationName: "GetDetector",
			Err:           err,
		}
	}
	out := result.(*GetDetectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDetectorInput struct {
	// The unique ID of the detector that you want to get.
	DetectorId *string
}

type GetDetectorOutput struct {
	// The last-updated timestamp for the detector.
	UpdatedAt *string
	// The timestamp of when the detector was created.
	CreatedAt *string
	// The tags of the detector resource.
	Tags map[string]*string
	// An object that describes which data sources are enabled for the detector.
	DataSources *types.DataSourceConfigurationsResult
	// The GuardDuty service role.
	ServiceRole *string
	// The detector status.
	Status types.DetectorStatus
	// The publishing frequency of the finding.
	FindingPublishingFrequency types.FindingPublishingFrequency

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDetectorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDetector{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDetector{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDetector(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "guardduty",
		OperationName: "GetDetector",
	}
}
