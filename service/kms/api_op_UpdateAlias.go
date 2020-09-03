// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates an existing AWS KMS alias with a different customer master key (CMK).
// Each alias is associated with only one CMK at a time, although a CMK can have
// multiple aliases. The alias and the CMK must be in the same AWS account and
// region. You cannot perform this operation on an alias in a different AWS
// account. The current and new CMK must be the same type (both symmetric or both
// asymmetric), and they must have the same key usage (ENCRYPT_DECRYPT or
// SIGN_VERIFY). This restriction prevents errors in code that uses aliases. If you
// must assign an alias to a different type of CMK, use DeleteAlias () to delete
// the old alias and CreateAlias () to create a new alias. You cannot use
// UpdateAlias to change an alias name. To change an alias name, use DeleteAlias ()
// to delete the old alias and CreateAlias () to create a new alias. Because an
// alias is not a property of a CMK, you can create, update, and delete the aliases
// of a CMK without affecting the CMK. Also, aliases do not appear in the response
// from the DescribeKey () operation. To get the aliases of all CMKs in the
// account, use the ListAliases () operation. The CMK that you use for this
// operation must be in a compatible key state. For details, see How Key State
// Affects Use of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) UpdateAlias(ctx context.Context, params *UpdateAliasInput, optFns ...func(*Options)) (*UpdateAliasOutput, error) {
	stack := middleware.NewStack("UpdateAlias", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateAliasMiddlewares(stack)
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
	addOpUpdateAliasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAlias(options.Region), middleware.Before)

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
			OperationName: "UpdateAlias",
			Err:           err,
		}
	}
	out := result.(*UpdateAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAliasInput struct {
	// Identifies the CMK to associate with the alias. When the update operation
	// completes, the alias will point to this CMK. The CMK must be in the same AWS
	// account and Region as the alias. Also, the new target CMK must be the same type
	// as the current target CMK (both symmetric or both asymmetric) and they must have
	// the same key usage. Specify the key ID or the Amazon Resource Name (ARN) of the
	// CMK. For example:
	//
	//     * Key ID: 1234abcd-12ab-34cd-56ef-1234567890ab
	//
	//     * Key
	// ARN:
	// arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab
	//
	// To
	// get the key ID and key ARN for a CMK, use ListKeys () or DescribeKey (). To
	// verify that the alias is mapped to the correct CMK, use ListAliases ().
	TargetKeyId *string
	// Identifies the alias that is changing its CMK. This value must begin with alias/
	// followed by the alias name, such as alias/ExampleAlias. You cannot use
	// UpdateAlias to change the alias name.
	AliasName *string
}

type UpdateAliasOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateAliasMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateAlias{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateAlias{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateAlias(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "UpdateAlias",
	}
}
