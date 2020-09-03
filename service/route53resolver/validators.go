// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpAssociateResolverEndpointIpAddress struct {
}

func (*validateOpAssociateResolverEndpointIpAddress) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateResolverEndpointIpAddress) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateResolverEndpointIpAddressInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateResolverEndpointIpAddressInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpAssociateResolverRule struct {
}

func (*validateOpAssociateResolverRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpAssociateResolverRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*AssociateResolverRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpAssociateResolverRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateResolverEndpoint struct {
}

func (*validateOpCreateResolverEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateResolverEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateResolverEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateResolverEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpCreateResolverRule struct {
}

func (*validateOpCreateResolverRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpCreateResolverRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*CreateResolverRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpCreateResolverRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteResolverEndpoint struct {
}

func (*validateOpDeleteResolverEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteResolverEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteResolverEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteResolverEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDeleteResolverRule struct {
}

func (*validateOpDeleteResolverRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteResolverRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteResolverRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteResolverRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateResolverEndpointIpAddress struct {
}

func (*validateOpDisassociateResolverEndpointIpAddress) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateResolverEndpointIpAddress) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateResolverEndpointIpAddressInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateResolverEndpointIpAddressInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpDisassociateResolverRule struct {
}

func (*validateOpDisassociateResolverRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDisassociateResolverRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DisassociateResolverRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDisassociateResolverRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResolverEndpoint struct {
}

func (*validateOpGetResolverEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResolverEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResolverEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResolverEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResolverRuleAssociation struct {
}

func (*validateOpGetResolverRuleAssociation) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResolverRuleAssociation) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResolverRuleAssociationInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResolverRuleAssociationInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResolverRule struct {
}

func (*validateOpGetResolverRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResolverRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResolverRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResolverRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetResolverRulePolicy struct {
}

func (*validateOpGetResolverRulePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetResolverRulePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetResolverRulePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetResolverRulePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListResolverEndpointIpAddresses struct {
}

func (*validateOpListResolverEndpointIpAddresses) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListResolverEndpointIpAddresses) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListResolverEndpointIpAddressesInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListResolverEndpointIpAddressesInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpListTagsForResource struct {
}

func (*validateOpListTagsForResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpListTagsForResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*ListTagsForResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpListTagsForResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutResolverRulePolicy struct {
}

func (*validateOpPutResolverRulePolicy) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutResolverRulePolicy) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutResolverRulePolicyInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutResolverRulePolicyInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpTagResource struct {
}

func (*validateOpTagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpTagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*TagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpTagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUntagResource struct {
}

func (*validateOpUntagResource) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUntagResource) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UntagResourceInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUntagResourceInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateResolverEndpoint struct {
}

func (*validateOpUpdateResolverEndpoint) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateResolverEndpoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateResolverEndpointInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateResolverEndpointInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpUpdateResolverRule struct {
}

func (*validateOpUpdateResolverRule) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpUpdateResolverRule) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*UpdateResolverRuleInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpUpdateResolverRuleInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpAssociateResolverEndpointIpAddressValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateResolverEndpointIpAddress{}, middleware.After)
}

func addOpAssociateResolverRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpAssociateResolverRule{}, middleware.After)
}

func addOpCreateResolverEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateResolverEndpoint{}, middleware.After)
}

func addOpCreateResolverRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpCreateResolverRule{}, middleware.After)
}

func addOpDeleteResolverEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteResolverEndpoint{}, middleware.After)
}

func addOpDeleteResolverRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteResolverRule{}, middleware.After)
}

func addOpDisassociateResolverEndpointIpAddressValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateResolverEndpointIpAddress{}, middleware.After)
}

func addOpDisassociateResolverRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDisassociateResolverRule{}, middleware.After)
}

func addOpGetResolverEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResolverEndpoint{}, middleware.After)
}

func addOpGetResolverRuleAssociationValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResolverRuleAssociation{}, middleware.After)
}

func addOpGetResolverRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResolverRule{}, middleware.After)
}

func addOpGetResolverRulePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetResolverRulePolicy{}, middleware.After)
}

func addOpListResolverEndpointIpAddressesValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListResolverEndpointIpAddresses{}, middleware.After)
}

func addOpListTagsForResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpListTagsForResource{}, middleware.After)
}

func addOpPutResolverRulePolicyValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutResolverRulePolicy{}, middleware.After)
}

func addOpTagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpTagResource{}, middleware.After)
}

func addOpUntagResourceValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUntagResource{}, middleware.After)
}

func addOpUpdateResolverEndpointValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateResolverEndpoint{}, middleware.After)
}

func addOpUpdateResolverRuleValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpUpdateResolverRule{}, middleware.After)
}

func validateIpAddressesRequest(v []*types.IpAddressRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IpAddressesRequest"}
	for i := range v {
		if err := validateIpAddressRequest(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateIpAddressRequest(v *types.IpAddressRequest) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IpAddressRequest"}
	if v.SubnetId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SubnetId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateResolverRuleConfig(v *types.ResolverRuleConfig) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ResolverRuleConfig"}
	if v.TargetIps != nil {
		if err := validateTargetList(v.TargetIps); err != nil {
			invalidParams.AddNested("TargetIps", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetAddress(v *types.TargetAddress) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetAddress"}
	if v.Ip == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Ip"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateTargetList(v []*types.TargetAddress) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TargetList"}
	for i := range v {
		if err := validateTargetAddress(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateResolverEndpointIpAddressInput(v *AssociateResolverEndpointIpAddressInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateResolverEndpointIpAddressInput"}
	if v.IpAddress == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IpAddress"))
	}
	if v.ResolverEndpointId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverEndpointId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpAssociateResolverRuleInput(v *AssociateResolverRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "AssociateResolverRuleInput"}
	if v.ResolverRuleId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverRuleId"))
	}
	if v.VPCId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VPCId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateResolverEndpointInput(v *CreateResolverEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateResolverEndpointInput"}
	if v.SecurityGroupIds == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SecurityGroupIds"))
	}
	if v.IpAddresses == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IpAddresses"))
	} else if v.IpAddresses != nil {
		if err := validateIpAddressesRequest(v.IpAddresses); err != nil {
			invalidParams.AddNested("IpAddresses", err.(smithy.InvalidParamsError))
		}
	}
	if v.CreatorRequestId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CreatorRequestId"))
	}
	if len(v.Direction) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Direction"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpCreateResolverRuleInput(v *CreateResolverRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "CreateResolverRuleInput"}
	if v.CreatorRequestId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("CreatorRequestId"))
	}
	if len(v.RuleType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("RuleType"))
	}
	if v.TargetIps != nil {
		if err := validateTargetList(v.TargetIps); err != nil {
			invalidParams.AddNested("TargetIps", err.(smithy.InvalidParamsError))
		}
	}
	if v.DomainName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("DomainName"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteResolverEndpointInput(v *DeleteResolverEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteResolverEndpointInput"}
	if v.ResolverEndpointId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverEndpointId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteResolverRuleInput(v *DeleteResolverRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteResolverRuleInput"}
	if v.ResolverRuleId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverRuleId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateResolverEndpointIpAddressInput(v *DisassociateResolverEndpointIpAddressInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateResolverEndpointIpAddressInput"}
	if v.ResolverEndpointId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverEndpointId"))
	}
	if v.IpAddress == nil {
		invalidParams.Add(smithy.NewErrParamRequired("IpAddress"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDisassociateResolverRuleInput(v *DisassociateResolverRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DisassociateResolverRuleInput"}
	if v.ResolverRuleId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverRuleId"))
	}
	if v.VPCId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("VPCId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResolverEndpointInput(v *GetResolverEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResolverEndpointInput"}
	if v.ResolverEndpointId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverEndpointId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResolverRuleAssociationInput(v *GetResolverRuleAssociationInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResolverRuleAssociationInput"}
	if v.ResolverRuleAssociationId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverRuleAssociationId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResolverRuleInput(v *GetResolverRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResolverRuleInput"}
	if v.ResolverRuleId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverRuleId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetResolverRulePolicyInput(v *GetResolverRulePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetResolverRulePolicyInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListResolverEndpointIpAddressesInput(v *ListResolverEndpointIpAddressesInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListResolverEndpointIpAddressesInput"}
	if v.ResolverEndpointId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverEndpointId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpListTagsForResourceInput(v *ListTagsForResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "ListTagsForResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutResolverRulePolicyInput(v *PutResolverRulePolicyInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutResolverRulePolicyInput"}
	if v.Arn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Arn"))
	}
	if v.ResolverRulePolicy == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverRulePolicy"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpTagResourceInput(v *TagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "TagResourceInput"}
	if v.Tags == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Tags"))
	}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUntagResourceInput(v *UntagResourceInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UntagResourceInput"}
	if v.ResourceArn == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResourceArn"))
	}
	if v.TagKeys == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TagKeys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateResolverEndpointInput(v *UpdateResolverEndpointInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateResolverEndpointInput"}
	if v.ResolverEndpointId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverEndpointId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpUpdateResolverRuleInput(v *UpdateResolverRuleInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "UpdateResolverRuleInput"}
	if v.Config == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Config"))
	} else if v.Config != nil {
		if err := validateResolverRuleConfig(v.Config); err != nil {
			invalidParams.AddNested("Config", err.(smithy.InvalidParamsError))
		}
	}
	if v.ResolverRuleId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ResolverRuleId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
