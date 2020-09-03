// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns metadata associated with a recovery point, including ID, status,
// encryption, and lifecycle.
func (c *Client) DescribeRecoveryPoint(ctx context.Context, params *DescribeRecoveryPointInput, optFns ...func(*Options)) (*DescribeRecoveryPointOutput, error) {
	stack := middleware.NewStack("DescribeRecoveryPoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeRecoveryPointMiddlewares(stack)
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
	addOpDescribeRecoveryPointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRecoveryPoint(options.Region), middleware.Before)

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
			OperationName: "DescribeRecoveryPoint",
			Err:           err,
		}
	}
	out := result.(*DescribeRecoveryPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRecoveryPointInput struct {
	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	BackupVaultName *string
	// An Amazon Resource Name (ARN) that uniquely identifies a recovery point; for
	// example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string
}

type DescribeRecoveryPointOutput struct {
	// The lifecycle defines when a protected resource is transitioned to cold storage
	// and when it expires. AWS Backup transitions and expires backups automatically
	// according to the lifecycle that you define. Backups that are transitioned to
	// cold storage must be stored in cold storage for a minimum of 90 days. Therefore,
	// the “expire after days” setting must be 90 days greater than the “transition to
	// cold after days” setting. The “transition to cold after days” setting cannot be
	// changed after a backup has been transitioned to cold.
	Lifecycle *types.Lifecycle
	// The size, in bytes, of a backup.
	BackupSizeInBytes *int64
	// An ARN that uniquely identifies a backup vault; for example,
	// arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string
	// A status code specifying the state of the recovery point. A partial status
	// indicates that the recovery point was not successfully re-created and must be
	// retried.
	Status types.RecoveryPointStatus
	// The date and time that a recovery point is created, in Unix format and
	// Coordinated Universal Time (UTC). The value of CreationDate is accurate to
	// milliseconds. For example, the value 1516925490.087 represents Friday, January
	// 26, 2018 12:11:30.087 AM.
	CreationDate *time.Time
	// The server-side encryption key used to protect your backups; for example,
	// arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	EncryptionKeyArn *string
	// Specifies the storage class of the recovery point. Valid values are WARM or
	// COLD.
	StorageClass types.StorageClass
	// Contains identifying information about the creation of a recovery point,
	// including the BackupPlanArn, BackupPlanId, BackupPlanVersion, and BackupRuleId
	// of the backup plan used to create it.
	CreatedBy *types.RecoveryPointCreator
	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Region where they are created. They consist of lowercase letters, numbers, and
	// hyphens.
	BackupVaultName *string
	// The date and time that a recovery point was last restored, in Unix format and
	// Coordinated Universal Time (UTC). The value of LastRestoreTime is accurate to
	// milliseconds. For example, the value 1516925490.087 represents Friday, January
	// 26, 2018 12:11:30.087 AM.
	LastRestoreTime *time.Time
	// The date and time that a job to create a recovery point is completed, in Unix
	// format and Coordinated Universal Time (UTC). The value of CompletionDate is
	// accurate to milliseconds. For example, the value 1516925490.087 represents
	// Friday, January 26, 2018 12:11:30.087 AM.
	CompletionDate *time.Time
	// Specifies the IAM role ARN used to create the target recovery point; for
	// example, arn:aws:iam::123456789012:role/S3Access.
	IamRoleArn *string
	// The type of AWS resource to save as a recovery point; for example, an Amazon
	// Elastic Block Store (Amazon EBS) volume or an Amazon Relational Database Service
	// (Amazon RDS) database.
	ResourceType *string
	// An ARN that uniquely identifies a recovery point; for example,
	// arn:aws:backup:us-east-1:123456789012:recovery-point:1EB3B5E7-9EB0-435A-A80B-108B488B0D45.
	RecoveryPointArn *string
	// An ARN that uniquely identifies a saved resource. The format of the ARN depends
	// on the resource type.
	ResourceArn *string
	// A CalculatedLifecycle object containing DeleteAt and MoveToColdStorageAt
	// timestamps.
	CalculatedLifecycle *types.CalculatedLifecycle
	// A Boolean value that is returned as TRUE if the specified recovery point is
	// encrypted, or FALSE if the recovery point is not encrypted.
	IsEncrypted *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeRecoveryPointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeRecoveryPoint{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeRecoveryPoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeRecoveryPoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "DescribeRecoveryPoint",
	}
}
