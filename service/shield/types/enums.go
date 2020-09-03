// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AttackLayer string

// Enum values for AttackLayer
const (
	AttackLayerNetwork     AttackLayer = "NETWORK"
	AttackLayerApplication AttackLayer = "APPLICATION"
)

type AttackPropertyIdentifier string

// Enum values for AttackPropertyIdentifier
const (
	AttackPropertyIdentifierDestination_url              AttackPropertyIdentifier = "DESTINATION_URL"
	AttackPropertyIdentifierReferrer                     AttackPropertyIdentifier = "REFERRER"
	AttackPropertyIdentifierSource_asn                   AttackPropertyIdentifier = "SOURCE_ASN"
	AttackPropertyIdentifierSource_country               AttackPropertyIdentifier = "SOURCE_COUNTRY"
	AttackPropertyIdentifierSource_ip_address            AttackPropertyIdentifier = "SOURCE_IP_ADDRESS"
	AttackPropertyIdentifierSource_user_agent            AttackPropertyIdentifier = "SOURCE_USER_AGENT"
	AttackPropertyIdentifierWordpress_pingback_reflector AttackPropertyIdentifier = "WORDPRESS_PINGBACK_REFLECTOR"
	AttackPropertyIdentifierWordpress_pingback_source    AttackPropertyIdentifier = "WORDPRESS_PINGBACK_SOURCE"
)

type AutoRenew string

// Enum values for AutoRenew
const (
	AutoRenewEnabled  AutoRenew = "ENABLED"
	AutoRenewDisabled AutoRenew = "DISABLED"
)

type ProactiveEngagementStatus string

// Enum values for ProactiveEngagementStatus
const (
	ProactiveEngagementStatusEnabled  ProactiveEngagementStatus = "ENABLED"
	ProactiveEngagementStatusDisabled ProactiveEngagementStatus = "DISABLED"
	ProactiveEngagementStatusPending  ProactiveEngagementStatus = "PENDING"
)

type SubResourceType string

// Enum values for SubResourceType
const (
	SubResourceTypeIp  SubResourceType = "IP"
	SubResourceTypeUrl SubResourceType = "URL"
)

type SubscriptionState string

// Enum values for SubscriptionState
const (
	SubscriptionStateActive   SubscriptionState = "ACTIVE"
	SubscriptionStateInactive SubscriptionState = "INACTIVE"
)

type Unit string

// Enum values for Unit
const (
	UnitBits     Unit = "BITS"
	UnitBytes    Unit = "BYTES"
	UnitPackets  Unit = "PACKETS"
	UnitRequests Unit = "REQUESTS"
)
