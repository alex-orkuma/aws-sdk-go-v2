// Code generated by smithy-go-codegen DO NOT EDIT.

package managedblockchain

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpCreateMember struct {
}

func (*validateOpCreateMember) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateMember) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateMemberInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateMemberInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateNetwork struct {
}

func (*validateOpCreateNetwork) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateNetwork) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateNetworkInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateNetworkInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateNode struct {
}

func (*validateOpCreateNode) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateNode) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateNodeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateNodeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateProposal struct {
}

func (*validateOpCreateProposal) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateProposal) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateProposalInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateProposalInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteMember struct {
}

func (*validateOpDeleteMember) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteMember) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteMemberInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteMemberInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteNode struct {
}

func (*validateOpDeleteNode) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteNode) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteNodeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteNodeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetMember struct {
}

func (*validateOpGetMember) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetMember) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetMemberInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetMemberInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetNetwork struct {
}

func (*validateOpGetNetwork) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetNetwork) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetNetworkInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetNetworkInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetNode struct {
}

func (*validateOpGetNode) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetNode) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetNodeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetNodeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetProposal struct {
}

func (*validateOpGetProposal) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetProposal) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetProposalInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetProposalInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListMembers struct {
}

func (*validateOpListMembers) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListMembers) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListMembersInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListMembersInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListNodes struct {
}

func (*validateOpListNodes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListNodes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListNodesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListNodesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListProposals struct {
}

func (*validateOpListProposals) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListProposals) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListProposalsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListProposalsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListProposalVotes struct {
}

func (*validateOpListProposalVotes) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListProposalVotes) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListProposalVotesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListProposalVotesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpRejectInvitation struct {
}

func (*validateOpRejectInvitation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpRejectInvitation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*RejectInvitationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpRejectInvitationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateMember struct {
}

func (*validateOpUpdateMember) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateMember) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateMemberInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateMemberInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateNode struct {
}

func (*validateOpUpdateNode) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateNode) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateNodeInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateNodeInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpVoteOnProposal struct {
}

func (*validateOpVoteOnProposal) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpVoteOnProposal) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*VoteOnProposalInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpVoteOnProposalInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpCreateMemberValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateMember{}, middleware.After)
}

func addOpCreateNetworkValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateNetwork{}, middleware.After)
}

func addOpCreateNodeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateNode{}, middleware.After)
}

func addOpCreateProposalValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateProposal{}, middleware.After)
}

func addOpDeleteMemberValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteMember{}, middleware.After)
}

func addOpDeleteNodeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteNode{}, middleware.After)
}

func addOpGetMemberValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetMember{}, middleware.After)
}

func addOpGetNetworkValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetNetwork{}, middleware.After)
}

func addOpGetNodeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetNode{}, middleware.After)
}

func addOpGetProposalValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetProposal{}, middleware.After)
}

func addOpListMembersValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListMembers{}, middleware.After)
}

func addOpListNodesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListNodes{}, middleware.After)
}

func addOpListProposalsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListProposals{}, middleware.After)
}

func addOpListProposalVotesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListProposalVotes{}, middleware.After)
}

func addOpRejectInvitationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpRejectInvitation{}, middleware.After)
}

func addOpUpdateMemberValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateMember{}, middleware.After)
}

func addOpUpdateNodeValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateNode{}, middleware.After)
}

func addOpVoteOnProposalValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpVoteOnProposal{}, middleware.After)
}

func validateInviteAction(v *types.InviteAction) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InviteAction"}
	if v.Principal == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Principal"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateInviteActionList(v []*types.InviteAction) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "InviteActionList"}
	for i := range v {
		if err := validateInviteAction(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMemberConfiguration(v *types.MemberConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MemberConfiguration"}
	if v.FrameworkConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FrameworkConfiguration"))
	} else if v.FrameworkConfiguration != nil {
		if err := validateMemberFrameworkConfiguration(v.FrameworkConfiguration); err != nil {
			invalidParams.AddNested("FrameworkConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMemberFabricConfiguration(v *types.MemberFabricConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MemberFabricConfiguration"}
	if v.AdminPassword == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdminPassword"))
	}
	if v.AdminUsername == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AdminUsername"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateMemberFrameworkConfiguration(v *types.MemberFrameworkConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "MemberFrameworkConfiguration"}
	if v.Fabric != nil {
		if err := validateMemberFabricConfiguration(v.Fabric); err != nil {
			invalidParams.AddNested("Fabric", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNetworkFabricConfiguration(v *types.NetworkFabricConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NetworkFabricConfiguration"}
	if len(v.Edition) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Edition"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNetworkFrameworkConfiguration(v *types.NetworkFrameworkConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NetworkFrameworkConfiguration"}
	if v.Fabric != nil {
		if err := validateNetworkFabricConfiguration(v.Fabric); err != nil {
			invalidParams.AddNested("Fabric", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateNodeConfiguration(v *types.NodeConfiguration) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "NodeConfiguration"}
	if v.AvailabilityZone == nil {
		invalidParams.Add(smithy.NewErrParamRequired("AvailabilityZone"))
	}
	if v.InstanceType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InstanceType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateProposalActions(v *types.ProposalActions) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ProposalActions"}
	if v.Removals != nil {
		if err := validateRemoveActionList(v.Removals); err != nil {
			invalidParams.AddNested("Removals", err.(smithy.InvalidParamsError))
		}
	}
	if v.Invitations != nil {
		if err := validateInviteActionList(v.Invitations); err != nil {
			invalidParams.AddNested("Invitations", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRemoveAction(v *types.RemoveAction) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveAction"}
	if v.MemberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateRemoveActionList(v []*types.RemoveAction) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RemoveActionList"}
	for i := range v {
		if err := validateRemoveAction(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateMemberInput(v *CreateMemberInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateMemberInput"}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if v.MemberConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberConfiguration"))
	} else if v.MemberConfiguration != nil {
		if err := validateMemberConfiguration(v.MemberConfiguration); err != nil {
			invalidParams.AddNested("MemberConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.InvitationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InvitationId"))
	}
	if v.ClientRequestToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientRequestToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateNetworkInput(v *CreateNetworkInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateNetworkInput"}
	if v.ClientRequestToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientRequestToken"))
	}
	if v.Name == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Name"))
	}
	if v.MemberConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberConfiguration"))
	} else if v.MemberConfiguration != nil {
		if err := validateMemberConfiguration(v.MemberConfiguration); err != nil {
			invalidParams.AddNested("MemberConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.VotingPolicy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VotingPolicy"))
	}
	if len(v.Framework) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Framework"))
	}
	if v.FrameworkConfiguration != nil {
		if err := validateNetworkFrameworkConfiguration(v.FrameworkConfiguration); err != nil {
			invalidParams.AddNested("FrameworkConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.FrameworkVersion == nil {
		invalidParams.Add(smithy.NewErrParamRequired("FrameworkVersion"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateNodeInput(v *CreateNodeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateNodeInput"}
	if v.ClientRequestToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientRequestToken"))
	}
	if v.MemberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberId"))
	}
	if v.NodeConfiguration == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NodeConfiguration"))
	} else if v.NodeConfiguration != nil {
		if err := validateNodeConfiguration(v.NodeConfiguration); err != nil {
			invalidParams.AddNested("NodeConfiguration", err.(smithy.InvalidParamsError))
		}
	}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateProposalInput(v *CreateProposalInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateProposalInput"}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if v.ClientRequestToken == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ClientRequestToken"))
	}
	if v.MemberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberId"))
	}
	if v.Actions == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Actions"))
	} else if v.Actions != nil {
		if err := validateProposalActions(v.Actions); err != nil {
			invalidParams.AddNested("Actions", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteMemberInput(v *DeleteMemberInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteMemberInput"}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if v.MemberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteNodeInput(v *DeleteNodeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteNodeInput"}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if v.NodeId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NodeId"))
	}
	if v.MemberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetMemberInput(v *GetMemberInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetMemberInput"}
	if v.MemberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberId"))
	}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetNetworkInput(v *GetNetworkInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetNetworkInput"}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetNodeInput(v *GetNodeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetNodeInput"}
	if v.NodeId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NodeId"))
	}
	if v.MemberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberId"))
	}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetProposalInput(v *GetProposalInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetProposalInput"}
	if v.ProposalId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProposalId"))
	}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListMembersInput(v *ListMembersInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListMembersInput"}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListNodesInput(v *ListNodesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListNodesInput"}
	if v.MemberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberId"))
	}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListProposalsInput(v *ListProposalsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListProposalsInput"}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListProposalVotesInput(v *ListProposalVotesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListProposalVotesInput"}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if v.ProposalId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProposalId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpRejectInvitationInput(v *RejectInvitationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "RejectInvitationInput"}
	if v.InvitationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InvitationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateMemberInput(v *UpdateMemberInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateMemberInput"}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if v.MemberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateNodeInput(v *UpdateNodeInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateNodeInput"}
	if v.NodeId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NodeId"))
	}
	if v.MemberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("MemberId"))
	}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpVoteOnProposalInput(v *VoteOnProposalInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "VoteOnProposalInput"}
	if v.NetworkId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("NetworkId"))
	}
	if len(v.Vote) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Vote"))
	}
	if v.VoterMemberId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VoterMemberId"))
	}
	if v.ProposalId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ProposalId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
