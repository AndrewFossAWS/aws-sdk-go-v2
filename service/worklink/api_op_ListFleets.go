// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package worklink

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/ListFleetsRequest
type ListFleetsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to be included in the next page.
	MaxResults *int64 `min:"1" type:"integer"`

	// The pagination token used to retrieve the next page of results for this operation.
	// If this value is null, it retrieves the first page.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListFleetsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListFleetsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListFleetsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if s.NextToken != nil && len(*s.NextToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("NextToken", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFleetsInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.MaxResults != nil {
		v := *s.MaxResults

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MaxResults", protocol.Int64Value(v), metadata)
	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/ListFleetsResponse
type ListFleetsOutput struct {
	_ struct{} `type:"structure"`

	// The summary list of the fleets.
	FleetSummaryList []FleetSummary `type:"list"`

	// The pagination token used to retrieve the next page of results for this operation.
	// If there are no more pages, this value is null.
	NextToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s ListFleetsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListFleetsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.FleetSummaryList) > 0 {
		v := s.FleetSummaryList

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "FleetSummaryList", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextToken != nil {
		v := *s.NextToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListFleets = "ListFleets"

// ListFleetsRequest returns a request value for making API operation for
// Amazon WorkLink.
//
// Retrieves a list of fleets for the current account and Region.
//
//    // Example sending a request using ListFleetsRequest.
//    req := client.ListFleetsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/worklink-2018-09-25/ListFleets
func (c *Client) ListFleetsRequest(input *ListFleetsInput) ListFleetsRequest {
	op := &aws.Operation{
		Name:       opListFleets,
		HTTPMethod: "POST",
		HTTPPath:   "/listFleets",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListFleetsInput{}
	}

	req := c.newRequest(op, input, &ListFleetsOutput{})
	return ListFleetsRequest{Request: req, Input: input, Copy: c.ListFleetsRequest}
}

// ListFleetsRequest is the request type for the
// ListFleets API operation.
type ListFleetsRequest struct {
	*aws.Request
	Input *ListFleetsInput
	Copy  func(*ListFleetsInput) ListFleetsRequest
}

// Send marshals and sends the ListFleets API request.
func (r ListFleetsRequest) Send(ctx context.Context) (*ListFleetsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListFleetsResponse{
		ListFleetsOutput: r.Request.Data.(*ListFleetsOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListFleetsRequestPaginator returns a paginator for ListFleets.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListFleetsRequest(input)
//   p := worklink.NewListFleetsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListFleetsPaginator(req ListFleetsRequest) ListFleetsPaginator {
	return ListFleetsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListFleetsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListFleetsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListFleetsPaginator struct {
	aws.Pager
}

func (p *ListFleetsPaginator) CurrentPage() *ListFleetsOutput {
	return p.Pager.CurrentPage().(*ListFleetsOutput)
}

// ListFleetsResponse is the response type for the
// ListFleets API operation.
type ListFleetsResponse struct {
	*ListFleetsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListFleets request.
func (r *ListFleetsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}