// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateSystemInstanceInput struct {
	_ struct{} `type:"structure"`

	// A document that defines an entity.
	//
	// Definition is a required field
	Definition *DefinitionDocument `locationName:"definition" type:"structure" required:"true"`

	// The ARN of the IAM role that AWS IoT Things Graph will assume when it executes
	// the flow. This role must have read and write access to AWS Lambda and AWS
	// IoT and any other AWS services that the flow uses when it executes. This
	// value is required if the value of the target parameter is CLOUD.
	FlowActionsRoleArn *string `locationName:"flowActionsRoleArn" min:"20" type:"string"`

	// The name of the Greengrass group where the system instance will be deployed.
	// This value is required if the value of the target parameter is GREENGRASS.
	GreengrassGroupName *string `locationName:"greengrassGroupName" type:"string"`

	// An object that specifies whether cloud metrics are collected in a deployment
	// and, if so, what role is used to collect metrics.
	MetricsConfiguration *MetricsConfiguration `locationName:"metricsConfiguration" type:"structure"`

	// The name of the Amazon Simple Storage Service bucket that will be used to
	// store and deploy the system instance's resource file. This value is required
	// if the value of the target parameter is GREENGRASS.
	S3BucketName *string `locationName:"s3BucketName" type:"string"`

	// Metadata, consisting of key-value pairs, that can be used to categorize your
	// system instances.
	Tags []Tag `locationName:"tags" type:"list"`

	// The target type of the deployment. Valid values are GREENGRASS and CLOUD.
	//
	// Target is a required field
	Target DeploymentTarget `locationName:"target" type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s CreateSystemInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateSystemInstanceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateSystemInstanceInput"}

	if s.Definition == nil {
		invalidParams.Add(aws.NewErrParamRequired("Definition"))
	}
	if s.FlowActionsRoleArn != nil && len(*s.FlowActionsRoleArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("FlowActionsRoleArn", 20))
	}
	if len(s.Target) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Target"))
	}
	if s.Definition != nil {
		if err := s.Definition.Validate(); err != nil {
			invalidParams.AddNested("Definition", err.(aws.ErrInvalidParams))
		}
	}
	if s.MetricsConfiguration != nil {
		if err := s.MetricsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("MetricsConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateSystemInstanceOutput struct {
	_ struct{} `type:"structure"`

	// The summary object that describes the new system instance.
	Summary *SystemInstanceSummary `locationName:"summary" type:"structure"`
}

// String returns the string representation
func (s CreateSystemInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateSystemInstance = "CreateSystemInstance"

// CreateSystemInstanceRequest returns a request value for making API operation for
// AWS IoT Things Graph.
//
// Creates a system instance.
//
// This action validates the system instance, prepares the deployment-related
// resources. For Greengrass deployments, it updates the Greengrass group that
// is specified by the greengrassGroupName parameter. It also adds a file to
// the S3 bucket specified by the s3BucketName parameter. You need to call DeploySystemInstance
// after running this action.
//
// For Greengrass deployments, since this action modifies and adds resources
// to a Greengrass group and an S3 bucket on the caller's behalf, the calling
// identity must have write permissions to both the specified Greengrass group
// and S3 bucket. Otherwise, the call will fail with an authorization error.
//
// For cloud deployments, this action requires a flowActionsRoleArn value. This
// is an IAM role that has permissions to access AWS services, such as AWS Lambda
// and AWS IoT, that the flow uses when it executes.
//
// If the definition document doesn't specify a version of the user's namespace,
// the latest version will be used by default.
//
//    // Example sending a request using CreateSystemInstanceRequest.
//    req := client.CreateSystemInstanceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iotthingsgraph-2018-09-06/CreateSystemInstance
func (c *Client) CreateSystemInstanceRequest(input *CreateSystemInstanceInput) CreateSystemInstanceRequest {
	op := &aws.Operation{
		Name:       opCreateSystemInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateSystemInstanceInput{}
	}

	req := c.newRequest(op, input, &CreateSystemInstanceOutput{})
	return CreateSystemInstanceRequest{Request: req, Input: input, Copy: c.CreateSystemInstanceRequest}
}

// CreateSystemInstanceRequest is the request type for the
// CreateSystemInstance API operation.
type CreateSystemInstanceRequest struct {
	*aws.Request
	Input *CreateSystemInstanceInput
	Copy  func(*CreateSystemInstanceInput) CreateSystemInstanceRequest
}

// Send marshals and sends the CreateSystemInstance API request.
func (r CreateSystemInstanceRequest) Send(ctx context.Context) (*CreateSystemInstanceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateSystemInstanceResponse{
		CreateSystemInstanceOutput: r.Request.Data.(*CreateSystemInstanceOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateSystemInstanceResponse is the response type for the
// CreateSystemInstance API operation.
type CreateSystemInstanceResponse struct {
	*CreateSystemInstanceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateSystemInstance request.
func (r *CreateSystemInstanceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}