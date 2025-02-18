// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package resourcemanager

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	resourcemanagerpb "cloud.google.com/go/resourcemanager/apiv3/resourcemanagerpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newTagValuesClientHook clientHook

// TagValuesCallOptions contains the retry settings for each method of TagValuesClient.
type TagValuesCallOptions struct {
	ListTagValues      []gax.CallOption
	GetTagValue        []gax.CallOption
	CreateTagValue     []gax.CallOption
	UpdateTagValue     []gax.CallOption
	DeleteTagValue     []gax.CallOption
	GetIamPolicy       []gax.CallOption
	SetIamPolicy       []gax.CallOption
	TestIamPermissions []gax.CallOption
}

func defaultTagValuesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("cloudresourcemanager.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("cloudresourcemanager.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://cloudresourcemanager.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultTagValuesCallOptions() *TagValuesCallOptions {
	return &TagValuesCallOptions{
		ListTagValues: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetTagValue: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreateTagValue: []gax.CallOption{},
		UpdateTagValue: []gax.CallOption{},
		DeleteTagValue: []gax.CallOption{},
		GetIamPolicy: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SetIamPolicy:       []gax.CallOption{},
		TestIamPermissions: []gax.CallOption{},
	}
}

// internalTagValuesClient is an interface that defines the methods available from Cloud Resource Manager API.
type internalTagValuesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListTagValues(context.Context, *resourcemanagerpb.ListTagValuesRequest, ...gax.CallOption) *TagValueIterator
	GetTagValue(context.Context, *resourcemanagerpb.GetTagValueRequest, ...gax.CallOption) (*resourcemanagerpb.TagValue, error)
	CreateTagValue(context.Context, *resourcemanagerpb.CreateTagValueRequest, ...gax.CallOption) (*CreateTagValueOperation, error)
	CreateTagValueOperation(name string) *CreateTagValueOperation
	UpdateTagValue(context.Context, *resourcemanagerpb.UpdateTagValueRequest, ...gax.CallOption) (*UpdateTagValueOperation, error)
	UpdateTagValueOperation(name string) *UpdateTagValueOperation
	DeleteTagValue(context.Context, *resourcemanagerpb.DeleteTagValueRequest, ...gax.CallOption) (*DeleteTagValueOperation, error)
	DeleteTagValueOperation(name string) *DeleteTagValueOperation
	GetIamPolicy(context.Context, *iampb.GetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	SetIamPolicy(context.Context, *iampb.SetIamPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	TestIamPermissions(context.Context, *iampb.TestIamPermissionsRequest, ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error)
}

// TagValuesClient is a client for interacting with Cloud Resource Manager API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Allow users to create and manage tag values.
type TagValuesClient struct {
	// The internal transport-dependent client.
	internalClient internalTagValuesClient

	// The call options for this service.
	CallOptions *TagValuesCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TagValuesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TagValuesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TagValuesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListTagValues lists all TagValues for a specific TagKey.
func (c *TagValuesClient) ListTagValues(ctx context.Context, req *resourcemanagerpb.ListTagValuesRequest, opts ...gax.CallOption) *TagValueIterator {
	return c.internalClient.ListTagValues(ctx, req, opts...)
}

// GetTagValue retrieves TagValue. If the TagValue or namespaced name does not exist, or
// if the user does not have permission to view it, this method will return
// PERMISSION_DENIED.
func (c *TagValuesClient) GetTagValue(ctx context.Context, req *resourcemanagerpb.GetTagValueRequest, opts ...gax.CallOption) (*resourcemanagerpb.TagValue, error) {
	return c.internalClient.GetTagValue(ctx, req, opts...)
}

// CreateTagValue creates a TagValue as a child of the specified TagKey. If a another
// request with the same parameters is sent while the original request is in
// process the second request will receive an error. A maximum of 300
// TagValues can exist under a TagKey at any given time.
func (c *TagValuesClient) CreateTagValue(ctx context.Context, req *resourcemanagerpb.CreateTagValueRequest, opts ...gax.CallOption) (*CreateTagValueOperation, error) {
	return c.internalClient.CreateTagValue(ctx, req, opts...)
}

// CreateTagValueOperation returns a new CreateTagValueOperation from a given name.
// The name must be that of a previously created CreateTagValueOperation, possibly from a different process.
func (c *TagValuesClient) CreateTagValueOperation(name string) *CreateTagValueOperation {
	return c.internalClient.CreateTagValueOperation(name)
}

// UpdateTagValue updates the attributes of the TagValue resource.
func (c *TagValuesClient) UpdateTagValue(ctx context.Context, req *resourcemanagerpb.UpdateTagValueRequest, opts ...gax.CallOption) (*UpdateTagValueOperation, error) {
	return c.internalClient.UpdateTagValue(ctx, req, opts...)
}

// UpdateTagValueOperation returns a new UpdateTagValueOperation from a given name.
// The name must be that of a previously created UpdateTagValueOperation, possibly from a different process.
func (c *TagValuesClient) UpdateTagValueOperation(name string) *UpdateTagValueOperation {
	return c.internalClient.UpdateTagValueOperation(name)
}

// DeleteTagValue deletes a TagValue. The TagValue cannot have any bindings when it is
// deleted.
func (c *TagValuesClient) DeleteTagValue(ctx context.Context, req *resourcemanagerpb.DeleteTagValueRequest, opts ...gax.CallOption) (*DeleteTagValueOperation, error) {
	return c.internalClient.DeleteTagValue(ctx, req, opts...)
}

// DeleteTagValueOperation returns a new DeleteTagValueOperation from a given name.
// The name must be that of a previously created DeleteTagValueOperation, possibly from a different process.
func (c *TagValuesClient) DeleteTagValueOperation(name string) *DeleteTagValueOperation {
	return c.internalClient.DeleteTagValueOperation(name)
}

// GetIamPolicy gets the access control policy for a TagValue. The returned policy may be
// empty if no such policy or resource exists. The resource field should
// be the TagValue’s resource name. For example: tagValues/1234.
// The caller must have the
// cloudresourcemanager.googleapis.com/tagValues.getIamPolicy permission on
// the identified TagValue to get the access control policy.
func (c *TagValuesClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetIamPolicy(ctx, req, opts...)
}

// SetIamPolicy sets the access control policy on a TagValue, replacing any existing
// policy. The resource field should be the TagValue’s resource name.
// For example: tagValues/1234.
// The caller must have resourcemanager.tagValues.setIamPolicy permission
// on the identified tagValue.
func (c *TagValuesClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.SetIamPolicy(ctx, req, opts...)
}

// TestIamPermissions returns permissions that a caller has on the specified TagValue.
// The resource field should be the TagValue’s resource name. For example:
// tagValues/1234.
//
// There are no permissions required for making this API call.
func (c *TagValuesClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	return c.internalClient.TestIamPermissions(ctx, req, opts...)
}

// tagValuesGRPCClient is a client for interacting with Cloud Resource Manager API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type tagValuesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing TagValuesClient
	CallOptions **TagValuesCallOptions

	// The gRPC API client.
	tagValuesClient resourcemanagerpb.TagValuesClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewTagValuesClient creates a new tag values client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Allow users to create and manage tag values.
func NewTagValuesClient(ctx context.Context, opts ...option.ClientOption) (*TagValuesClient, error) {
	clientOpts := defaultTagValuesGRPCClientOptions()
	if newTagValuesClientHook != nil {
		hookOpts, err := newTagValuesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TagValuesClient{CallOptions: defaultTagValuesCallOptions()}

	c := &tagValuesGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		tagValuesClient:  resourcemanagerpb.NewTagValuesClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *tagValuesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *tagValuesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *tagValuesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *tagValuesGRPCClient) ListTagValues(ctx context.Context, req *resourcemanagerpb.ListTagValuesRequest, opts ...gax.CallOption) *TagValueIterator {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ListTagValues[0:len((*c.CallOptions).ListTagValues):len((*c.CallOptions).ListTagValues)], opts...)
	it := &TagValueIterator{}
	req = proto.Clone(req).(*resourcemanagerpb.ListTagValuesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*resourcemanagerpb.TagValue, string, error) {
		resp := &resourcemanagerpb.ListTagValuesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.tagValuesClient.ListTagValues(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetTagValues(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *tagValuesGRPCClient) GetTagValue(ctx context.Context, req *resourcemanagerpb.GetTagValueRequest, opts ...gax.CallOption) (*resourcemanagerpb.TagValue, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetTagValue[0:len((*c.CallOptions).GetTagValue):len((*c.CallOptions).GetTagValue)], opts...)
	var resp *resourcemanagerpb.TagValue
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.tagValuesClient.GetTagValue(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *tagValuesGRPCClient) CreateTagValue(ctx context.Context, req *resourcemanagerpb.CreateTagValueRequest, opts ...gax.CallOption) (*CreateTagValueOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).CreateTagValue[0:len((*c.CallOptions).CreateTagValue):len((*c.CallOptions).CreateTagValue)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.tagValuesClient.CreateTagValue(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreateTagValueOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *tagValuesGRPCClient) UpdateTagValue(ctx context.Context, req *resourcemanagerpb.UpdateTagValueRequest, opts ...gax.CallOption) (*UpdateTagValueOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "tag_value.name", url.QueryEscape(req.GetTagValue().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdateTagValue[0:len((*c.CallOptions).UpdateTagValue):len((*c.CallOptions).UpdateTagValue)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.tagValuesClient.UpdateTagValue(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdateTagValueOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *tagValuesGRPCClient) DeleteTagValue(ctx context.Context, req *resourcemanagerpb.DeleteTagValueRequest, opts ...gax.CallOption) (*DeleteTagValueOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeleteTagValue[0:len((*c.CallOptions).DeleteTagValue):len((*c.CallOptions).DeleteTagValue)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.tagValuesClient.DeleteTagValue(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeleteTagValueOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *tagValuesGRPCClient) GetIamPolicy(ctx context.Context, req *iampb.GetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetIamPolicy[0:len((*c.CallOptions).GetIamPolicy):len((*c.CallOptions).GetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.tagValuesClient.GetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *tagValuesGRPCClient) SetIamPolicy(ctx context.Context, req *iampb.SetIamPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SetIamPolicy[0:len((*c.CallOptions).SetIamPolicy):len((*c.CallOptions).SetIamPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.tagValuesClient.SetIamPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *tagValuesGRPCClient) TestIamPermissions(ctx context.Context, req *iampb.TestIamPermissionsRequest, opts ...gax.CallOption) (*iampb.TestIamPermissionsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource", url.QueryEscape(req.GetResource())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).TestIamPermissions[0:len((*c.CallOptions).TestIamPermissions):len((*c.CallOptions).TestIamPermissions)], opts...)
	var resp *iampb.TestIamPermissionsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.tagValuesClient.TestIamPermissions(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateTagValueOperation manages a long-running operation from CreateTagValue.
type CreateTagValueOperation struct {
	lro *longrunning.Operation
}

// CreateTagValueOperation returns a new CreateTagValueOperation from a given name.
// The name must be that of a previously created CreateTagValueOperation, possibly from a different process.
func (c *tagValuesGRPCClient) CreateTagValueOperation(name string) *CreateTagValueOperation {
	return &CreateTagValueOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreateTagValueOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*resourcemanagerpb.TagValue, error) {
	var resp resourcemanagerpb.TagValue
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreateTagValueOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*resourcemanagerpb.TagValue, error) {
	var resp resourcemanagerpb.TagValue
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreateTagValueOperation) Metadata() (*resourcemanagerpb.CreateTagValueMetadata, error) {
	var meta resourcemanagerpb.CreateTagValueMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreateTagValueOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreateTagValueOperation) Name() string {
	return op.lro.Name()
}

// DeleteTagValueOperation manages a long-running operation from DeleteTagValue.
type DeleteTagValueOperation struct {
	lro *longrunning.Operation
}

// DeleteTagValueOperation returns a new DeleteTagValueOperation from a given name.
// The name must be that of a previously created DeleteTagValueOperation, possibly from a different process.
func (c *tagValuesGRPCClient) DeleteTagValueOperation(name string) *DeleteTagValueOperation {
	return &DeleteTagValueOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeleteTagValueOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*resourcemanagerpb.TagValue, error) {
	var resp resourcemanagerpb.TagValue
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeleteTagValueOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*resourcemanagerpb.TagValue, error) {
	var resp resourcemanagerpb.TagValue
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeleteTagValueOperation) Metadata() (*resourcemanagerpb.DeleteTagValueMetadata, error) {
	var meta resourcemanagerpb.DeleteTagValueMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeleteTagValueOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeleteTagValueOperation) Name() string {
	return op.lro.Name()
}

// UpdateTagValueOperation manages a long-running operation from UpdateTagValue.
type UpdateTagValueOperation struct {
	lro *longrunning.Operation
}

// UpdateTagValueOperation returns a new UpdateTagValueOperation from a given name.
// The name must be that of a previously created UpdateTagValueOperation, possibly from a different process.
func (c *tagValuesGRPCClient) UpdateTagValueOperation(name string) *UpdateTagValueOperation {
	return &UpdateTagValueOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdateTagValueOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*resourcemanagerpb.TagValue, error) {
	var resp resourcemanagerpb.TagValue
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UpdateTagValueOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*resourcemanagerpb.TagValue, error) {
	var resp resourcemanagerpb.TagValue
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UpdateTagValueOperation) Metadata() (*resourcemanagerpb.UpdateTagValueMetadata, error) {
	var meta resourcemanagerpb.UpdateTagValueMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdateTagValueOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdateTagValueOperation) Name() string {
	return op.lro.Name()
}

// TagValueIterator manages a stream of *resourcemanagerpb.TagValue.
type TagValueIterator struct {
	items    []*resourcemanagerpb.TagValue
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*resourcemanagerpb.TagValue, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *TagValueIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *TagValueIterator) Next() (*resourcemanagerpb.TagValue, error) {
	var item *resourcemanagerpb.TagValue
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *TagValueIterator) bufLen() int {
	return len(it.items)
}

func (it *TagValueIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
