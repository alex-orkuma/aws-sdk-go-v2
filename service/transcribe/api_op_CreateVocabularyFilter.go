// Code generated by smithy-go-codegen DO NOT EDIT.

package transcribe

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transcribe/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a new vocabulary filter that you can use to filter words, such as
// profane words, from the output of a transcription job.
func (c *Client) CreateVocabularyFilter(ctx context.Context, params *CreateVocabularyFilterInput, optFns ...func(*Options)) (*CreateVocabularyFilterOutput, error) {
	stack := middleware.NewStack("CreateVocabularyFilter", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateVocabularyFilterMiddlewares(stack)
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
	addOpCreateVocabularyFilterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVocabularyFilter(options.Region), middleware.Before)

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
			OperationName: "CreateVocabularyFilter",
			Err:           err,
		}
	}
	out := result.(*CreateVocabularyFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVocabularyFilterInput struct {
	// The Amazon S3 location of a text file used as input to create the vocabulary
	// filter. Only use characters from the character set defined for custom
	// vocabularies. For a list of character sets, see Character Sets for Custom
	// Vocabularies
	// (https://docs.aws.amazon.com/transcribe/latest/dg/how-vocabulary.html#charsets).
	// The specified file must be less than 50 KB of UTF-8 characters. If you provide
	// the location of a list of words in the VocabularyFilterFileUri parameter, you
	// can't use the Words parameter.
	VocabularyFilterFileUri *string
	// The language code of the words in the vocabulary filter. All words in the filter
	// must be in the same language. The vocabulary filter can only be used with
	// transcription jobs in the specified language.
	LanguageCode types.LanguageCode
	// The words to use in the vocabulary filter. Only use characters from the
	// character set defined for custom vocabularies. For a list of character sets, see
	// Character Sets for Custom Vocabularies
	// (https://docs.aws.amazon.com/transcribe/latest/dg/how-vocabulary.html#charsets).
	// If you provide a list of words in the Words parameter, you can't use the
	// VocabularyFilterFileUri parameter.
	Words []*string
	// The vocabulary filter name. The name must be unique within the account that
	// contains it.If you try to create a vocabulary filter with the same name as a
	// previous vocabulary filter you will receive a ConflictException error.
	VocabularyFilterName *string
}

type CreateVocabularyFilterOutput struct {
	// The name of the vocabulary filter.
	VocabularyFilterName *string
	// The date and time that the vocabulary filter was modified.
	LastModifiedTime *time.Time
	// The language code of the words in the collection.
	LanguageCode types.LanguageCode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateVocabularyFilterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateVocabularyFilter{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateVocabularyFilter{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateVocabularyFilter(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transcribe",
		OperationName: "CreateVocabularyFilter",
	}
}
