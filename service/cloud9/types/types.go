// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// Information about an AWS Cloud9 development environment.
type Environment struct {
	// The Amazon Resource Name (ARN) of the environment.
	Arn *string
	// The state of the environment in its creation or deletion lifecycle.
	Lifecycle *EnvironmentLifecycle
	// The type of environment. Valid values include the following:
	//
	//     * ec2: An
	// Amazon Elastic Compute Cloud (Amazon EC2) instance connects to the
	// environment.
	//
	//     * ssh: Your own server connects to the environment.
	Type EnvironmentType
	// The description for the environment.
	Description *string
	// The name of the environment.
	Name *string
	// The Amazon Resource Name (ARN) of the environment owner.
	OwnerArn *string
	// The ID of the environment.
	Id *string
}

// Information about the current creation or deletion lifecycle state of an AWS
// Cloud9 development environment.
type EnvironmentLifecycle struct {
	// The current creation or deletion lifecycle state of the environment.
	//
	//     *
	// CREATING: The environment is in the process of being created.
	//
	//     * CREATED:
	// The environment was successfully created.
	//
	//     * CREATE_FAILED: The environment
	// failed to be created.
	//
	//     * DELETING: The environment is in the process of
	// being deleted.
	//
	//     * DELETE_FAILED: The environment failed to delete.
	Status EnvironmentLifecycleStatus
	// If the environment failed to delete, the Amazon Resource Name (ARN) of the
	// related AWS resource.
	FailureResource *string
	// Any informational message about the lifecycle state of the environment.
	Reason *string
}

// Information about an environment member for an AWS Cloud9 development
// environment.
type EnvironmentMember struct {
	// The user ID in AWS Identity and Access Management (AWS IAM) of the environment
	// member.
	UserId *string
	// The time, expressed in epoch time format, when the environment member last
	// opened the environment.
	LastAccess *time.Time
	// The ID of the environment for the environment member.
	EnvironmentId *string
	// The Amazon Resource Name (ARN) of the environment member.
	UserArn *string
	// The type of environment member permissions associated with this environment
	// member. Available values include:
	//
	//     * owner: Owns the environment.
	//
	//     *
	// read-only: Has read-only access to the environment.
	//
	//     * read-write: Has
	// read-write access to the environment.
	Permissions Permissions
}

// Metadata that is associated with AWS resources. In particular, a name-value pair
// that can be associated with an AWS Cloud9 development environment. There are two
// types of tags: user tags and system tags. A user tag is created by the user. A
// system tag is automatically created by AWS services. A system tag is prefixed
// with "aws:" and cannot be modified by the user.
type Tag struct {
	// The value part of a tag.
	Value *string
	// The name part of a tag.
	Key *string
}
