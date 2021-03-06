// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists available reserved Elasticsearch instance offerings.
func (c *Client) DescribeReservedElasticsearchInstanceOfferings(ctx context.Context, params *DescribeReservedElasticsearchInstanceOfferingsInput, optFns ...func(*Options)) (*DescribeReservedElasticsearchInstanceOfferingsOutput, error) {
	if params == nil {
		params = &DescribeReservedElasticsearchInstanceOfferingsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReservedElasticsearchInstanceOfferings", params, optFns, addOperationDescribeReservedElasticsearchInstanceOfferingsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReservedElasticsearchInstanceOfferingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for parameters to DescribeReservedElasticsearchInstanceOfferings
type DescribeReservedElasticsearchInstanceOfferingsInput struct {

	// Set this value to limit the number of results returned. If not specified,
	// defaults to 100.
	MaxResults *int32

	// NextToken should be sent in case if earlier API call produced result containing
	// NextToken. It is used for pagination.
	NextToken *string

	// The offering identifier filter value. Use this parameter to show only the
	// available offering that matches the specified reservation identifier.
	ReservedElasticsearchInstanceOfferingId *string
}

// Container for results from DescribeReservedElasticsearchInstanceOfferings
type DescribeReservedElasticsearchInstanceOfferingsOutput struct {

	// Provides an identifier to allow retrieval of paginated results.
	NextToken *string

	// List of reserved Elasticsearch instance offerings
	ReservedElasticsearchInstanceOfferings []*types.ReservedElasticsearchInstanceOffering

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeReservedElasticsearchInstanceOfferingsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeReservedElasticsearchInstanceOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeReservedElasticsearchInstanceOfferings{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReservedElasticsearchInstanceOfferings(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribeReservedElasticsearchInstanceOfferings(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "DescribeReservedElasticsearchInstanceOfferings",
	}
}
