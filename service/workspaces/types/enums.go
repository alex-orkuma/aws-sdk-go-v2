// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccessPropertyValue string

// Enum values for AccessPropertyValue
const (
	AccessPropertyValueAllow AccessPropertyValue = "ALLOW"
	AccessPropertyValueDeny  AccessPropertyValue = "DENY"
)

type Compute string

// Enum values for Compute
const (
	ComputeValue       Compute = "VALUE"
	ComputeStandard    Compute = "STANDARD"
	ComputePerformance Compute = "PERFORMANCE"
	ComputePower       Compute = "POWER"
	ComputeGraphics    Compute = "GRAPHICS"
	ComputePowerpro    Compute = "POWERPRO"
	ComputeGraphicspro Compute = "GRAPHICSPRO"
)

type ConnectionState string

// Enum values for ConnectionState
const (
	ConnectionStateConnected    ConnectionState = "CONNECTED"
	ConnectionStateDisconnected ConnectionState = "DISCONNECTED"
	ConnectionStateUnknown      ConnectionState = "UNKNOWN"
)

type DedicatedTenancyModificationStateEnum string

// Enum values for DedicatedTenancyModificationStateEnum
const (
	DedicatedTenancyModificationStateEnumPending   DedicatedTenancyModificationStateEnum = "PENDING"
	DedicatedTenancyModificationStateEnumCompleted DedicatedTenancyModificationStateEnum = "COMPLETED"
	DedicatedTenancyModificationStateEnumFailed    DedicatedTenancyModificationStateEnum = "FAILED"
)

type DedicatedTenancySupportEnum string

// Enum values for DedicatedTenancySupportEnum
const (
	DedicatedTenancySupportEnumEnabled DedicatedTenancySupportEnum = "ENABLED"
)

type DedicatedTenancySupportResultEnum string

// Enum values for DedicatedTenancySupportResultEnum
const (
	DedicatedTenancySupportResultEnumEnabled  DedicatedTenancySupportResultEnum = "ENABLED"
	DedicatedTenancySupportResultEnumDisabled DedicatedTenancySupportResultEnum = "DISABLED"
)

type ImageType string

// Enum values for ImageType
const (
	ImageTypeOwned  ImageType = "OWNED"
	ImageTypeShared ImageType = "SHARED"
)

type ModificationResourceEnum string

// Enum values for ModificationResourceEnum
const (
	ModificationResourceEnumRoot_volume  ModificationResourceEnum = "ROOT_VOLUME"
	ModificationResourceEnumUser_volume  ModificationResourceEnum = "USER_VOLUME"
	ModificationResourceEnumCompute_type ModificationResourceEnum = "COMPUTE_TYPE"
)

type ModificationStateEnum string

// Enum values for ModificationStateEnum
const (
	ModificationStateEnumUpdate_initiated   ModificationStateEnum = "UPDATE_INITIATED"
	ModificationStateEnumUpdate_in_progress ModificationStateEnum = "UPDATE_IN_PROGRESS"
)

type OperatingSystemType string

// Enum values for OperatingSystemType
const (
	OperatingSystemTypeWindows OperatingSystemType = "WINDOWS"
	OperatingSystemTypeLinux   OperatingSystemType = "LINUX"
)

type ReconnectEnum string

// Enum values for ReconnectEnum
const (
	ReconnectEnumEnabled  ReconnectEnum = "ENABLED"
	ReconnectEnumDisabled ReconnectEnum = "DISABLED"
)

type RunningMode string

// Enum values for RunningMode
const (
	RunningModeAuto_stop RunningMode = "AUTO_STOP"
	RunningModeAlways_on RunningMode = "ALWAYS_ON"
)

type TargetWorkspaceState string

// Enum values for TargetWorkspaceState
const (
	TargetWorkspaceStateAvailable         TargetWorkspaceState = "AVAILABLE"
	TargetWorkspaceStateAdmin_maintenance TargetWorkspaceState = "ADMIN_MAINTENANCE"
)

type Tenancy string

// Enum values for Tenancy
const (
	TenancyDedicated Tenancy = "DEDICATED"
	TenancyShared    Tenancy = "SHARED"
)

type WorkspaceDirectoryState string

// Enum values for WorkspaceDirectoryState
const (
	WorkspaceDirectoryStateRegistering   WorkspaceDirectoryState = "REGISTERING"
	WorkspaceDirectoryStateRegistered    WorkspaceDirectoryState = "REGISTERED"
	WorkspaceDirectoryStateDeregistering WorkspaceDirectoryState = "DEREGISTERING"
	WorkspaceDirectoryStateDeregistered  WorkspaceDirectoryState = "DEREGISTERED"
	WorkspaceDirectoryStateError         WorkspaceDirectoryState = "ERROR"
)

type WorkspaceDirectoryType string

// Enum values for WorkspaceDirectoryType
const (
	WorkspaceDirectoryTypeSimple_ad    WorkspaceDirectoryType = "SIMPLE_AD"
	WorkspaceDirectoryTypeAd_connector WorkspaceDirectoryType = "AD_CONNECTOR"
)

type WorkspaceImageIngestionProcess string

// Enum values for WorkspaceImageIngestionProcess
const (
	WorkspaceImageIngestionProcessByol_regular     WorkspaceImageIngestionProcess = "BYOL_REGULAR"
	WorkspaceImageIngestionProcessByol_graphics    WorkspaceImageIngestionProcess = "BYOL_GRAPHICS"
	WorkspaceImageIngestionProcessByol_graphicspro WorkspaceImageIngestionProcess = "BYOL_GRAPHICSPRO"
)

type WorkspaceImageRequiredTenancy string

// Enum values for WorkspaceImageRequiredTenancy
const (
	WorkspaceImageRequiredTenancyDefault   WorkspaceImageRequiredTenancy = "DEFAULT"
	WorkspaceImageRequiredTenancyDedicated WorkspaceImageRequiredTenancy = "DEDICATED"
)

type WorkspaceImageState string

// Enum values for WorkspaceImageState
const (
	WorkspaceImageStateAvailable WorkspaceImageState = "AVAILABLE"
	WorkspaceImageStatePending   WorkspaceImageState = "PENDING"
	WorkspaceImageStateError     WorkspaceImageState = "ERROR"
)

type WorkspaceState string

// Enum values for WorkspaceState
const (
	WorkspaceStatePending           WorkspaceState = "PENDING"
	WorkspaceStateAvailable         WorkspaceState = "AVAILABLE"
	WorkspaceStateImpaired          WorkspaceState = "IMPAIRED"
	WorkspaceStateUnhealthy         WorkspaceState = "UNHEALTHY"
	WorkspaceStateRebooting         WorkspaceState = "REBOOTING"
	WorkspaceStateStarting          WorkspaceState = "STARTING"
	WorkspaceStateRebuilding        WorkspaceState = "REBUILDING"
	WorkspaceStateRestoring         WorkspaceState = "RESTORING"
	WorkspaceStateMaintenance       WorkspaceState = "MAINTENANCE"
	WorkspaceStateAdmin_maintenance WorkspaceState = "ADMIN_MAINTENANCE"
	WorkspaceStateTerminating       WorkspaceState = "TERMINATING"
	WorkspaceStateTerminated        WorkspaceState = "TERMINATED"
	WorkspaceStateSuspended         WorkspaceState = "SUSPENDED"
	WorkspaceStateUpdating          WorkspaceState = "UPDATING"
	WorkspaceStateStopping          WorkspaceState = "STOPPING"
	WorkspaceStateStopped           WorkspaceState = "STOPPED"
	WorkspaceStateError             WorkspaceState = "ERROR"
)
