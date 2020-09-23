// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package marketplacemetering provides the client and types for making API
// requests to AWSMarketplace Metering.
//
// This reference provides descriptions of the low-level AWS Marketplace Metering
// Service API.
//
// AWS Marketplace sellers can use this API to submit usage data for custom
// usage dimensions.
//
// For information on the permissions you need to use this API, see AWS Marketing
// metering and entitlement API permissions (https://docs.aws.amazon.com/marketplace/latest/userguide/iam-user-policy-for-aws-marketplace-actions.html)
// in the AWS Marketplace Seller Guide.
//
// Submitting Metering Records
//
//    * MeterUsage- Submits the metering record for a Marketplace product. MeterUsage
//    is called from an EC2 instance or a container running on EKS or ECS.
//
//    * BatchMeterUsage- Submits the metering record for a set of customers.
//    BatchMeterUsage is called from a software-as-a-service (SaaS) application.
//
// Accepting New Customers
//
//    * ResolveCustomer- Called by a SaaS application during the registration
//    process. When a buyer visits your website during the registration process,
//    the buyer submits a Registration Token through the browser. The Registration
//    Token is resolved through this API to obtain a CustomerIdentifier and
//    Product Code.
//
// Entitlement and Metering for Paid Container Products
//
//    * Paid container software products sold through AWS Marketplace must integrate
//    with the AWS Marketplace Metering Service and call the RegisterUsage operation
//    for software entitlement and metering. Free and BYOL products for Amazon
//    ECS or Amazon EKS aren't required to call RegisterUsage, but you can do
//    so if you want to receive usage data in your seller reports. For more
//    information on using the RegisterUsage operation, see Container-Based
//    Products (https://docs.aws.amazon.com/marketplace/latest/userguide/container-based-products.html).
//
// BatchMeterUsage API calls are captured by AWS CloudTrail. You can use Cloudtrail
// to verify that the SaaS metering records that you sent are accurate by searching
// for records with the eventName of BatchMeterUsage. You can also use CloudTrail
// to audit records over time. For more information, see the AWS CloudTrail
// User Guide (http://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html) .
//
// See https://docs.aws.amazon.com/goto/WebAPI/meteringmarketplace-2016-01-14 for more information on this service.
//
// See marketplacemetering package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/marketplacemetering/
//
// Using the Client
//
// To use AWSMarketplace Metering with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWSMarketplace Metering client for more information on
// creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/marketplacemetering/#New
package marketplacemetering