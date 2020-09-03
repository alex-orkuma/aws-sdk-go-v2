// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the backup plan that is specified by the plan ID as a backup template.
func (c *Client) ExportBackupPlanTemplate(ctx context.Context, params *ExportBackupPlanTemplateInput, optFns ...func(*Options)) (*ExportBackupPlanTemplateOutput, error) {
	stack := middleware.NewStack("ExportBackupPlanTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpExportBackupPlanTemplateMiddlewares(stack)
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
	addOpExportBackupPlanTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opExportBackupPlanTemplate(options.Region), middleware.Before)

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
			OperationName: "ExportBackupPlanTemplate",
			Err:           err,
		}
	}
	out := result.(*ExportBackupPlanTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportBackupPlanTemplateInput struct {
	// Uniquely identifies a backup plan.
	BackupPlanId *string
}

type ExportBackupPlanTemplateOutput struct {
	// The body of a backup plan template in JSON format. This is a signed JSON
	// document that cannot be modified before being passed to GetBackupPlanFromJSON.
	BackupPlanTemplateJson *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpExportBackupPlanTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpExportBackupPlanTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpExportBackupPlanTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opExportBackupPlanTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "ExportBackupPlanTemplate",
	}
}