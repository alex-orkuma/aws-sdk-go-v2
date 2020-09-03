// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Inspects text for syntax and the part of speech of words in the document. For
// more information, how-syntax ().
func (c *Client) DetectSyntax(ctx context.Context, params *DetectSyntaxInput, optFns ...func(*Options)) (*DetectSyntaxOutput, error) {
	stack := middleware.NewStack("DetectSyntax", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDetectSyntaxMiddlewares(stack)
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
	addOpDetectSyntaxValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetectSyntax(options.Region), middleware.Before)

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
			OperationName: "DetectSyntax",
			Err:           err,
		}
	}
	out := result.(*DetectSyntaxOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectSyntaxInput struct {
	// A UTF-8 string. Each string must contain fewer that 5,000 bytes of UTF encoded
	// characters.
	Text *string
	// The language code of the input documents. You can specify any of the following
	// languages supported by Amazon Comprehend: German ("de"), English ("en"), Spanish
	// ("es"), French ("fr"), Italian ("it"), or Portuguese ("pt").
	LanguageCode types.SyntaxLanguageCode
}

type DetectSyntaxOutput struct {
	// A collection of syntax tokens describing the text. For each token, the response
	// provides the text, the token type, where the text begins and ends, and the level
	// of confidence that Amazon Comprehend has that the token is correct. For a list
	// of token types, see how-syntax ().
	SyntaxTokens []*types.SyntaxToken

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDetectSyntaxMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDetectSyntax{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectSyntax{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetectSyntax(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DetectSyntax",
	}
}
