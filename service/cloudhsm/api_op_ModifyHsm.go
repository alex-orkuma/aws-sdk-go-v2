// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudhsm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is documentation for AWS CloudHSM Classic. For more information, see AWS
// CloudHSM Classic FAQs (http://aws.amazon.com/cloudhsm/faqs-classic/), the AWS
// CloudHSM Classic User Guide
// (https://docs.aws.amazon.com/cloudhsm/classic/userguide/), and the AWS CloudHSM
// Classic API Reference
// (https://docs.aws.amazon.com/cloudhsm/classic/APIReference/). For information
// about the current version of AWS CloudHSM, see AWS CloudHSM
// (http://aws.amazon.com/cloudhsm/), the AWS CloudHSM User Guide
// (https://docs.aws.amazon.com/cloudhsm/latest/userguide/), and the AWS CloudHSM
// API Reference (https://docs.aws.amazon.com/cloudhsm/latest/APIReference/).
// Modifies an HSM. This operation can result in the HSM being offline for up to 15
// minutes while the AWS CloudHSM service is reconfigured. If you are modifying a
// production HSM, you should ensure that your AWS CloudHSM service is configured
// for high availability, and consider executing this operation during a
// maintenance window.
func (c *Client) ModifyHsm(ctx context.Context, params *ModifyHsmInput, optFns ...func(*Options)) (*ModifyHsmOutput, error) {
	stack := middleware.NewStack("ModifyHsm", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpModifyHsmMiddlewares(stack)
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
	addOpModifyHsmValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyHsm(options.Region), middleware.Before)

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
			OperationName: "ModifyHsm",
			Err:           err,
		}
	}
	out := result.(*ModifyHsmOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the inputs for the ModifyHsm () operation.
type ModifyHsmInput struct {
	// The new identifier of the subnet that the HSM is in. The new subnet must be in
	// the same Availability Zone as the current subnet.
	SubnetId *string
	// The new external ID.
	ExternalId *string
	// The new IP address for the elastic network interface (ENI) attached to the HSM.
	// If the HSM is moved to a different subnet, and an IP address is not specified,
	// an IP address will be randomly chosen from the CIDR range of the new subnet.
	EniIp *string
	// The ARN of the HSM to modify.
	HsmArn *string
	// The new IAM role ARN.
	IamRoleArn *string
	// The new IP address for the syslog monitoring server. The AWS CloudHSM service
	// only supports one syslog monitoring server.
	SyslogIp *string
}

// Contains the output of the ModifyHsm () operation.
type ModifyHsmOutput struct {
	// The ARN of the HSM.
	HsmArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpModifyHsmMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpModifyHsm{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyHsm{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyHsm(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudhsm",
		OperationName: "ModifyHsm",
	}
}
