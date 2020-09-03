// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns all of the runtime parameters offered by the underlying database
// software, or engine, for a specific database in Amazon Lightsail. In addition to
// the parameter names and values, this operation returns other information about
// each parameter. This information includes whether changes require a reboot,
// whether the parameter is modifiable, the allowed values, and the data types.
func (c *Client) GetRelationalDatabaseParameters(ctx context.Context, params *GetRelationalDatabaseParametersInput, optFns ...func(*Options)) (*GetRelationalDatabaseParametersOutput, error) {
	stack := middleware.NewStack("GetRelationalDatabaseParameters", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetRelationalDatabaseParametersMiddlewares(stack)
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
	addOpGetRelationalDatabaseParametersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRelationalDatabaseParameters(options.Region), middleware.Before)

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
			OperationName: "GetRelationalDatabaseParameters",
			Err:           err,
		}
	}
	out := result.(*GetRelationalDatabaseParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabaseParametersInput struct {
	// The token to advance to the next page of results from your request. To get a
	// page token, perform an initial GetRelationalDatabaseParameters request. If your
	// results are paginated, the response will return a next page token that you can
	// specify as the page token in a subsequent request.
	PageToken *string
	// The name of your database for which to get parameters.
	RelationalDatabaseName *string
}

type GetRelationalDatabaseParametersOutput struct {
	// The token to advance to the next page of resutls from your request. A next page
	// token is not returned if there are no more results to display. To get the next
	// page of results, perform another GetRelationalDatabaseParameters request and
	// specify the next page token using the pageToken parameter.
	NextPageToken *string
	// An object describing the result of your get relational database parameters
	// request.
	Parameters []*types.RelationalDatabaseParameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetRelationalDatabaseParametersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetRelationalDatabaseParameters{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRelationalDatabaseParameters{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRelationalDatabaseParameters(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRelationalDatabaseParameters",
	}
}
