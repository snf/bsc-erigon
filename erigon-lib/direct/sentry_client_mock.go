// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ledgerwatch/erigon-lib/direct (interfaces: SentryClient)
//
// Generated by this command:
//
//	mockgen -typed=true -destination=./sentry_client_mock.go -package=direct . SentryClient
//

// Package direct is a generated GoMock package.
package direct

import (
	context "context"
	reflect "reflect"

	sentry "github.com/ledgerwatch/erigon-lib/gointerfaces/sentry"
	types "github.com/ledgerwatch/erigon-lib/gointerfaces/types"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockSentryClient is a mock of SentryClient interface.
type MockSentryClient struct {
	ctrl     *gomock.Controller
	recorder *MockSentryClientMockRecorder
}

// MockSentryClientMockRecorder is the mock recorder for MockSentryClient.
type MockSentryClientMockRecorder struct {
	mock *MockSentryClient
}

// NewMockSentryClient creates a new mock instance.
func NewMockSentryClient(ctrl *gomock.Controller) *MockSentryClient {
	mock := &MockSentryClient{ctrl: ctrl}
	mock.recorder = &MockSentryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSentryClient) EXPECT() *MockSentryClientMockRecorder {
	return m.recorder
}

// AddPeer mocks base method.
func (m *MockSentryClient) AddPeer(arg0 context.Context, arg1 *sentry.AddPeerRequest, arg2 ...grpc.CallOption) (*sentry.AddPeerReply, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddPeer", varargs...)
	ret0, _ := ret[0].(*sentry.AddPeerReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPeer indicates an expected call of AddPeer.
func (mr *MockSentryClientMockRecorder) AddPeer(arg0, arg1 any, arg2 ...any) *MockSentryClientAddPeerCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPeer", reflect.TypeOf((*MockSentryClient)(nil).AddPeer), varargs...)
	return &MockSentryClientAddPeerCall{Call: call}
}

// MockSentryClientAddPeerCall wrap *gomock.Call
type MockSentryClientAddPeerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientAddPeerCall) Return(arg0 *sentry.AddPeerReply, arg1 error) *MockSentryClientAddPeerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientAddPeerCall) Do(f func(context.Context, *sentry.AddPeerRequest, ...grpc.CallOption) (*sentry.AddPeerReply, error)) *MockSentryClientAddPeerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientAddPeerCall) DoAndReturn(f func(context.Context, *sentry.AddPeerRequest, ...grpc.CallOption) (*sentry.AddPeerReply, error)) *MockSentryClientAddPeerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// HandShake mocks base method.
func (m *MockSentryClient) HandShake(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*sentry.HandShakeReply, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HandShake", varargs...)
	ret0, _ := ret[0].(*sentry.HandShakeReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandShake indicates an expected call of HandShake.
func (mr *MockSentryClientMockRecorder) HandShake(arg0, arg1 any, arg2 ...any) *MockSentryClientHandShakeCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandShake", reflect.TypeOf((*MockSentryClient)(nil).HandShake), varargs...)
	return &MockSentryClientHandShakeCall{Call: call}
}

// MockSentryClientHandShakeCall wrap *gomock.Call
type MockSentryClientHandShakeCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientHandShakeCall) Return(arg0 *sentry.HandShakeReply, arg1 error) *MockSentryClientHandShakeCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientHandShakeCall) Do(f func(context.Context, *emptypb.Empty, ...grpc.CallOption) (*sentry.HandShakeReply, error)) *MockSentryClientHandShakeCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientHandShakeCall) DoAndReturn(f func(context.Context, *emptypb.Empty, ...grpc.CallOption) (*sentry.HandShakeReply, error)) *MockSentryClientHandShakeCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// MarkDisconnected mocks base method.
func (m *MockSentryClient) MarkDisconnected() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "MarkDisconnected")
}

// MarkDisconnected indicates an expected call of MarkDisconnected.
func (mr *MockSentryClientMockRecorder) MarkDisconnected() *MockSentryClientMarkDisconnectedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkDisconnected", reflect.TypeOf((*MockSentryClient)(nil).MarkDisconnected))
	return &MockSentryClientMarkDisconnectedCall{Call: call}
}

// MockSentryClientMarkDisconnectedCall wrap *gomock.Call
type MockSentryClientMarkDisconnectedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientMarkDisconnectedCall) Return() *MockSentryClientMarkDisconnectedCall {
	c.Call = c.Call.Return()
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientMarkDisconnectedCall) Do(f func()) *MockSentryClientMarkDisconnectedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientMarkDisconnectedCall) DoAndReturn(f func()) *MockSentryClientMarkDisconnectedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Messages mocks base method.
func (m *MockSentryClient) Messages(arg0 context.Context, arg1 *sentry.MessagesRequest, arg2 ...grpc.CallOption) (sentry.Sentry_MessagesClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Messages", varargs...)
	ret0, _ := ret[0].(sentry.Sentry_MessagesClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Messages indicates an expected call of Messages.
func (mr *MockSentryClientMockRecorder) Messages(arg0, arg1 any, arg2 ...any) *MockSentryClientMessagesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Messages", reflect.TypeOf((*MockSentryClient)(nil).Messages), varargs...)
	return &MockSentryClientMessagesCall{Call: call}
}

// MockSentryClientMessagesCall wrap *gomock.Call
type MockSentryClientMessagesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientMessagesCall) Return(arg0 sentry.Sentry_MessagesClient, arg1 error) *MockSentryClientMessagesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientMessagesCall) Do(f func(context.Context, *sentry.MessagesRequest, ...grpc.CallOption) (sentry.Sentry_MessagesClient, error)) *MockSentryClientMessagesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientMessagesCall) DoAndReturn(f func(context.Context, *sentry.MessagesRequest, ...grpc.CallOption) (sentry.Sentry_MessagesClient, error)) *MockSentryClientMessagesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// NodeInfo mocks base method.
func (m *MockSentryClient) NodeInfo(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*types.NodeInfoReply, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NodeInfo", varargs...)
	ret0, _ := ret[0].(*types.NodeInfoReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NodeInfo indicates an expected call of NodeInfo.
func (mr *MockSentryClientMockRecorder) NodeInfo(arg0, arg1 any, arg2 ...any) *MockSentryClientNodeInfoCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeInfo", reflect.TypeOf((*MockSentryClient)(nil).NodeInfo), varargs...)
	return &MockSentryClientNodeInfoCall{Call: call}
}

// MockSentryClientNodeInfoCall wrap *gomock.Call
type MockSentryClientNodeInfoCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientNodeInfoCall) Return(arg0 *types.NodeInfoReply, arg1 error) *MockSentryClientNodeInfoCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientNodeInfoCall) Do(f func(context.Context, *emptypb.Empty, ...grpc.CallOption) (*types.NodeInfoReply, error)) *MockSentryClientNodeInfoCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientNodeInfoCall) DoAndReturn(f func(context.Context, *emptypb.Empty, ...grpc.CallOption) (*types.NodeInfoReply, error)) *MockSentryClientNodeInfoCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PeerById mocks base method.
func (m *MockSentryClient) PeerById(arg0 context.Context, arg1 *sentry.PeerByIdRequest, arg2 ...grpc.CallOption) (*sentry.PeerByIdReply, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PeerById", varargs...)
	ret0, _ := ret[0].(*sentry.PeerByIdReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerById indicates an expected call of PeerById.
func (mr *MockSentryClientMockRecorder) PeerById(arg0, arg1 any, arg2 ...any) *MockSentryClientPeerByIdCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerById", reflect.TypeOf((*MockSentryClient)(nil).PeerById), varargs...)
	return &MockSentryClientPeerByIdCall{Call: call}
}

// MockSentryClientPeerByIdCall wrap *gomock.Call
type MockSentryClientPeerByIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientPeerByIdCall) Return(arg0 *sentry.PeerByIdReply, arg1 error) *MockSentryClientPeerByIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientPeerByIdCall) Do(f func(context.Context, *sentry.PeerByIdRequest, ...grpc.CallOption) (*sentry.PeerByIdReply, error)) *MockSentryClientPeerByIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientPeerByIdCall) DoAndReturn(f func(context.Context, *sentry.PeerByIdRequest, ...grpc.CallOption) (*sentry.PeerByIdReply, error)) *MockSentryClientPeerByIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PeerCount mocks base method.
func (m *MockSentryClient) PeerCount(arg0 context.Context, arg1 *sentry.PeerCountRequest, arg2 ...grpc.CallOption) (*sentry.PeerCountReply, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PeerCount", varargs...)
	ret0, _ := ret[0].(*sentry.PeerCountReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerCount indicates an expected call of PeerCount.
func (mr *MockSentryClientMockRecorder) PeerCount(arg0, arg1 any, arg2 ...any) *MockSentryClientPeerCountCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerCount", reflect.TypeOf((*MockSentryClient)(nil).PeerCount), varargs...)
	return &MockSentryClientPeerCountCall{Call: call}
}

// MockSentryClientPeerCountCall wrap *gomock.Call
type MockSentryClientPeerCountCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientPeerCountCall) Return(arg0 *sentry.PeerCountReply, arg1 error) *MockSentryClientPeerCountCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientPeerCountCall) Do(f func(context.Context, *sentry.PeerCountRequest, ...grpc.CallOption) (*sentry.PeerCountReply, error)) *MockSentryClientPeerCountCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientPeerCountCall) DoAndReturn(f func(context.Context, *sentry.PeerCountRequest, ...grpc.CallOption) (*sentry.PeerCountReply, error)) *MockSentryClientPeerCountCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PeerEvents mocks base method.
func (m *MockSentryClient) PeerEvents(arg0 context.Context, arg1 *sentry.PeerEventsRequest, arg2 ...grpc.CallOption) (sentry.Sentry_PeerEventsClient, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PeerEvents", varargs...)
	ret0, _ := ret[0].(sentry.Sentry_PeerEventsClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerEvents indicates an expected call of PeerEvents.
func (mr *MockSentryClientMockRecorder) PeerEvents(arg0, arg1 any, arg2 ...any) *MockSentryClientPeerEventsCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerEvents", reflect.TypeOf((*MockSentryClient)(nil).PeerEvents), varargs...)
	return &MockSentryClientPeerEventsCall{Call: call}
}

// MockSentryClientPeerEventsCall wrap *gomock.Call
type MockSentryClientPeerEventsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientPeerEventsCall) Return(arg0 sentry.Sentry_PeerEventsClient, arg1 error) *MockSentryClientPeerEventsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientPeerEventsCall) Do(f func(context.Context, *sentry.PeerEventsRequest, ...grpc.CallOption) (sentry.Sentry_PeerEventsClient, error)) *MockSentryClientPeerEventsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientPeerEventsCall) DoAndReturn(f func(context.Context, *sentry.PeerEventsRequest, ...grpc.CallOption) (sentry.Sentry_PeerEventsClient, error)) *MockSentryClientPeerEventsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PeerMinBlock mocks base method.
func (m *MockSentryClient) PeerMinBlock(arg0 context.Context, arg1 *sentry.PeerMinBlockRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PeerMinBlock", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerMinBlock indicates an expected call of PeerMinBlock.
func (mr *MockSentryClientMockRecorder) PeerMinBlock(arg0, arg1 any, arg2 ...any) *MockSentryClientPeerMinBlockCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerMinBlock", reflect.TypeOf((*MockSentryClient)(nil).PeerMinBlock), varargs...)
	return &MockSentryClientPeerMinBlockCall{Call: call}
}

// MockSentryClientPeerMinBlockCall wrap *gomock.Call
type MockSentryClientPeerMinBlockCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientPeerMinBlockCall) Return(arg0 *emptypb.Empty, arg1 error) *MockSentryClientPeerMinBlockCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientPeerMinBlockCall) Do(f func(context.Context, *sentry.PeerMinBlockRequest, ...grpc.CallOption) (*emptypb.Empty, error)) *MockSentryClientPeerMinBlockCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientPeerMinBlockCall) DoAndReturn(f func(context.Context, *sentry.PeerMinBlockRequest, ...grpc.CallOption) (*emptypb.Empty, error)) *MockSentryClientPeerMinBlockCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Peers mocks base method.
func (m *MockSentryClient) Peers(arg0 context.Context, arg1 *emptypb.Empty, arg2 ...grpc.CallOption) (*sentry.PeersReply, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Peers", varargs...)
	ret0, _ := ret[0].(*sentry.PeersReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Peers indicates an expected call of Peers.
func (mr *MockSentryClientMockRecorder) Peers(arg0, arg1 any, arg2 ...any) *MockSentryClientPeersCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peers", reflect.TypeOf((*MockSentryClient)(nil).Peers), varargs...)
	return &MockSentryClientPeersCall{Call: call}
}

// MockSentryClientPeersCall wrap *gomock.Call
type MockSentryClientPeersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientPeersCall) Return(arg0 *sentry.PeersReply, arg1 error) *MockSentryClientPeersCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientPeersCall) Do(f func(context.Context, *emptypb.Empty, ...grpc.CallOption) (*sentry.PeersReply, error)) *MockSentryClientPeersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientPeersCall) DoAndReturn(f func(context.Context, *emptypb.Empty, ...grpc.CallOption) (*sentry.PeersReply, error)) *MockSentryClientPeersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// PenalizePeer mocks base method.
func (m *MockSentryClient) PenalizePeer(arg0 context.Context, arg1 *sentry.PenalizePeerRequest, arg2 ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PenalizePeer", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PenalizePeer indicates an expected call of PenalizePeer.
func (mr *MockSentryClientMockRecorder) PenalizePeer(arg0, arg1 any, arg2 ...any) *MockSentryClientPenalizePeerCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PenalizePeer", reflect.TypeOf((*MockSentryClient)(nil).PenalizePeer), varargs...)
	return &MockSentryClientPenalizePeerCall{Call: call}
}

// MockSentryClientPenalizePeerCall wrap *gomock.Call
type MockSentryClientPenalizePeerCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientPenalizePeerCall) Return(arg0 *emptypb.Empty, arg1 error) *MockSentryClientPenalizePeerCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientPenalizePeerCall) Do(f func(context.Context, *sentry.PenalizePeerRequest, ...grpc.CallOption) (*emptypb.Empty, error)) *MockSentryClientPenalizePeerCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientPenalizePeerCall) DoAndReturn(f func(context.Context, *sentry.PenalizePeerRequest, ...grpc.CallOption) (*emptypb.Empty, error)) *MockSentryClientPenalizePeerCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Protocol mocks base method.
func (m *MockSentryClient) Protocol() uint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Protocol")
	ret0, _ := ret[0].(uint)
	return ret0
}

// Protocol indicates an expected call of Protocol.
func (mr *MockSentryClientMockRecorder) Protocol() *MockSentryClientProtocolCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Protocol", reflect.TypeOf((*MockSentryClient)(nil).Protocol))
	return &MockSentryClientProtocolCall{Call: call}
}

// MockSentryClientProtocolCall wrap *gomock.Call
type MockSentryClientProtocolCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientProtocolCall) Return(arg0 uint) *MockSentryClientProtocolCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientProtocolCall) Do(f func() uint) *MockSentryClientProtocolCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientProtocolCall) DoAndReturn(f func() uint) *MockSentryClientProtocolCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// Ready mocks base method.
func (m *MockSentryClient) Ready() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ready")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Ready indicates an expected call of Ready.
func (mr *MockSentryClientMockRecorder) Ready() *MockSentryClientReadyCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*MockSentryClient)(nil).Ready))
	return &MockSentryClientReadyCall{Call: call}
}

// MockSentryClientReadyCall wrap *gomock.Call
type MockSentryClientReadyCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientReadyCall) Return(arg0 bool) *MockSentryClientReadyCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientReadyCall) Do(f func() bool) *MockSentryClientReadyCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientReadyCall) DoAndReturn(f func() bool) *MockSentryClientReadyCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SendMessageById mocks base method.
func (m *MockSentryClient) SendMessageById(arg0 context.Context, arg1 *sentry.SendMessageByIdRequest, arg2 ...grpc.CallOption) (*sentry.SentPeers, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessageById", varargs...)
	ret0, _ := ret[0].(*sentry.SentPeers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessageById indicates an expected call of SendMessageById.
func (mr *MockSentryClientMockRecorder) SendMessageById(arg0, arg1 any, arg2 ...any) *MockSentryClientSendMessageByIdCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessageById", reflect.TypeOf((*MockSentryClient)(nil).SendMessageById), varargs...)
	return &MockSentryClientSendMessageByIdCall{Call: call}
}

// MockSentryClientSendMessageByIdCall wrap *gomock.Call
type MockSentryClientSendMessageByIdCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientSendMessageByIdCall) Return(arg0 *sentry.SentPeers, arg1 error) *MockSentryClientSendMessageByIdCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientSendMessageByIdCall) Do(f func(context.Context, *sentry.SendMessageByIdRequest, ...grpc.CallOption) (*sentry.SentPeers, error)) *MockSentryClientSendMessageByIdCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientSendMessageByIdCall) DoAndReturn(f func(context.Context, *sentry.SendMessageByIdRequest, ...grpc.CallOption) (*sentry.SentPeers, error)) *MockSentryClientSendMessageByIdCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SendMessageByMinBlock mocks base method.
func (m *MockSentryClient) SendMessageByMinBlock(arg0 context.Context, arg1 *sentry.SendMessageByMinBlockRequest, arg2 ...grpc.CallOption) (*sentry.SentPeers, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessageByMinBlock", varargs...)
	ret0, _ := ret[0].(*sentry.SentPeers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessageByMinBlock indicates an expected call of SendMessageByMinBlock.
func (mr *MockSentryClientMockRecorder) SendMessageByMinBlock(arg0, arg1 any, arg2 ...any) *MockSentryClientSendMessageByMinBlockCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessageByMinBlock", reflect.TypeOf((*MockSentryClient)(nil).SendMessageByMinBlock), varargs...)
	return &MockSentryClientSendMessageByMinBlockCall{Call: call}
}

// MockSentryClientSendMessageByMinBlockCall wrap *gomock.Call
type MockSentryClientSendMessageByMinBlockCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientSendMessageByMinBlockCall) Return(arg0 *sentry.SentPeers, arg1 error) *MockSentryClientSendMessageByMinBlockCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientSendMessageByMinBlockCall) Do(f func(context.Context, *sentry.SendMessageByMinBlockRequest, ...grpc.CallOption) (*sentry.SentPeers, error)) *MockSentryClientSendMessageByMinBlockCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientSendMessageByMinBlockCall) DoAndReturn(f func(context.Context, *sentry.SendMessageByMinBlockRequest, ...grpc.CallOption) (*sentry.SentPeers, error)) *MockSentryClientSendMessageByMinBlockCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SendMessageToAll mocks base method.
func (m *MockSentryClient) SendMessageToAll(arg0 context.Context, arg1 *sentry.OutboundMessageData, arg2 ...grpc.CallOption) (*sentry.SentPeers, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessageToAll", varargs...)
	ret0, _ := ret[0].(*sentry.SentPeers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessageToAll indicates an expected call of SendMessageToAll.
func (mr *MockSentryClientMockRecorder) SendMessageToAll(arg0, arg1 any, arg2 ...any) *MockSentryClientSendMessageToAllCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessageToAll", reflect.TypeOf((*MockSentryClient)(nil).SendMessageToAll), varargs...)
	return &MockSentryClientSendMessageToAllCall{Call: call}
}

// MockSentryClientSendMessageToAllCall wrap *gomock.Call
type MockSentryClientSendMessageToAllCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientSendMessageToAllCall) Return(arg0 *sentry.SentPeers, arg1 error) *MockSentryClientSendMessageToAllCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientSendMessageToAllCall) Do(f func(context.Context, *sentry.OutboundMessageData, ...grpc.CallOption) (*sentry.SentPeers, error)) *MockSentryClientSendMessageToAllCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientSendMessageToAllCall) DoAndReturn(f func(context.Context, *sentry.OutboundMessageData, ...grpc.CallOption) (*sentry.SentPeers, error)) *MockSentryClientSendMessageToAllCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SendMessageToRandomPeers mocks base method.
func (m *MockSentryClient) SendMessageToRandomPeers(arg0 context.Context, arg1 *sentry.SendMessageToRandomPeersRequest, arg2 ...grpc.CallOption) (*sentry.SentPeers, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SendMessageToRandomPeers", varargs...)
	ret0, _ := ret[0].(*sentry.SentPeers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendMessageToRandomPeers indicates an expected call of SendMessageToRandomPeers.
func (mr *MockSentryClientMockRecorder) SendMessageToRandomPeers(arg0, arg1 any, arg2 ...any) *MockSentryClientSendMessageToRandomPeersCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessageToRandomPeers", reflect.TypeOf((*MockSentryClient)(nil).SendMessageToRandomPeers), varargs...)
	return &MockSentryClientSendMessageToRandomPeersCall{Call: call}
}

// MockSentryClientSendMessageToRandomPeersCall wrap *gomock.Call
type MockSentryClientSendMessageToRandomPeersCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientSendMessageToRandomPeersCall) Return(arg0 *sentry.SentPeers, arg1 error) *MockSentryClientSendMessageToRandomPeersCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientSendMessageToRandomPeersCall) Do(f func(context.Context, *sentry.SendMessageToRandomPeersRequest, ...grpc.CallOption) (*sentry.SentPeers, error)) *MockSentryClientSendMessageToRandomPeersCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientSendMessageToRandomPeersCall) DoAndReturn(f func(context.Context, *sentry.SendMessageToRandomPeersRequest, ...grpc.CallOption) (*sentry.SentPeers, error)) *MockSentryClientSendMessageToRandomPeersCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetStatus mocks base method.
func (m *MockSentryClient) SetStatus(arg0 context.Context, arg1 *sentry.StatusData, arg2 ...grpc.CallOption) (*sentry.SetStatusReply, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetStatus", varargs...)
	ret0, _ := ret[0].(*sentry.SetStatusReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetStatus indicates an expected call of SetStatus.
func (mr *MockSentryClientMockRecorder) SetStatus(arg0, arg1 any, arg2 ...any) *MockSentryClientSetStatusCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1}, arg2...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetStatus", reflect.TypeOf((*MockSentryClient)(nil).SetStatus), varargs...)
	return &MockSentryClientSetStatusCall{Call: call}
}

// MockSentryClientSetStatusCall wrap *gomock.Call
type MockSentryClientSetStatusCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockSentryClientSetStatusCall) Return(arg0 *sentry.SetStatusReply, arg1 error) *MockSentryClientSetStatusCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockSentryClientSetStatusCall) Do(f func(context.Context, *sentry.StatusData, ...grpc.CallOption) (*sentry.SetStatusReply, error)) *MockSentryClientSetStatusCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockSentryClientSetStatusCall) DoAndReturn(f func(context.Context, *sentry.StatusData, ...grpc.CallOption) (*sentry.SetStatusReply, error)) *MockSentryClientSetStatusCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}