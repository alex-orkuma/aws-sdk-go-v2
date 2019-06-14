// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workspaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/StartWorkspacesRequest
type StartWorkspacesInput struct {
	_ struct{} `type:"structure"`

	// The WorkSpaces to start. You can specify up to 25 WorkSpaces.
	//
	// StartWorkspaceRequests is a required field
	StartWorkspaceRequests []StartRequest `min:"1" type:"list" required:"true"`
}

// String returns the string representation
func (s StartWorkspacesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartWorkspacesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartWorkspacesInput"}

	if s.StartWorkspaceRequests == nil {
		invalidParams.Add(aws.NewErrParamRequired("StartWorkspaceRequests"))
	}
	if s.StartWorkspaceRequests != nil && len(s.StartWorkspaceRequests) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("StartWorkspaceRequests", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/StartWorkspacesResult
type StartWorkspacesOutput struct {
	_ struct{} `type:"structure"`

	// Information about the WorkSpaces that could not be started.
	FailedRequests []FailedWorkspaceChangeRequest `type:"list"`
}

// String returns the string representation
func (s StartWorkspacesOutput) String() string {
	return awsutil.Prettify(s)
}

const opStartWorkspaces = "StartWorkspaces"

// StartWorkspacesRequest returns a request value for making API operation for
// Amazon WorkSpaces.
//
// Starts the specified WorkSpaces.
//
// You cannot start a WorkSpace unless it has a running mode of AutoStop and
// a state of STOPPED.
//
//    // Example sending a request using StartWorkspacesRequest.
//    req := client.StartWorkspacesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workspaces-2015-04-08/StartWorkspaces
func (c *Client) StartWorkspacesRequest(input *StartWorkspacesInput) StartWorkspacesRequest {
	op := &aws.Operation{
		Name:       opStartWorkspaces,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartWorkspacesInput{}
	}

	req := c.newRequest(op, input, &StartWorkspacesOutput{})
	return StartWorkspacesRequest{Request: req, Input: input, Copy: c.StartWorkspacesRequest}
}

// StartWorkspacesRequest is the request type for the
// StartWorkspaces API operation.
type StartWorkspacesRequest struct {
	*aws.Request
	Input *StartWorkspacesInput
	Copy  func(*StartWorkspacesInput) StartWorkspacesRequest
}

// Send marshals and sends the StartWorkspaces API request.
func (r StartWorkspacesRequest) Send(ctx context.Context) (*StartWorkspacesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &StartWorkspacesResponse{
		StartWorkspacesOutput: r.Request.Data.(*StartWorkspacesOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// StartWorkspacesResponse is the response type for the
// StartWorkspaces API operation.
type StartWorkspacesResponse struct {
	*StartWorkspacesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// StartWorkspaces request.
func (r *StartWorkspacesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}