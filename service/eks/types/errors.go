// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

// This exception is thrown if the request contains a semantic error. The precise
// meaning will depend on the API, and will be documented in the error message.
type BadRequestException struct {
	Message *string
}

func (e *BadRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *BadRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *BadRequestException) ErrorCode() string             { return "BadRequestException" }
func (e *BadRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *BadRequestException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *BadRequestException) HasMessage() bool {
	return e.Message != nil
}

// These errors are usually caused by a client action. Actions can include using an
// action or resource on behalf of a user that doesn't have permissions to use the
// action or resource or specifying an identifier that is not valid.
type ClientException struct {
	Message *string

	ClusterName   *string
	NodegroupName *string
}

func (e *ClientException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClientException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClientException) ErrorCode() string             { return "ClientException" }
func (e *ClientException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ClientException) GetClusterName() string {
	return ptr.ToString(e.ClusterName)
}
func (e *ClientException) HasClusterName() bool {
	return e.ClusterName != nil
}
func (e *ClientException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ClientException) HasMessage() bool {
	return e.Message != nil
}
func (e *ClientException) GetNodegroupName() string {
	return ptr.ToString(e.NodegroupName)
}
func (e *ClientException) HasNodegroupName() bool {
	return e.NodegroupName != nil
}

// The specified parameter is invalid. Review the available parameters for the API
// request.
type InvalidParameterException struct {
	Message *string

	NodegroupName      *string
	FargateProfileName *string
	ClusterName        *string
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidParameterException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidParameterException) ErrorCode() string             { return "InvalidParameterException" }
func (e *InvalidParameterException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidParameterException) GetNodegroupName() string {
	return ptr.ToString(e.NodegroupName)
}
func (e *InvalidParameterException) HasNodegroupName() bool {
	return e.NodegroupName != nil
}
func (e *InvalidParameterException) GetFargateProfileName() string {
	return ptr.ToString(e.FargateProfileName)
}
func (e *InvalidParameterException) HasFargateProfileName() bool {
	return e.FargateProfileName != nil
}
func (e *InvalidParameterException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidParameterException) HasMessage() bool {
	return e.Message != nil
}
func (e *InvalidParameterException) GetClusterName() string {
	return ptr.ToString(e.ClusterName)
}
func (e *InvalidParameterException) HasClusterName() bool {
	return e.ClusterName != nil
}

// The request is invalid given the state of the cluster. Check the state of the
// cluster and the associated operations.
type InvalidRequestException struct {
	Message *string

	ClusterName   *string
	NodegroupName *string
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidRequestException) GetClusterName() string {
	return ptr.ToString(e.ClusterName)
}
func (e *InvalidRequestException) HasClusterName() bool {
	return e.ClusterName != nil
}
func (e *InvalidRequestException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidRequestException) HasMessage() bool {
	return e.Message != nil
}
func (e *InvalidRequestException) GetNodegroupName() string {
	return ptr.ToString(e.NodegroupName)
}
func (e *InvalidRequestException) HasNodegroupName() bool {
	return e.NodegroupName != nil
}

// A service resource associated with the request could not be found. Clients
// should not retry such requests.
type NotFoundException struct {
	Message *string
}

func (e *NotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotFoundException) ErrorCode() string             { return "NotFoundException" }
func (e *NotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *NotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *NotFoundException) HasMessage() bool {
	return e.Message != nil
}

// The specified resource is in use.
type ResourceInUseException struct {
	Message *string

	NodegroupName *string
	ClusterName   *string
}

func (e *ResourceInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceInUseException) ErrorCode() string             { return "ResourceInUseException" }
func (e *ResourceInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceInUseException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceInUseException) HasMessage() bool {
	return e.Message != nil
}
func (e *ResourceInUseException) GetNodegroupName() string {
	return ptr.ToString(e.NodegroupName)
}
func (e *ResourceInUseException) HasNodegroupName() bool {
	return e.NodegroupName != nil
}
func (e *ResourceInUseException) GetClusterName() string {
	return ptr.ToString(e.ClusterName)
}
func (e *ResourceInUseException) HasClusterName() bool {
	return e.ClusterName != nil
}

// You have encountered a service limit on the specified resource.
type ResourceLimitExceededException struct {
	Message *string

	NodegroupName *string
	ClusterName   *string
}

func (e *ResourceLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceLimitExceededException) ErrorCode() string             { return "ResourceLimitExceededException" }
func (e *ResourceLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceLimitExceededException) GetNodegroupName() string {
	return ptr.ToString(e.NodegroupName)
}
func (e *ResourceLimitExceededException) HasNodegroupName() bool {
	return e.NodegroupName != nil
}
func (e *ResourceLimitExceededException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceLimitExceededException) HasMessage() bool {
	return e.Message != nil
}
func (e *ResourceLimitExceededException) GetClusterName() string {
	return ptr.ToString(e.ClusterName)
}
func (e *ResourceLimitExceededException) HasClusterName() bool {
	return e.ClusterName != nil
}

// The specified resource could not be found. You can view your available clusters
// with ListClusters (). You can view your available managed node groups with
// ListNodegroups (). Amazon EKS clusters and node groups are Region-specific.
type ResourceNotFoundException struct {
	Message *string

	FargateProfileName *string
	NodegroupName      *string
	ClusterName        *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFoundException) HasMessage() bool {
	return e.Message != nil
}
func (e *ResourceNotFoundException) GetFargateProfileName() string {
	return ptr.ToString(e.FargateProfileName)
}
func (e *ResourceNotFoundException) HasFargateProfileName() bool {
	return e.FargateProfileName != nil
}
func (e *ResourceNotFoundException) GetNodegroupName() string {
	return ptr.ToString(e.NodegroupName)
}
func (e *ResourceNotFoundException) HasNodegroupName() bool {
	return e.NodegroupName != nil
}
func (e *ResourceNotFoundException) GetClusterName() string {
	return ptr.ToString(e.ClusterName)
}
func (e *ResourceNotFoundException) HasClusterName() bool {
	return e.ClusterName != nil
}

// These errors are usually caused by a server-side issue.
type ServerException struct {
	Message *string

	ClusterName   *string
	NodegroupName *string
}

func (e *ServerException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServerException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServerException) ErrorCode() string             { return "ServerException" }
func (e *ServerException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *ServerException) GetClusterName() string {
	return ptr.ToString(e.ClusterName)
}
func (e *ServerException) HasClusterName() bool {
	return e.ClusterName != nil
}
func (e *ServerException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ServerException) HasMessage() bool {
	return e.Message != nil
}
func (e *ServerException) GetNodegroupName() string {
	return ptr.ToString(e.NodegroupName)
}
func (e *ServerException) HasNodegroupName() bool {
	return e.NodegroupName != nil
}

// The service is unavailable. Back off and retry the operation.
type ServiceUnavailableException struct {
	Message *string
}

func (e *ServiceUnavailableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ServiceUnavailableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ServiceUnavailableException) ErrorCode() string             { return "ServiceUnavailableException" }
func (e *ServiceUnavailableException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *ServiceUnavailableException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ServiceUnavailableException) HasMessage() bool {
	return e.Message != nil
}

// At least one of your specified cluster subnets is in an Availability Zone that
// does not support Amazon EKS. The exception output specifies the supported
// Availability Zones for your account, from which you can choose subnets for your
// cluster.
type UnsupportedAvailabilityZoneException struct {
	Message *string

	ClusterName   *string
	NodegroupName *string
	ValidZones    []*string
}

func (e *UnsupportedAvailabilityZoneException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *UnsupportedAvailabilityZoneException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *UnsupportedAvailabilityZoneException) ErrorCode() string {
	return "UnsupportedAvailabilityZoneException"
}
func (e *UnsupportedAvailabilityZoneException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
func (e *UnsupportedAvailabilityZoneException) GetClusterName() string {
	return ptr.ToString(e.ClusterName)
}
func (e *UnsupportedAvailabilityZoneException) HasClusterName() bool {
	return e.ClusterName != nil
}
func (e *UnsupportedAvailabilityZoneException) GetNodegroupName() string {
	return ptr.ToString(e.NodegroupName)
}
func (e *UnsupportedAvailabilityZoneException) HasNodegroupName() bool {
	return e.NodegroupName != nil
}
func (e *UnsupportedAvailabilityZoneException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *UnsupportedAvailabilityZoneException) HasMessage() bool {
	return e.Message != nil
}
func (e *UnsupportedAvailabilityZoneException) GetValidZones() []*string {
	return e.ValidZones
}
func (e *UnsupportedAvailabilityZoneException) HasValidZones() bool {
	return e.ValidZones != nil
}
