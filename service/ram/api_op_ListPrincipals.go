// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the principals that you have shared resources with or that have shared
// resources with you.
func (c *Client) ListPrincipals(ctx context.Context, params *ListPrincipalsInput, optFns ...func(*Options)) (*ListPrincipalsOutput, error) {
	stack := middleware.NewStack("ListPrincipals", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPrincipalsMiddlewares(stack)
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
	addOpListPrincipalsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPrincipals(options.Region), middleware.Before)

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
			OperationName: "ListPrincipals",
			Err:           err,
		}
	}
	out := result.(*ListPrincipalsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPrincipalsInput struct {
	// The principals.
	Principals []*string
	// The token for the next page of results.
	NextToken *string
	// The Amazon Resource Names (ARN) of the resource shares.
	ResourceShareArns []*string
	// The type of owner.
	ResourceOwner types.ResourceOwner
	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32
	// The Amazon Resource Name (ARN) of the resource.
	ResourceArn *string
	// The resource type. Valid values: codebuild:Project | codebuild:ReportGroup |
	// ec2:CapacityReservation | ec2:DedicatedHost | ec2:Subnet |
	// ec2:TrafficMirrorTarget | ec2:TransitGateway | imagebuilder:Component |
	// imagebuilder:Image | imagebuilder:ImageRecipe |
	// license-manager:LicenseConfiguration I resource-groups:Group | rds:Cluster |
	// route53resolver:ResolverRule
	ResourceType *string
}

type ListPrincipalsOutput struct {
	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string
	// The principals.
	Principals []*types.Principal

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPrincipalsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPrincipals{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPrincipals{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPrincipals(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "ListPrincipals",
	}
}
