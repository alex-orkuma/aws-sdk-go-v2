// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Edition string

// Enum values for Edition
const (
	EditionStarter  Edition = "STARTER"
	EditionStandard Edition = "STANDARD"
)

type Framework string

// Enum values for Framework
const (
	FrameworkHyperledger_fabric Framework = "HYPERLEDGER_FABRIC"
)

type InvitationStatus string

// Enum values for InvitationStatus
const (
	InvitationStatusPending   InvitationStatus = "PENDING"
	InvitationStatusAccepted  InvitationStatus = "ACCEPTED"
	InvitationStatusAccepting InvitationStatus = "ACCEPTING"
	InvitationStatusRejected  InvitationStatus = "REJECTED"
	InvitationStatusExpired   InvitationStatus = "EXPIRED"
)

type MemberStatus string

// Enum values for MemberStatus
const (
	MemberStatusCreating      MemberStatus = "CREATING"
	MemberStatusAvailable     MemberStatus = "AVAILABLE"
	MemberStatusCreate_failed MemberStatus = "CREATE_FAILED"
	MemberStatusUpdating      MemberStatus = "UPDATING"
	MemberStatusDeleting      MemberStatus = "DELETING"
	MemberStatusDeleted       MemberStatus = "DELETED"
)

type NetworkStatus string

// Enum values for NetworkStatus
const (
	NetworkStatusCreating      NetworkStatus = "CREATING"
	NetworkStatusAvailable     NetworkStatus = "AVAILABLE"
	NetworkStatusCreate_failed NetworkStatus = "CREATE_FAILED"
	NetworkStatusDeleting      NetworkStatus = "DELETING"
	NetworkStatusDeleted       NetworkStatus = "DELETED"
)

type NodeStatus string

// Enum values for NodeStatus
const (
	NodeStatusCreating      NodeStatus = "CREATING"
	NodeStatusAvailable     NodeStatus = "AVAILABLE"
	NodeStatusCreate_failed NodeStatus = "CREATE_FAILED"
	NodeStatusUpdating      NodeStatus = "UPDATING"
	NodeStatusDeleting      NodeStatus = "DELETING"
	NodeStatusDeleted       NodeStatus = "DELETED"
	NodeStatusFailed        NodeStatus = "FAILED"
)

type ProposalStatus string

// Enum values for ProposalStatus
const (
	ProposalStatusIn_progress   ProposalStatus = "IN_PROGRESS"
	ProposalStatusApproved      ProposalStatus = "APPROVED"
	ProposalStatusRejected      ProposalStatus = "REJECTED"
	ProposalStatusExpired       ProposalStatus = "EXPIRED"
	ProposalStatusAction_failed ProposalStatus = "ACTION_FAILED"
)

type ThresholdComparator string

// Enum values for ThresholdComparator
const (
	ThresholdComparatorGreater_than             ThresholdComparator = "GREATER_THAN"
	ThresholdComparatorGreater_than_or_equal_to ThresholdComparator = "GREATER_THAN_OR_EQUAL_TO"
)

type VoteValue string

// Enum values for VoteValue
const (
	VoteValueYes VoteValue = "YES"
	VoteValueNo  VoteValue = "NO"
)
