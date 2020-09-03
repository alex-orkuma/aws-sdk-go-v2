// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Takes a set of configuration settings and either a configuration template or
// environment, and determines whether those values are valid. This action returns
// a list of messages indicating any errors or warnings associated with the
// selection of option values.
func (c *Client) ValidateConfigurationSettings(ctx context.Context, params *ValidateConfigurationSettingsInput, optFns ...func(*Options)) (*ValidateConfigurationSettingsOutput, error) {
	stack := middleware.NewStack("ValidateConfigurationSettings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpValidateConfigurationSettingsMiddlewares(stack)
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
	addOpValidateConfigurationSettingsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opValidateConfigurationSettings(options.Region), middleware.Before)

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
			OperationName: "ValidateConfigurationSettings",
			Err:           err,
		}
	}
	out := result.(*ValidateConfigurationSettingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A list of validation messages for a specified configuration template.
type ValidateConfigurationSettingsInput struct {
	// The name of the environment to validate the settings against. Condition: You
	// cannot specify both this and a configuration template name.
	EnvironmentName *string
	// The name of the configuration template to validate the settings against.
	// Condition: You cannot specify both this and an environment name.
	TemplateName *string
	// The name of the application that the configuration template or environment
	// belongs to.
	ApplicationName *string
	// A list of the options and desired values to evaluate.
	OptionSettings []*types.ConfigurationOptionSetting
}

// Provides a list of validation messages.
type ValidateConfigurationSettingsOutput struct {
	// A list of ValidationMessage ().
	Messages []*types.ValidationMessage

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpValidateConfigurationSettingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpValidateConfigurationSettings{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpValidateConfigurationSettings{}, middleware.After)
}

func newServiceMetadataMiddleware_opValidateConfigurationSettings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "ValidateConfigurationSettings",
	}
}