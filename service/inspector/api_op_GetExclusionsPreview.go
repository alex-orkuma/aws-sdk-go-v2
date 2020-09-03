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

// Retrieves the exclusions preview (a list of ExclusionPreview objects) specified
// by the preview token. You can obtain the preview token by running the
// CreateExclusionsPreview API.
func (c *Client) GetExclusionsPreview(ctx context.Context, params *GetExclusionsPreviewInput, optFns ...func(*Options)) (*GetExclusionsPreviewOutput, error) {
	stack := middleware.NewStack("GetExclusionsPreview", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetExclusionsPreviewMiddlewares(stack)
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
	addOpGetExclusionsPreviewValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetExclusionsPreview(options.Region), middleware.Before)

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
			OperationName: "GetExclusionsPreview",
			Err:           err,
		}
	}
	out := result.(*GetExclusionsPreviewOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetExclusionsPreviewInput struct {
	// The ARN that specifies the assessment template for which the exclusions preview
	// was requested.
	AssessmentTemplateArn *string
	// The locale into which you want to translate the exclusion's title, description,
	// and recommendation.
	Locale types.Locale
	// The unique identifier associated of the exclusions preview.
	PreviewToken *string
	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the GetExclusionsPreviewRequest action.
	// Subsequent calls to the action fill nextToken in the request with the value of
	// nextToken from the previous response to continue listing data.
	NextToken *string
	// You can use this parameter to indicate the maximum number of items you want in
	// the response. The default value is 100. The maximum value is 500.
	MaxResults *int32
}

type GetExclusionsPreviewOutput struct {
	// Information about the exclusions included in the preview.
	ExclusionPreviews []*types.ExclusionPreview
	// When a response is generated, if there is more data to be listed, this
	// parameters is present in the response and contains the value to use for the
	// nextToken parameter in a subsequent pagination request. If there is no more data
	// to be listed, this parameter is set to null.
	NextToken *string
	// Specifies the status of the request to generate an exclusions preview.
	PreviewStatus types.PreviewStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetExclusionsPreviewMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetExclusionsPreview{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetExclusionsPreview{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetExclusionsPreview(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "GetExclusionsPreview",
	}
}
