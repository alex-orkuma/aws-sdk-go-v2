// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// This structure contains all information about one canary in your account.
type Canary struct {
	// The unique ID of this canary.
	Id *string
	// The ARN of the Lambda function that is used as your canary's engine. For more
	// information about Lambda ARN format, see Resources and Conditions for Lambda
	// Actions
	// (https://docs.aws.amazon.com/lambda/latest/dg/lambda-api-permissions-ref.html).
	EngineArn *string
	// If this canary is to test an endpoint in a VPC, this structure contains
	// information about the subnets and security groups of the VPC endpoint. For more
	// information, see  Running a Canary in a VPC
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html).
	VpcConfig *VpcConfigOutput
	// The number of days to retain data about successful runs of this canary.
	SuccessRetentionPeriodInDays *int32
	// Specifies the runtime version to use for the canary. Currently, the only valid
	// value is syn-1.0. For more information about runtime versions, see  Canary
	// Runtime Versions
	// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html).
	RuntimeVersion *string
	// The list of key-value pairs that are associated with the canary.
	Tags map[string]*string
	// The number of days to retain data about failed runs of this canary.
	FailureRetentionPeriodInDays *int32
	// A structure that contains information about the canary's status.
	Status *CanaryStatus
	// The location in Amazon S3 where Synthetics stores artifacts from the runs of
	// this canary. Artifacts include the log file, screenshots, and HAR files.
	ArtifactS3Location *string
	// The ARN of the IAM role used to run the canary. This role must include
	// lambda.amazonaws.com as a principal in the trust policy.
	ExecutionRoleArn *string
	// This structure contains information about the canary's Lambda handler and where
	// its code is stored by CloudWatch Synthetics.
	Code *CanaryCodeOutput
	// The name of the canary.
	Name *string
	// A structure that contains information about how often the canary is to run, and
	// when these runs are to stop.
	Schedule *CanaryScheduleOutput
	// A structure that contains information about when the canary was created,
	// modified, and most recently run.
	Timeline *CanaryTimeline
	// A structure that contains information for a canary run.
	RunConfig *CanaryRunConfigOutput
}

// Use this structure to input your script code for the canary. This structure
// contains the Lambda handler with the location where the canary should start
// running the script. If the script is stored in an S3 bucket, the bucket name,
// key, and version are also included. If the script was passed into the canary
// directly, the script code is contained in the value of Zipfile.
type CanaryCodeInput struct {
	// The S3 version ID of your script.
	S3Version *string
	// The entry point to use for the source code when running the canary. This value
	// must end with the string .handler.
	Handler *string
	// The S3 key of your script. For more information, see Working with Amazon S3
	// Objects (https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingObjects.html).
	S3Key *string
	// If your canary script is located in S3, specify the full bucket name here. The
	// bucket must already exist. Specify the full bucket name, including s3:// as the
	// start of the bucket name.
	S3Bucket *string
	// If you input your canary script directly into the canary instead of referring to
	// an S3 location, the value of this parameter is the .zip file that contains the
	// script. It can be up to 5 MB.
	ZipFile []byte
}

// This structure contains information about the canary's Lambda handler and where
// its code is stored by CloudWatch Synthetics.
type CanaryCodeOutput struct {
	// The ARN of the Lambda layer where Synthetics stores the canary script code.
	SourceLocationArn *string
	// The entry point to use for the source code when running the canary.
	Handler *string
}

// This structure contains information about the most recent run of a single
// canary.
type CanaryLastRun struct {
	// The name of the canary.
	CanaryName *string
	// The results from this canary's most recent run.
	LastRun *CanaryRun
}

// This structure contains the details about one run of one canary.
type CanaryRun struct {
	// The name of the canary.
	Name *string
	// The location where the canary stored artifacts from the run. Artifacts include
	// the log file, screenshots, and HAR files.
	ArtifactS3Location *string
	// A structure that contains the start and end times of this run.
	Timeline *CanaryRunTimeline
	// The status of this run.
	Status *CanaryRunStatus
}

// A structure that contains input information for a canary run.
type CanaryRunConfigInput struct {
	// The maximum amount of memory available to the canary while it is running, in MB.
	// The value you specify must be a multiple of 64.
	MemoryInMB *int32
	// How long the canary is allowed to run before it must stop. If you omit this
	// field, the frequency of the canary is used as this value, up to a maximum of 14
	// minutes.
	TimeoutInSeconds *int32
}

// A structure that contains information for a canary run.
type CanaryRunConfigOutput struct {
	// How long the canary is allowed to run before it must stop.
	TimeoutInSeconds *int32
	// The maximum amount of memory available to the canary while it is running, in MB.
	// The value you must be a multiple of 64.
	MemoryInMB *int32
}

// This structure contains the status information about a canary run.
type CanaryRunStatus struct {
	// If this value is CANARY_FAILURE, an exception occurred in the canary code. If
	// this value is EXECUTION_FAILURE, an exception occurred in CloudWatch Synthetics.
	StateReasonCode CanaryRunStateReasonCode
	// If run of the canary failed, this field contains the reason for the error.
	StateReason *string
	// The current state of the run.
	State CanaryRunState
}

// This structure contains the start and end times of a single canary run.
type CanaryRunTimeline struct {
	// The start time of the run.
	Started *time.Time
	// The end time of the run.
	Completed *time.Time
}

// This structure specifies how often a canary is to make runs and the date and
// time when it should stop making runs.
type CanaryScheduleInput struct {
	// How long, in seconds, for the canary to continue making regular runs according
	// to the schedule in the Expression value. If you specify 0, the canary continues
	// making runs until you stop it. If you omit this field, the default of 0 is used.
	DurationInSeconds *int64
	// A rate expression that defines how often the canary is to run. The syntax is
	// rate(number unit). unit can be minute, minutes, or hour. For example, rate(1
	// minute) runs the canary once a minute, rate(10 minutes) runs it once every 10
	// minutes, and rate(1 hour) runs it once every hour. You can specify a frequency
	// between rate(1 minute) and rate(1 hour). Specifying rate(0 minute) or rate(0
	// hour) is a special value that causes the canary to run only once when it is
	// started.
	Expression *string
}

// How long, in seconds, for the canary to continue making regular runs according
// to the schedule in the Expression value.
type CanaryScheduleOutput struct {
	// A rate expression that defines how often the canary is to run. The syntax is
	// rate(number unit). unit can be minute, minutes, or hour. For example, rate(1
	// minute) runs the canary once a minute, rate(10 minutes) runs it once every 10
	// minutes, and rate(1 hour) runs it once every hour. Specifying rate(0 minute) or
	// rate(0 hour) is a special value that causes the canary to run only once when it
	// is started.
	Expression *string
	// How long, in seconds, for the canary to continue making regular runs after it
	// was created. The runs are performed according to the schedule in the Expression
	// value.
	DurationInSeconds *int64
}

// A structure that contains the current state of the canary.
type CanaryStatus struct {
	// The current state of the canary.
	State CanaryState
	// If the canary has insufficient permissions to run, this field provides more
	// details.
	StateReason *string
	// If the canary cannot run or has failed, this field displays the reason.
	StateReasonCode CanaryStateReasonCode
}

// This structure contains information about when the canary was created and
// modified.
type CanaryTimeline struct {
	// The date and time that the canary's most recent run started.
	LastStarted *time.Time
	// The date and time the canary was created.
	Created *time.Time
	// The date and time the canary was most recently modified.
	LastModified *time.Time
	// The date and time that the canary's most recent run ended.
	LastStopped *time.Time
}

// This structure contains information about one canary runtime version. For more
// information about runtime versions, see  Canary Runtime Versions
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_Library.html).
type RuntimeVersion struct {
	// The name of the runtime version. Currently, the only valid value is syn-1.0.
	// Specifies the runtime version to use for the canary. Currently, the only valid
	// value is syn-1.0.
	VersionName *string
	// The date that the runtime version was released.
	ReleaseDate *time.Time
	// A description of the runtime version, created by Amazon.
	Description *string
	// If this runtime version is deprecated, this value is the date of deprecation.
	DeprecationDate *time.Time
}

// If this canary is to test an endpoint in a VPC, this structure contains
// information about the subnets and security groups of the VPC endpoint. For more
// information, see  Running a Canary in a VPC
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html).
type VpcConfigInput struct {
	// The IDs of the subnets where this canary is to run.
	SubnetIds []*string
	// The IDs of the security groups for this canary.
	SecurityGroupIds []*string
}

// If this canary is to test an endpoint in a VPC, this structure contains
// information about the subnets and security groups of the VPC endpoint. For more
// information, see  Running a Canary in a VPC
// (https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/CloudWatch_Synthetics_Canaries_VPC.html).
type VpcConfigOutput struct {
	// The IDs of the subnets where this canary is to run.
	SubnetIds []*string
	// The IDs of the security groups for this canary.
	SecurityGroupIds []*string
	// The IDs of the VPC where this canary is to run.
	VpcId *string
}
