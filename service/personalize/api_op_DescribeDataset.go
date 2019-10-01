// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeDatasetInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the dataset to describe.
	//
	// DatasetArn is a required field
	DatasetArn *string `locationName:"datasetArn" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDatasetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDatasetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeDatasetInput"}

	if s.DatasetArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DatasetArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeDatasetOutput struct {
	_ struct{} `type:"structure"`

	// A listing of the dataset's properties.
	Dataset *Dataset `locationName:"dataset" type:"structure"`
}

// String returns the string representation
func (s DescribeDatasetOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeDataset = "DescribeDataset"

// DescribeDatasetRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Describes the given dataset. For more information on datasets, see CreateDataset.
//
//    // Example sending a request using DescribeDatasetRequest.
//    req := client.DescribeDatasetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/DescribeDataset
func (c *Client) DescribeDatasetRequest(input *DescribeDatasetInput) DescribeDatasetRequest {
	op := &aws.Operation{
		Name:       opDescribeDataset,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDatasetInput{}
	}

	req := c.newRequest(op, input, &DescribeDatasetOutput{})
	return DescribeDatasetRequest{Request: req, Input: input, Copy: c.DescribeDatasetRequest}
}

// DescribeDatasetRequest is the request type for the
// DescribeDataset API operation.
type DescribeDatasetRequest struct {
	*aws.Request
	Input *DescribeDatasetInput
	Copy  func(*DescribeDatasetInput) DescribeDatasetRequest
}

// Send marshals and sends the DescribeDataset API request.
func (r DescribeDatasetRequest) Send(ctx context.Context) (*DescribeDatasetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDatasetResponse{
		DescribeDatasetOutput: r.Request.Data.(*DescribeDatasetOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDatasetResponse is the response type for the
// DescribeDataset API operation.
type DescribeDatasetResponse struct {
	*DescribeDatasetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDataset request.
func (r *DescribeDatasetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}