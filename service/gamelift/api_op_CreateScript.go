// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new script record for your Realtime Servers script. Realtime scripts
// are JavaScript that provide configuration settings and optional custom game
// logic for your game. The script is deployed when you create a Realtime Servers
// fleet to host your game sessions. Script logic is executed during an active game
// session. To create a new script record, specify a script name and provide the
// script file(s). The script files and all dependencies must be zipped into a
// single file. You can pull the zip file from either of these locations:
//
//     * A
// locally available directory. Use the ZipFile parameter for this option.
//
//     *
// An Amazon Simple Storage Service (Amazon S3) bucket under your AWS account. Use
// the StorageLocation parameter for this option. You'll need to have an Identity
// Access Management (IAM) role that allows the Amazon GameLift service to access
// your S3 bucket.
//
// If the call is successful, a new script record is created with
// a unique script ID. If the script file is provided as a local file, the file is
// uploaded to an Amazon GameLift-owned S3 bucket and the script record's storage
// location reflects this location. If the script file is provided as an S3 bucket,
// Amazon GameLift accesses the file at this storage location as needed for
// deployment. Learn more  <p> <a
// href="https://docs.aws.amazon.com/gamelift/latest/developerguide/realtime-intro.html">Amazon
// GameLift Realtime Servers</a> </p> <p> <a
// href="https://docs.aws.amazon.com/gamelift/latest/developerguide/setting-up-role.html">Set
// Up a Role for Amazon GameLift Access</a> </p> <p> <b>Related operations</b> </p>
// <ul> <li> <p> <a>CreateScript</a> </p> </li> <li> <p> <a>ListScripts</a> </p>
// </li> <li> <p> <a>DescribeScript</a> </p> </li> <li> <p> <a>UpdateScript</a>
// </p> </li> <li> <p> <a>DeleteScript</a> </p> </li> </ul>
func (c *Client) CreateScript(ctx context.Context, params *CreateScriptInput, optFns ...func(*Options)) (*CreateScriptOutput, error) {
	stack := middleware.NewStack("CreateScript", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateScriptMiddlewares(stack)
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
	addOpCreateScriptValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateScript(options.Region), middleware.Before)

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
			OperationName: "CreateScript",
			Err:           err,
		}
	}
	out := result.(*CreateScriptOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateScriptInput struct {
	// The version that is associated with a build or script. Version strings do not
	// need to be unique. You can use UpdateScript () to change this value later.
	Version *string
	// A descriptive label that is associated with a script. Script names do not need
	// to be unique. You can use UpdateScript () to change this value later.
	Name *string
	// A list of labels to assign to the new script resource. Tags are
	// developer-defined key-value pairs. Tagging AWS resources are useful for resource
	// management, access management and cost allocation. For more information, see
	// Tagging AWS Resources
	// (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in the AWS
	// General Reference. Once the resource is created, you can use TagResource (),
	// UntagResource (), and ListTagsForResource () to add, remove, and view tags. The
	// maximum tag limit may be lower than stated. See the AWS General Reference for
	// actual tagging limits.
	Tags []*types.Tag
	// A data object containing your Realtime scripts and dependencies as a zip file.
	// The zip file can have one or multiple files. Maximum size of a zip file is 5 MB.
	// When using the AWS CLI tool to create a script, this parameter is set to the zip
	// file name. It must be prepended with the string "fileb://" to indicate that the
	// file data is a binary object. For example: --zip-file
	// fileb://myRealtimeScript.zip.
	ZipFile []byte
	// The location of the Amazon S3 bucket where a zipped file containing your
	// Realtime scripts is stored. The storage location must specify the Amazon S3
	// bucket name, the zip file name (the "key"), and a role ARN that allows Amazon
	// GameLift to access the Amazon S3 storage location. The S3 bucket must be in the
	// same Region where you want to create a new script. By default, Amazon GameLift
	// uploads the latest version of the zip file; if you have S3 object versioning
	// turned on, you can use the ObjectVersion parameter to specify an earlier
	// version.
	StorageLocation *types.S3Location
}

type CreateScriptOutput struct {
	// The newly created script record with a unique script ID and ARN. The new
	// script's storage location reflects an Amazon S3 location: (1) If the script was
	// uploaded from an S3 bucket under your account, the storage location reflects the
	// information that was provided in the CreateScript request; (2) If the script
	// file was uploaded from a local zip file, the storage location reflects an S3
	// location controls by the Amazon GameLift service.
	Script *types.Script

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateScriptMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateScript{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateScript{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateScript(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "CreateScript",
	}
}
