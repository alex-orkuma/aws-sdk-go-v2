// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of existing service meshes.
func (c *Client) ListMeshes(ctx context.Context, params *ListMeshesInput, optFns ...func(*Options)) (*ListMeshesOutput, error) {
	stack := middleware.NewStack("ListMeshes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListMeshesMiddlewares(stack)
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
			OperationName: "ListMeshes",
			Err:           err,
		}
	}
	out := result.(*ListMeshesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ListMeshesInput struct {
	// The nextToken value returned from a previous paginated ListMeshes request where
	// limit was used and the results exceeded the value of that parameter. Pagination
	// continues from the end of the previous results that returned the nextToken
	// value. This token should be treated as an opaque identifier that is used only to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string
	// The maximum number of results returned by ListMeshes in paginated output. When
	// you use this parameter, ListMeshes returns only limit results in a single page
	// along with a nextToken response element. You can see the remaining results of
	// the initial request by sending another ListMeshes request with the returned
	// nextToken value. This value can be between 1 and 100. If you don't use this
	// parameter, ListMeshes returns up to 100 results and a nextToken value if
	// applicable.
	Limit *int32
}

//
type ListMeshesOutput struct {
	// The list of existing service meshes.
	Meshes []*types.MeshRef
	// The nextToken value to include in a future ListMeshes request. When the results
	// of a ListMeshes request exceed limit, you can use this value to retrieve the
	// next page of results. This value is null when there are no more results to
	// return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListMeshesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListMeshes{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListMeshes{}, middleware.After)
}
