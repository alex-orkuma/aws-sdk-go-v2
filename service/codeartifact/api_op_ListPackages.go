// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of PackageSummary
// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageSummary.html)
// objects for packages in a repository that match the request parameters.
func (c *Client) ListPackages(ctx context.Context, params *ListPackagesInput, optFns ...func(*Options)) (*ListPackagesOutput, error) {
	stack := middleware.NewStack("ListPackages", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListPackagesMiddlewares(stack)
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
	addOpListPackagesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPackages(options.Region), middleware.Before)

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
			OperationName: "ListPackages",
			Err:           err,
		}
	}
	out := result.(*ListPackagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPackagesInput struct {
	// The namespace of the package. The package component that specifies its namespace
	// depends on its type. For example:
	//
	//     * The namespace of a Maven package is its
	// groupId.
	//
	//     * The namespace of an npm package is its scope.
	//
	//     * A Python
	// package does not contain a corresponding component, so Python packages do not
	// have a namespace.
	Namespace *string
	// The format of the packages. The valid package types are:
	//
	//     * npm: A Node
	// Package Manager (npm) package.
	//
	//     * pypi: A Python Package Index (PyPI)
	// package.
	//
	//     * maven: A Maven package that contains compiled code in a
	// distributable format, such as a JAR file.
	Format types.PackageFormat
	// The 12-digit account number of the AWS account that owns the domain. It does not
	// include dashes or spaces.
	DomainOwner *string
	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string
	// The maximum number of results to return per page.
	MaxResults *int32
	// A prefix used to filter returned repositories. Only repositories with names that
	// start with repositoryPrefix are returned.
	PackagePrefix *string
	// The name of the repository from which packages are to be listed.
	Repository *string
	// The domain that contains the repository that contains the requested list of
	// packages.
	Domain *string
}

type ListPackagesOutput struct {
	// If there are additional results, this is the token for the next set of results.
	NextToken *string
	// The list of returned PackageSummary
	// (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_PackageSummary.html)
	// objects.
	Packages []*types.PackageSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListPackagesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListPackages{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListPackages{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPackages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "ListPackages",
	}
}
