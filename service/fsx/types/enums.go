// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActiveDirectoryErrorType string

// Enum values for ActiveDirectoryErrorType
const (
	ActiveDirectoryErrorTypeDomain_not_found         ActiveDirectoryErrorType = "DOMAIN_NOT_FOUND"
	ActiveDirectoryErrorTypeIncompatible_domain_mode ActiveDirectoryErrorType = "INCOMPATIBLE_DOMAIN_MODE"
	ActiveDirectoryErrorTypeWrong_vpc                ActiveDirectoryErrorType = "WRONG_VPC"
	ActiveDirectoryErrorTypeInvalid_domain_stage     ActiveDirectoryErrorType = "INVALID_DOMAIN_STAGE"
)

type AdministrativeActionType string

// Enum values for AdministrativeActionType
const (
	AdministrativeActionTypeFile_system_update   AdministrativeActionType = "FILE_SYSTEM_UPDATE"
	AdministrativeActionTypeStorage_optimization AdministrativeActionType = "STORAGE_OPTIMIZATION"
)

type AutoImportPolicyType string

// Enum values for AutoImportPolicyType
const (
	AutoImportPolicyTypeNone        AutoImportPolicyType = "NONE"
	AutoImportPolicyTypeNew         AutoImportPolicyType = "NEW"
	AutoImportPolicyTypeNew_changed AutoImportPolicyType = "NEW_CHANGED"
)

type BackupLifecycle string

// Enum values for BackupLifecycle
const (
	BackupLifecycleAvailable BackupLifecycle = "AVAILABLE"
	BackupLifecycleCreating  BackupLifecycle = "CREATING"
	BackupLifecycleDeleted   BackupLifecycle = "DELETED"
	BackupLifecycleFailed    BackupLifecycle = "FAILED"
)

type BackupType string

// Enum values for BackupType
const (
	BackupTypeAutomatic      BackupType = "AUTOMATIC"
	BackupTypeUser_initiated BackupType = "USER_INITIATED"
)

type DataRepositoryLifecycle string

// Enum values for DataRepositoryLifecycle
const (
	DataRepositoryLifecycleCreating      DataRepositoryLifecycle = "CREATING"
	DataRepositoryLifecycleAvailable     DataRepositoryLifecycle = "AVAILABLE"
	DataRepositoryLifecycleMisconfigured DataRepositoryLifecycle = "MISCONFIGURED"
	DataRepositoryLifecycleUpdating      DataRepositoryLifecycle = "UPDATING"
	DataRepositoryLifecycleDeleting      DataRepositoryLifecycle = "DELETING"
)

type DataRepositoryTaskFilterName string

// Enum values for DataRepositoryTaskFilterName
const (
	DataRepositoryTaskFilterNameFile_system_id DataRepositoryTaskFilterName = "file-system-id"
	DataRepositoryTaskFilterNameTask_lifecycle DataRepositoryTaskFilterName = "task-lifecycle"
)

type DataRepositoryTaskLifecycle string

// Enum values for DataRepositoryTaskLifecycle
const (
	DataRepositoryTaskLifecyclePending   DataRepositoryTaskLifecycle = "PENDING"
	DataRepositoryTaskLifecycleExecuting DataRepositoryTaskLifecycle = "EXECUTING"
	DataRepositoryTaskLifecycleFailed    DataRepositoryTaskLifecycle = "FAILED"
	DataRepositoryTaskLifecycleSucceeded DataRepositoryTaskLifecycle = "SUCCEEDED"
	DataRepositoryTaskLifecycleCanceled  DataRepositoryTaskLifecycle = "CANCELED"
	DataRepositoryTaskLifecycleCanceling DataRepositoryTaskLifecycle = "CANCELING"
)

type DataRepositoryTaskType string

// Enum values for DataRepositoryTaskType
const (
	DataRepositoryTaskTypeExport DataRepositoryTaskType = "EXPORT_TO_REPOSITORY"
)

type FileSystemLifecycle string

// Enum values for FileSystemLifecycle
const (
	FileSystemLifecycleAvailable     FileSystemLifecycle = "AVAILABLE"
	FileSystemLifecycleCreating      FileSystemLifecycle = "CREATING"
	FileSystemLifecycleFailed        FileSystemLifecycle = "FAILED"
	FileSystemLifecycleDeleting      FileSystemLifecycle = "DELETING"
	FileSystemLifecycleMisconfigured FileSystemLifecycle = "MISCONFIGURED"
	FileSystemLifecycleUpdating      FileSystemLifecycle = "UPDATING"
)

type FileSystemMaintenanceOperation string

// Enum values for FileSystemMaintenanceOperation
const (
	FileSystemMaintenanceOperationPatching   FileSystemMaintenanceOperation = "PATCHING"
	FileSystemMaintenanceOperationBacking_up FileSystemMaintenanceOperation = "BACKING_UP"
)

type FileSystemType string

// Enum values for FileSystemType
const (
	FileSystemTypeWindows FileSystemType = "WINDOWS"
	FileSystemTypeLustre  FileSystemType = "LUSTRE"
)

type FilterName string

// Enum values for FilterName
const (
	FilterNameFile_system_id   FilterName = "file-system-id"
	FilterNameBackup_type      FilterName = "backup-type"
	FilterNameFile_system_type FilterName = "file-system-type"
)

type LustreDeploymentType string

// Enum values for LustreDeploymentType
const (
	LustreDeploymentTypeScratch_1    LustreDeploymentType = "SCRATCH_1"
	LustreDeploymentTypeScratch_2    LustreDeploymentType = "SCRATCH_2"
	LustreDeploymentTypePersistent_1 LustreDeploymentType = "PERSISTENT_1"
)

type ReportFormat string

// Enum values for ReportFormat
const (
	ReportFormatReport_csv_20191124 ReportFormat = "REPORT_CSV_20191124"
)

type ReportScope string

// Enum values for ReportScope
const (
	ReportScopeFailed_files_only ReportScope = "FAILED_FILES_ONLY"
)

type ServiceLimit string

// Enum values for ServiceLimit
const (
	ServiceLimitFile_system_count            ServiceLimit = "FILE_SYSTEM_COUNT"
	ServiceLimitTotal_throughput_capacity    ServiceLimit = "TOTAL_THROUGHPUT_CAPACITY"
	ServiceLimitTotal_storage                ServiceLimit = "TOTAL_STORAGE"
	ServiceLimitTotal_user_initiated_backups ServiceLimit = "TOTAL_USER_INITIATED_BACKUPS"
)

type Status string

// Enum values for Status
const (
	StatusFailed             Status = "FAILED"
	StatusIn_progress        Status = "IN_PROGRESS"
	StatusPending            Status = "PENDING"
	StatusCompleted          Status = "COMPLETED"
	StatusUpdated_optimizing Status = "UPDATED_OPTIMIZING"
)

type StorageType string

// Enum values for StorageType
const (
	StorageTypeSsd StorageType = "SSD"
	StorageTypeHdd StorageType = "HDD"
)

type WindowsDeploymentType string

// Enum values for WindowsDeploymentType
const (
	WindowsDeploymentTypeMulti_az_1  WindowsDeploymentType = "MULTI_AZ_1"
	WindowsDeploymentTypeSingle_az_1 WindowsDeploymentType = "SINGLE_AZ_1"
	WindowsDeploymentTypeSingle_az_2 WindowsDeploymentType = "SINGLE_AZ_2"
)
