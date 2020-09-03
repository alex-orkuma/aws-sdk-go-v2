// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ResourceOwner string

// Enum values for ResourceOwner
const (
	ResourceOwnerSelf           ResourceOwner = "SELF"
	ResourceOwnerOther_accounts ResourceOwner = "OTHER-ACCOUNTS"
)

type ResourceShareAssociationStatus string

// Enum values for ResourceShareAssociationStatus
const (
	ResourceShareAssociationStatusAssociating    ResourceShareAssociationStatus = "ASSOCIATING"
	ResourceShareAssociationStatusAssociated     ResourceShareAssociationStatus = "ASSOCIATED"
	ResourceShareAssociationStatusFailed         ResourceShareAssociationStatus = "FAILED"
	ResourceShareAssociationStatusDisassociating ResourceShareAssociationStatus = "DISASSOCIATING"
	ResourceShareAssociationStatusDisassociated  ResourceShareAssociationStatus = "DISASSOCIATED"
)

type ResourceShareAssociationType string

// Enum values for ResourceShareAssociationType
const (
	ResourceShareAssociationTypePrincipal ResourceShareAssociationType = "PRINCIPAL"
	ResourceShareAssociationTypeResource  ResourceShareAssociationType = "RESOURCE"
)

type ResourceShareFeatureSet string

// Enum values for ResourceShareFeatureSet
const (
	ResourceShareFeatureSetCreated_from_policy   ResourceShareFeatureSet = "CREATED_FROM_POLICY"
	ResourceShareFeatureSetPromoting_to_standard ResourceShareFeatureSet = "PROMOTING_TO_STANDARD"
	ResourceShareFeatureSetStandard              ResourceShareFeatureSet = "STANDARD"
)

type ResourceShareInvitationStatus string

// Enum values for ResourceShareInvitationStatus
const (
	ResourceShareInvitationStatusPending  ResourceShareInvitationStatus = "PENDING"
	ResourceShareInvitationStatusAccepted ResourceShareInvitationStatus = "ACCEPTED"
	ResourceShareInvitationStatusRejected ResourceShareInvitationStatus = "REJECTED"
	ResourceShareInvitationStatusExpired  ResourceShareInvitationStatus = "EXPIRED"
)

type ResourceShareStatus string

// Enum values for ResourceShareStatus
const (
	ResourceShareStatusPending  ResourceShareStatus = "PENDING"
	ResourceShareStatusActive   ResourceShareStatus = "ACTIVE"
	ResourceShareStatusFailed   ResourceShareStatus = "FAILED"
	ResourceShareStatusDeleting ResourceShareStatus = "DELETING"
	ResourceShareStatusDeleted  ResourceShareStatus = "DELETED"
)

type ResourceStatus string

// Enum values for ResourceStatus
const (
	ResourceStatusAvailable                   ResourceStatus = "AVAILABLE"
	ResourceStatusZonal_resource_inaccessible ResourceStatus = "ZONAL_RESOURCE_INACCESSIBLE"
	ResourceStatusLimit_exceeded              ResourceStatus = "LIMIT_EXCEEDED"
	ResourceStatusUnavailable                 ResourceStatus = "UNAVAILABLE"
	ResourceStatusPending                     ResourceStatus = "PENDING"
)
