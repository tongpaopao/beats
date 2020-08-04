// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeletePolicyRequest
type DeletePolicyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the IAM policy you want to delete.
	//
	// For more information about ARNs, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// in the AWS General Reference.
	//
	// PolicyArn is a required field
	PolicyArn *string `min:"20" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePolicyInput"}

	if s.PolicyArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("PolicyArn"))
	}
	if s.PolicyArn != nil && len(*s.PolicyArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("PolicyArn", 20))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeletePolicyOutput
type DeletePolicyOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePolicyOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeletePolicy = "DeletePolicy"

// DeletePolicyRequest returns a request value for making API operation for
// AWS Identity and Access Management.
//
// Deletes the specified managed policy.
//
// Before you can delete a managed policy, you must first detach the policy
// from all users, groups, and roles that it is attached to. In addition, you
// must delete all the policy's versions. The following steps describe the process
// for deleting a managed policy:
//
//    * Detach the policy from all users, groups, and roles that the policy
//    is attached to, using the DetachUserPolicy, DetachGroupPolicy, or DetachRolePolicy
//    API operations. To list all the users, groups, and roles that a policy
//    is attached to, use ListEntitiesForPolicy.
//
//    * Delete all versions of the policy using DeletePolicyVersion. To list
//    the policy's versions, use ListPolicyVersions. You cannot use DeletePolicyVersion
//    to delete the version that is marked as the default version. You delete
//    the policy's default version in the next step of the process.
//
//    * Delete the policy (this automatically deletes the policy's default version)
//    using this API.
//
// For information about managed policies, see Managed Policies and Inline Policies
// (https://docs.aws.amazon.com/IAM/latest/UserGuide/policies-managed-vs-inline.html)
// in the IAM User Guide.
//
//    // Example sending a request using DeletePolicyRequest.
//    req := client.DeletePolicyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iam-2010-05-08/DeletePolicy
func (c *Client) DeletePolicyRequest(input *DeletePolicyInput) DeletePolicyRequest {
	op := &aws.Operation{
		Name:       opDeletePolicy,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeletePolicyInput{}
	}

	req := c.newRequest(op, input, &DeletePolicyOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeletePolicyRequest{Request: req, Input: input, Copy: c.DeletePolicyRequest}
}

// DeletePolicyRequest is the request type for the
// DeletePolicy API operation.
type DeletePolicyRequest struct {
	*aws.Request
	Input *DeletePolicyInput
	Copy  func(*DeletePolicyInput) DeletePolicyRequest
}

// Send marshals and sends the DeletePolicy API request.
func (r DeletePolicyRequest) Send(ctx context.Context) (*DeletePolicyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePolicyResponse{
		DeletePolicyOutput: r.Request.Data.(*DeletePolicyOutput),
		response:           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePolicyResponse is the response type for the
// DeletePolicy API operation.
type DeletePolicyResponse struct {
	*DeletePolicyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePolicy request.
func (r *DeletePolicyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}