// Code generated by smithy-go-codegen DO NOT EDIT.

package lexruntimeservice

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/lexruntimeservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpDeleteSession struct {
}

func (*validateOpDeleteSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpDeleteSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*DeleteSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpDeleteSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpGetSession struct {
}

func (*validateOpGetSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpGetSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*GetSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpGetSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPostContent struct {
}

func (*validateOpPostContent) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPostContent) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PostContentInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPostContentInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPostText struct {
}

func (*validateOpPostText) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPostText) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PostTextInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPostTextInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

type validateOpPutSession struct {
}

func (*validateOpPutSession) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutSession) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutSessionInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutSessionInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpDeleteSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpDeleteSession{}, middleware.After)
}

func addOpGetSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpGetSession{}, middleware.After)
}

func addOpPostContentValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPostContent{}, middleware.After)
}

func addOpPostTextValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPostText{}, middleware.After)
}

func addOpPutSessionValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutSession{}, middleware.After)
}

func validateDialogAction(v *types.DialogAction) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DialogAction"}
	if len(v.Type) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateIntentSummary(v *types.IntentSummary) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IntentSummary"}
	if len(v.DialogActionType) == 0 {
		invalidParams.Add(smithy.NewErrParamRequired("DialogActionType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateIntentSummaryList(v []*types.IntentSummary) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "IntentSummaryList"}
	for i := range v {
		if err := validateIntentSummary(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpDeleteSessionInput(v *DeleteSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "DeleteSessionInput"}
	if v.UserId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserId"))
	}
	if v.BotName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotName"))
	}
	if v.BotAlias == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotAlias"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpGetSessionInput(v *GetSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "GetSessionInput"}
	if v.BotName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotName"))
	}
	if v.BotAlias == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotAlias"))
	}
	if v.UserId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPostContentInput(v *PostContentInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PostContentInput"}
	if v.InputStream == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputStream"))
	}
	if v.ContentType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("ContentType"))
	}
	if v.UserId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserId"))
	}
	if v.BotName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotName"))
	}
	if v.BotAlias == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotAlias"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPostTextInput(v *PostTextInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PostTextInput"}
	if v.UserId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserId"))
	}
	if v.InputText == nil {
		invalidParams.Add(smithy.NewErrParamRequired("InputText"))
	}
	if v.BotName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotName"))
	}
	if v.BotAlias == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotAlias"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutSessionInput(v *PutSessionInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutSessionInput"}
	if v.UserId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("UserId"))
	}
	if v.BotName == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotName"))
	}
	if v.BotAlias == nil {
		invalidParams.Add(smithy.NewErrParamRequired("BotAlias"))
	}
	if v.DialogAction != nil {
		if err := validateDialogAction(v.DialogAction); err != nil {
			invalidParams.AddNested("DialogAction", err.(smithy.InvalidParamsError))
		}
	}
	if v.RecentIntentSummaryView != nil {
		if err := validateIntentSummaryList(v.RecentIntentSummaryView); err != nil {
			invalidParams.AddNested("RecentIntentSummaryView", err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
