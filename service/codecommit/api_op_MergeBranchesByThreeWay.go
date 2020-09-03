// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Merges two specified branches using the three-way merge strategy.
func (c *Client) MergeBranchesByThreeWay(ctx context.Context, params *MergeBranchesByThreeWayInput, optFns ...func(*Options)) (*MergeBranchesByThreeWayOutput, error) {
	stack := middleware.NewStack("MergeBranchesByThreeWay", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpMergeBranchesByThreeWayMiddlewares(stack)
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
	addOpMergeBranchesByThreeWayValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opMergeBranchesByThreeWay(options.Region), middleware.Before)

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
			OperationName: "MergeBranchesByThreeWay",
			Err:           err,
		}
	}
	out := result.(*MergeBranchesByThreeWayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type MergeBranchesByThreeWayInput struct {
	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	SourceCommitSpecifier *string
	// The name of the repository where you want to merge two branches.
	RepositoryName *string
	// If AUTOMERGE is the conflict resolution strategy, a list of inputs to use when
	// resolving conflicts during a merge.
	ConflictResolution *types.ConflictResolution
	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation is
	// successful.
	ConflictResolutionStrategy types.ConflictResolutionStrategyTypeEnum
	// If the commit contains deletions, whether to keep a folder or folder structure
	// if the changes leave the folders empty. If true, a .gitkeep file is created for
	// empty folders. The default is false.
	KeepEmptyFolders *bool
	// The email address of the person merging the branches. This information is used
	// in the commit information for the merge.
	Email *string
	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL is
	// used, which returns a not-mergeable result if the same file has differences in
	// both branches. If LINE_LEVEL is specified, a conflict is considered not
	// mergeable if the same file in both branches has differences on the same line.
	ConflictDetailLevel types.ConflictDetailLevelTypeEnum
	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	DestinationCommitSpecifier *string
	// The branch where the merge is applied.
	TargetBranch *string
	// The name of the author who created the commit. This information is used as both
	// the author and committer for the commit.
	AuthorName *string
	// The commit message to include in the commit information for the merge.
	CommitMessage *string
}

type MergeBranchesByThreeWayOutput struct {
	// The tree ID of the merge in the destination or target branch.
	TreeId *string
	// The commit ID of the merge in the destination or target branch.
	CommitId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpMergeBranchesByThreeWayMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpMergeBranchesByThreeWay{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpMergeBranchesByThreeWay{}, middleware.After)
}

func newServiceMetadataMiddleware_opMergeBranchesByThreeWay(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "MergeBranchesByThreeWay",
	}
}
