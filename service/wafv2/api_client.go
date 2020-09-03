// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	"net/http"
)

const ServiceID = "WAFV2"

// This is the latest version of the AWS WAF API, released in November, 2019. The
// names of the entities that you use to access this API, like endpoints and
// namespaces, all have the versioning information added, like "V2" or "v2", to
// distinguish from the prior version. We recommend migrating your resources to
// this version, because it has a number of significant improvements. If you used
// AWS WAF prior to this release, you can't use this AWS WAFV2 API to access any
// AWS WAF resources that you created before. You can access your old rules, web
// ACLs, and other AWS WAF resources only through the AWS WAF Classic APIs. The AWS
// WAF Classic APIs have retained the prior names, endpoints, and namespaces. For
// information, including how to migrate your AWS WAF resources to this version,
// see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). AWS
// WAF is a web application firewall that lets you monitor the HTTP and HTTPS
// requests that are forwarded to Amazon CloudFront, an Amazon API Gateway API, or
// an Application Load Balancer. AWS WAF also lets you control access to your
// content. Based on conditions that you specify, such as the IP addresses that
// requests originate from or the values of query strings, API Gateway, CloudFront,
// or the Application Load Balancer responds to requests either with the requested
// content or with an HTTP 403 status code (Forbidden). You also can configure
// CloudFront to return a custom error page when a request is blocked. This API
// guide is for developers who need detailed information about AWS WAF API actions,
// data types, and errors. For detailed information about AWS WAF features and an
// overview of how to use AWS WAF, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/). You can make calls
// using the endpoints listed in AWS Service Endpoints for AWS WAF
// (https://docs.aws.amazon.com/general/latest/gr/rande.html#waf_region).
//
//     *
// For regional applications, you can use any of the endpoints in the list. A
// regional application can be an Application Load Balancer (ALB) or an API Gateway
// stage.
//
//     * For AWS CloudFront applications, you must use the API endpoint
// listed for US East (N. Virginia): us-east-1.
//
// Alternatively, you can use one of
// the AWS SDKs to access an API that's tailored to the programming language or
// platform that you're using. For more information, see AWS SDKs
// (http://aws.amazon.com/tools/#SDKs). We currently provide two versions of the
// AWS WAF API: this API and the prior versions, the classic AWS WAF APIs. This new
// API provides the same functionality as the older versions, with the following
// major improvements:
//
//     * You use one API for both global and regional
// applications. Where you need to distinguish the scope, you specify a Scope
// parameter and set it to CLOUDFRONT or REGIONAL.
//
//     * You can define a Web ACL
// or rule group with a single call, and update it with a single call. You define
// all rule specifications in JSON format, and pass them to your rule group or Web
// ACL calls.
//
//     * The limits AWS WAF places on the use of rules more closely
// reflects the cost of running each type of rule. Rule groups include capacity
// settings, so you know the maximum cost of a rule group when you use it.
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveDefaultEndpointConfiguration(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []APIOptionFunc

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions ResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// An integer value representing the logging level.
	LogLevel aws.LogLevel

	// The logger writer interface to write logging messages to.
	Logger aws.Logger

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

func (o Options) GetCredentials() aws.CredentialsProvider {
	return o.Credentials
}

func (o Options) GetEndpointOptions() ResolverOptions {
	return o.EndpointOptions
}

func (o Options) GetEndpointResolver() EndpointResolver {
	return o.EndpointResolver
}

func (o Options) GetLogLevel() aws.LogLevel {
	return o.LogLevel
}

func (o Options) GetLogger() aws.Logger {
	return o.Logger
}

func (o Options) GetRegion() string {
	return o.Region
}

func (o Options) GetRetryer() retry.Retryer {
	return o.Retryer
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]APIOptionFunc, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}

type APIOptionFunc func(*middleware.Stack) error

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:      cfg.Region,
		Retryer:     cfg.Retryer,
		LogLevel:    cfg.LogLevel,
		Logger:      cfg.Logger,
		HTTPClient:  cfg.HTTPClient,
		Credentials: cfg.Credentials,
	}
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = aws.NewBuildableHTTPClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func addClientUserAgent(stack *middleware.Stack) {
	awsmiddleware.AddUserAgentKey("wafv2")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) {
	signer := v4.Signer{}
	stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, signer), middleware.After)
}
