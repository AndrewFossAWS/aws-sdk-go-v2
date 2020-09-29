// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an EFS access point. An access point is an application-specific view
// into an EFS file system that applies an operating system user and group, and a
// file system path, to any file system request made through the access point. The
// operating system user and group override any identity information provided by
// the NFS client. The file system path is exposed as the access point's root
// directory. Applications using the access point can only access data in its own
// directory and below. To learn more, see Mounting a File System Using EFS Access
// Points (https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html). This
// operation requires permissions for the elasticfilesystem:CreateAccessPoint
// action.
func (c *Client) CreateAccessPoint(ctx context.Context, params *CreateAccessPointInput, optFns ...func(*Options)) (*CreateAccessPointOutput, error) {
	stack := middleware.NewStack("CreateAccessPoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateAccessPointMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addIdempotencyToken_opCreateAccessPointMiddleware(stack, options)
	addOpCreateAccessPointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccessPoint(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "CreateAccessPoint",
			Err:           err,
		}
	}
	out := result.(*CreateAccessPointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessPointInput struct {
	// Specifies the directory on the Amazon EFS file system that the access point
	// exposes as the root directory of your file system to NFS clients using the
	// access point. The clients using the access point can only access the root
	// directory and below. If the RootDirectory > Path specified does not exist, EFS
	// creates it and applies the CreationInfo settings when a client connects to an
	// access point. When specifying a RootDirectory, you need to provide the Path, and
	// the CreationInfo is optional.
	RootDirectory *types.RootDirectory
	// The ID of the EFS file system that the access point provides access to.
	FileSystemId *string
	// A string of up to 64 ASCII characters that Amazon EFS uses to ensure idempotent
	// creation.
	ClientToken *string
	// The operating system user and group applied to all file system requests made
	// using the access point.
	PosixUser *types.PosixUser
	// Creates tags associated with the access point. Each tag is a key-value pair.
	Tags []*types.Tag
}

// Provides a description of an EFS file system access point.
type CreateAccessPointOutput struct {
	// The unique Amazon Resource Name (ARN) associated with the access point.
	AccessPointArn *string
	// Identified the AWS account that owns the access point resource.
	OwnerId *string
	// The ID of the access point, assigned by Amazon EFS.
	AccessPointId *string
	// Identifies the lifecycle phase of the access point.
	LifeCycleState types.LifeCycleState
	// The name of the access point. This is the value of the Name tag.
	Name *string
	// The ID of the EFS file system that the access point applies to.
	FileSystemId *string
	// The tags associated with the access point, presented as an array of Tag objects.
	Tags []*types.Tag
	// The directory on the Amazon EFS file system that the access point exposes as the
	// root directory to NFS clients using the access point.
	RootDirectory *types.RootDirectory
	// The opaque string specified in the request to ensure idempotent creation.
	ClientToken *string
	// The full POSIX identity, including the user ID, group ID, and secondary group
	// IDs on the access point that is used for all file operations by NFS clients
	// using the access point.
	PosixUser *types.PosixUser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateAccessPointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateAccessPoint{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAccessPoint{}, middleware.After)
}

type idempotencyToken_initializeOpCreateAccessPoint struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAccessPoint) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAccessPoint) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAccessPointInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAccessPointInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateAccessPointMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateAccessPoint{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAccessPoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "CreateAccessPoint",
	}
}