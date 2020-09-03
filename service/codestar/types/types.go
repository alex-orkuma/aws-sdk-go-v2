// Code generated by smithy-go-codegen DO NOT EDIT.

package types

// Location and destination information about the source code files provided with
// the project request. The source code is uploaded to the new project source
// repository after project creation.
type Code struct {
	// The location where the source code files provided with the project request are
	// stored. AWS CodeStar retrieves the files during project creation.
	Source *CodeSource
	// The repository to be created in AWS CodeStar. Valid values are AWS CodeCommit or
	// GitHub. After AWS CodeStar provisions the new repository, the source code files
	// provided with the project request are placed in the repository.
	Destination *CodeDestination
}

// Information about the AWS CodeCommit repository to be created in AWS CodeStar.
// This is where the source code files provided with the project request will be
// uploaded after project creation.
type CodeCommitCodeDestination struct {
	// The name of the AWS CodeCommit repository to be created in AWS CodeStar.
	Name *string
}

// The repository to be created in AWS CodeStar. Valid values are AWS CodeCommit or
// GitHub. After AWS CodeStar provisions the new repository, the source code files
// provided with the project request are placed in the repository.
type CodeDestination struct {
	// Information about the GitHub repository to be created in AWS CodeStar. This is
	// where the source code files provided with the project request will be uploaded
	// after project creation.
	GitHub *GitHubCodeDestination
	// Information about the AWS CodeCommit repository to be created in AWS CodeStar.
	// This is where the source code files provided with the project request will be
	// uploaded after project creation.
	CodeCommit *CodeCommitCodeDestination
}

// The location where the source code files provided with the project request are
// stored. AWS CodeStar retrieves the files during project creation.
type CodeSource struct {
	// Information about the Amazon S3 location where the source code files provided
	// with the project request are stored.
	S3 *S3Location
}

// Information about the GitHub repository to be created in AWS CodeStar. This is
// where the source code files provided with the project request will be uploaded
// after project creation.
type GitHubCodeDestination struct {
	// The type of GitHub repository to be created in AWS CodeStar. Valid values are
	// User or Organization.
	Type *string
	// Whether the GitHub repository is to be a private repository.
	PrivateRepository *bool
	// The GitHub user's personal access token for the GitHub repository.
	Token *string
	// Name of the GitHub repository to be created in AWS CodeStar.
	Name *string
	// Whether to enable issues for the GitHub repository.
	IssuesEnabled *bool
	// The GitHub username for the owner of the GitHub repository to be created in AWS
	// CodeStar. If this repository should be owned by a GitHub organization, provide
	// its name.
	Owner *string
	// Description for the GitHub repository to be created in AWS CodeStar. This
	// description displays in GitHub after the repository is created.
	Description *string
}

// An indication of whether a project creation or deletion is failed or successful.
type ProjectStatus struct {
	// The phase of completion for a project creation or deletion.
	State *string
	// In the case of a project creation or deletion failure, a reason for the failure.
	Reason *string
}

// Information about the metadata for a project.
type ProjectSummary struct {
	// The ID of the project.
	ProjectId *string
	// The Amazon Resource Name (ARN) of the project.
	ProjectArn *string
}

// Information about a resource for a project.
type Resource struct {
	// The Amazon Resource Name (ARN) of the resource.
	Id *string
}

// The Amazon S3 location where the source code files provided with the project
// request are stored.
type S3Location struct {
	// The Amazon S3 object key where the source code files provided with the project
	// request are stored.
	BucketKey *string
	// The Amazon S3 bucket name where the source code files provided with the project
	// request are stored.
	BucketName *string
}

// Information about a team member in a project.
type TeamMember struct {
	// The role assigned to the user in the project. Project roles have different
	// levels of access. For more information, see Working with Teams
	// (http://docs.aws.amazon.com/codestar/latest/userguide/working-with-teams.html)
	// in the AWS CodeStar User Guide.
	ProjectRole *string
	// The Amazon Resource Name (ARN) of the user in IAM.
	UserArn *string
	// Whether the user is allowed to remotely access project resources using an SSH
	// public/private key pair.
	RemoteAccessAllowed *bool
}

// The toolchain template file provided with the project request. AWS CodeStar uses
// the template to provision the toolchain stack in AWS CloudFormation.
type Toolchain struct {
	// The Amazon S3 location where the toolchain template file provided with the
	// project request is stored. AWS CodeStar retrieves the file during project
	// creation.
	Source *ToolchainSource
	// The list of parameter overrides to be passed into the toolchain template during
	// stack provisioning, if any.
	StackParameters map[string]*string
	// The service role ARN for AWS CodeStar to use for the toolchain template during
	// stack provisioning.
	RoleArn *string
}

// The Amazon S3 location where the toolchain template file provided with the
// project request is stored. AWS CodeStar retrieves the file during project
// creation.
type ToolchainSource struct {
	// The Amazon S3 bucket where the toolchain template file provided with the project
	// request is stored.
	S3 *S3Location
}

// Information about a user's profile in AWS CodeStar.
type UserProfileSummary struct {
	// The email address associated with the user.
	EmailAddress *string
	// The display name of a user in AWS CodeStar. For example, this could be set to
	// both first and last name ("Mary Major") or a single name ("Mary"). The display
	// name is also used to generate the initial icon associated with the user in AWS
	// CodeStar projects. If spaces are included in the display name, the first
	// character that appears after the space will be used as the second character in
	// the user initial icon. The initial icon displays a maximum of two characters, so
	// a display name with more than one space (for example "Mary Jane Major") would
	// generate an initial icon using the first character and the first character after
	// the space ("MJ", not "MM").
	DisplayName *string
	// The Amazon Resource Name (ARN) of the user in IAM.
	UserArn *string
	// The SSH public key associated with the user in AWS CodeStar. If a project owner
	// allows the user remote access to project resources, this public key will be used
	// along with the user's private key for SSH access.
	SshPublicKey *string
}
