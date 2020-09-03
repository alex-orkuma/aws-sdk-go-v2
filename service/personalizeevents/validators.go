// Code generated by smithy-go-codegen DO NOT EDIT.

package personalizeevents

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/personalizeevents/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
)

type validateOpPutEvents struct {
}

func (*validateOpPutEvents) ID() string {
	return "OperationInputValidation"
}

func (m *validateOpPutEvents) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	input, ok := in.Parameters.(*PutEventsInput)
	if !ok {
		return out, metadata, fmt.Errorf("unknown input parameters type %T", in.Parameters)
	}
	if err := validateOpPutEventsInput(input); err != nil {
		return out, metadata, err
	}
	return next.HandleInitialize(ctx, in)
}

func addOpPutEventsValidationMiddleware(stack *middleware.Stack) error {
	return stack.Initialize.Add(&validateOpPutEvents{}, middleware.After)
}

func validateEvent(v *types.Event) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "Event"}
	if v.EventType == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EventType"))
	}
	if v.Properties == nil {
		invalidParams.Add(smithy.NewErrParamRequired("Properties"))
	}
	if v.SentAt == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SentAt"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateEventList(v []*types.Event) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "EventList"}
	for i := range v {
		if err := validateEvent(v[i]); err != nil {
			invalidParams.AddNested(fmt.Sprintf("[%d]", i), err.(smithy.InvalidParamsError))
		}
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}

func validateOpPutEventsInput(v *PutEventsInput) error {
	if v == nil {
		return nil
	}
	invalidParams := smithy.InvalidParamsError{Context: "PutEventsInput"}
	if v.EventList == nil {
		invalidParams.Add(smithy.NewErrParamRequired("EventList"))
	} else if v.EventList != nil {
		if err := validateEventList(v.EventList); err != nil {
			invalidParams.AddNested("EventList", err.(smithy.InvalidParamsError))
		}
	}
	if v.TrackingId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("TrackingId"))
	}
	if v.SessionId == nil {
		invalidParams.Add(smithy.NewErrParamRequired("SessionId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	} else {
		return nil
	}
}
