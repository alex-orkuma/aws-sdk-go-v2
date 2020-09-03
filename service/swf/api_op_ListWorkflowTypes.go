// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about workflow types in the specified domain. The results
// may be split into multiple pages that can be retrieved by making the call
// repeatedly. Access Control You can use IAM policies to control this action's
// access to Amazon SWF resources as follows:
//
//     * Use a Resource element with
// the domain name to limit the action to only specified domains.
//
//     * Use an
// Action element to allow or deny permission to call this action.
//
//     * You
// cannot use an IAM policy to constrain this action's parameters.
//
// If the caller
// doesn't have sufficient permissions to invoke the action, or the parameter
// values fall outside the specified constraints, the action fails. The associated
// event attribute's cause parameter is set to OPERATION_NOT_PERMITTED. For details
// and example IAM policies, see Using IAM to Manage Access to Amazon SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) ListWorkflowTypes(ctx context.Context, params *ListWorkflowTypesInput, optFns ...func(*Options)) (*ListWorkflowTypesOutput, error) {
	stack := middleware.NewStack("ListWorkflowTypes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpListWorkflowTypesMiddlewares(stack)
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
	addOpListWorkflowTypesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkflowTypes(options.Region), middleware.Before)

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
			OperationName: "ListWorkflowTypes",
			Err:           err,
		}
	}
	out := result.(*ListWorkflowTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkflowTypesInput struct {
	// Specifies the registration status of the workflow types to list.
	RegistrationStatus types.RegistrationStatus
	// The name of the domain in which the workflow types have been registered.
	Domain *string
	// When set to true, returns the results in reverse order. By default the results
	// are returned in ascending alphabetical order of the name of the workflow types.
	ReverseOrder *bool
	// If NextPageToken is returned there are more results available. The value of
	// NextPageToken is a unique pagination token for each page. Make the call again
	// using the returned token to retrieve the next page. Keep all other arguments
	// unchanged. Each pagination token expires after 60 seconds. Using an expired
	// pagination token will return a 400 error: "Specified token has exceeded its
	// maximum lifetime".  <p>The configured <code>maximumPageSize</code> determines
	// how many results can be returned in a single call. </p>
	NextPageToken *string
	// The maximum number of results that are returned per call. Use nextPageToken to
	// obtain further pages of results.
	MaximumPageSize *int32
	// If specified, lists the workflow type with this name.
	Name *string
}

// Contains a paginated list of information structures about workflow types.
type ListWorkflowTypesOutput struct {
	// The list of workflow type information.
	TypeInfos []*types.WorkflowTypeInfo
	// If a NextPageToken was returned by a previous call, there are more results
	// available. To retrieve the next page of results, make the call again using the
	// returned token in nextPageToken. Keep all other arguments unchanged. The
	// configured maximumPageSize determines how many results can be returned in a
	// single call.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpListWorkflowTypesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpListWorkflowTypes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpListWorkflowTypes{}, middleware.After)
}

func newServiceMetadataMiddleware_opListWorkflowTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "ListWorkflowTypes",
	}
}
