// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/go/lib/infra (interfaces: TrustStore,Messenger,ResponseWriter)

// Package mock_infra is a generated GoMock package.
package mock_infra

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	addr "github.com/scionproto/scion/go/lib/addr"
	common "github.com/scionproto/scion/go/lib/common"
	ctrl "github.com/scionproto/scion/go/lib/ctrl"
	ack "github.com/scionproto/scion/go/lib/ctrl/ack"
	cert_mgmt "github.com/scionproto/scion/go/lib/ctrl/cert_mgmt"
	ifid "github.com/scionproto/scion/go/lib/ctrl/ifid"
	path_mgmt "github.com/scionproto/scion/go/lib/ctrl/path_mgmt"
	seg "github.com/scionproto/scion/go/lib/ctrl/seg"
	infra "github.com/scionproto/scion/go/lib/infra"
	cert "github.com/scionproto/scion/go/lib/scrypto/cert"
	trc "github.com/scionproto/scion/go/lib/scrypto/trc"
	proto "github.com/scionproto/scion/go/proto"
	net "net"
	reflect "reflect"
)

// MockTrustStore is a mock of TrustStore interface
type MockTrustStore struct {
	ctrl     *gomock.Controller
	recorder *MockTrustStoreMockRecorder
}

// MockTrustStoreMockRecorder is the mock recorder for MockTrustStore
type MockTrustStoreMockRecorder struct {
	mock *MockTrustStore
}

// NewMockTrustStore creates a new mock instance
func NewMockTrustStore(ctrl *gomock.Controller) *MockTrustStore {
	mock := &MockTrustStore{ctrl: ctrl}
	mock.recorder = &MockTrustStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrustStore) EXPECT() *MockTrustStoreMockRecorder {
	return m.recorder
}

// GetChain mocks base method
func (m *MockTrustStore) GetChain(arg0 context.Context, arg1 addr.IA, arg2 uint64) (*cert.Chain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChain", arg0, arg1, arg2)
	ret0, _ := ret[0].(*cert.Chain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChain indicates an expected call of GetChain
func (mr *MockTrustStoreMockRecorder) GetChain(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChain", reflect.TypeOf((*MockTrustStore)(nil).GetChain), arg0, arg1, arg2)
}

// GetTRC mocks base method
func (m *MockTrustStore) GetTRC(arg0 context.Context, arg1 addr.ISD, arg2 uint64) (*trc.TRC, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTRC", arg0, arg1, arg2)
	ret0, _ := ret[0].(*trc.TRC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTRC indicates an expected call of GetTRC
func (mr *MockTrustStoreMockRecorder) GetTRC(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTRC", reflect.TypeOf((*MockTrustStore)(nil).GetTRC), arg0, arg1, arg2)
}

// GetValidCachedTRC mocks base method
func (m *MockTrustStore) GetValidCachedTRC(arg0 context.Context, arg1 addr.ISD) (*trc.TRC, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidCachedTRC", arg0, arg1)
	ret0, _ := ret[0].(*trc.TRC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidCachedTRC indicates an expected call of GetValidCachedTRC
func (mr *MockTrustStoreMockRecorder) GetValidCachedTRC(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidCachedTRC", reflect.TypeOf((*MockTrustStore)(nil).GetValidCachedTRC), arg0, arg1)
}

// GetValidChain mocks base method
func (m *MockTrustStore) GetValidChain(arg0 context.Context, arg1 addr.IA, arg2 net.Addr) (*cert.Chain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidChain", arg0, arg1, arg2)
	ret0, _ := ret[0].(*cert.Chain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidChain indicates an expected call of GetValidChain
func (mr *MockTrustStoreMockRecorder) GetValidChain(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidChain", reflect.TypeOf((*MockTrustStore)(nil).GetValidChain), arg0, arg1, arg2)
}

// GetValidTRC mocks base method
func (m *MockTrustStore) GetValidTRC(arg0 context.Context, arg1 addr.ISD, arg2 net.Addr) (*trc.TRC, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValidTRC", arg0, arg1, arg2)
	ret0, _ := ret[0].(*trc.TRC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValidTRC indicates an expected call of GetValidTRC
func (mr *MockTrustStoreMockRecorder) GetValidTRC(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValidTRC", reflect.TypeOf((*MockTrustStore)(nil).GetValidTRC), arg0, arg1, arg2)
}

// NewChainReqHandler mocks base method
func (m *MockTrustStore) NewChainReqHandler(arg0 bool) infra.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewChainReqHandler", arg0)
	ret0, _ := ret[0].(infra.Handler)
	return ret0
}

// NewChainReqHandler indicates an expected call of NewChainReqHandler
func (mr *MockTrustStoreMockRecorder) NewChainReqHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewChainReqHandler", reflect.TypeOf((*MockTrustStore)(nil).NewChainReqHandler), arg0)
}

// NewSigVerifier mocks base method
func (m *MockTrustStore) NewSigVerifier() ctrl.SigVerifier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSigVerifier")
	ret0, _ := ret[0].(ctrl.SigVerifier)
	return ret0
}

// NewSigVerifier indicates an expected call of NewSigVerifier
func (mr *MockTrustStoreMockRecorder) NewSigVerifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSigVerifier", reflect.TypeOf((*MockTrustStore)(nil).NewSigVerifier))
}

// NewSigner mocks base method
func (m *MockTrustStore) NewSigner(arg0 *proto.SignS, arg1 common.RawBytes) ctrl.Signer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSigner", arg0, arg1)
	ret0, _ := ret[0].(ctrl.Signer)
	return ret0
}

// NewSigner indicates an expected call of NewSigner
func (mr *MockTrustStoreMockRecorder) NewSigner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSigner", reflect.TypeOf((*MockTrustStore)(nil).NewSigner), arg0, arg1)
}

// NewTRCReqHandler mocks base method
func (m *MockTrustStore) NewTRCReqHandler(arg0 bool) infra.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTRCReqHandler", arg0)
	ret0, _ := ret[0].(infra.Handler)
	return ret0
}

// NewTRCReqHandler indicates an expected call of NewTRCReqHandler
func (mr *MockTrustStoreMockRecorder) NewTRCReqHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTRCReqHandler", reflect.TypeOf((*MockTrustStore)(nil).NewTRCReqHandler), arg0)
}

// SetMessenger mocks base method
func (m *MockTrustStore) SetMessenger(arg0 infra.Messenger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMessenger", arg0)
}

// SetMessenger indicates an expected call of SetMessenger
func (mr *MockTrustStoreMockRecorder) SetMessenger(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMessenger", reflect.TypeOf((*MockTrustStore)(nil).SetMessenger), arg0)
}

// MockMessenger is a mock of Messenger interface
type MockMessenger struct {
	ctrl     *gomock.Controller
	recorder *MockMessengerMockRecorder
}

// MockMessengerMockRecorder is the mock recorder for MockMessenger
type MockMessengerMockRecorder struct {
	mock *MockMessenger
}

// NewMockMessenger creates a new mock instance
func NewMockMessenger(ctrl *gomock.Controller) *MockMessenger {
	mock := &MockMessenger{ctrl: ctrl}
	mock.recorder = &MockMessengerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessenger) EXPECT() *MockMessengerMockRecorder {
	return m.recorder
}

// AddHandler mocks base method
func (m *MockMessenger) AddHandler(arg0 infra.MessageType, arg1 infra.Handler) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddHandler", arg0, arg1)
}

// AddHandler indicates an expected call of AddHandler
func (mr *MockMessengerMockRecorder) AddHandler(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddHandler", reflect.TypeOf((*MockMessenger)(nil).AddHandler), arg0, arg1)
}

// CloseServer mocks base method
func (m *MockMessenger) CloseServer() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseServer")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseServer indicates an expected call of CloseServer
func (mr *MockMessengerMockRecorder) CloseServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseServer", reflect.TypeOf((*MockMessenger)(nil).CloseServer))
}

// GetCertChain mocks base method
func (m *MockMessenger) GetCertChain(arg0 context.Context, arg1 *cert_mgmt.ChainReq, arg2 net.Addr, arg3 uint64) (*cert_mgmt.Chain, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCertChain", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*cert_mgmt.Chain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertChain indicates an expected call of GetCertChain
func (mr *MockMessengerMockRecorder) GetCertChain(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertChain", reflect.TypeOf((*MockMessenger)(nil).GetCertChain), arg0, arg1, arg2, arg3)
}

// GetSegChanges mocks base method
func (m *MockMessenger) GetSegChanges(arg0 context.Context, arg1 *path_mgmt.SegChangesReq, arg2 net.Addr, arg3 uint64) (*path_mgmt.SegChangesReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSegChanges", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*path_mgmt.SegChangesReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSegChanges indicates an expected call of GetSegChanges
func (mr *MockMessengerMockRecorder) GetSegChanges(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSegChanges", reflect.TypeOf((*MockMessenger)(nil).GetSegChanges), arg0, arg1, arg2, arg3)
}

// GetSegChangesIds mocks base method
func (m *MockMessenger) GetSegChangesIds(arg0 context.Context, arg1 *path_mgmt.SegChangesIdReq, arg2 net.Addr, arg3 uint64) (*path_mgmt.SegChangesIdReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSegChangesIds", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*path_mgmt.SegChangesIdReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSegChangesIds indicates an expected call of GetSegChangesIds
func (mr *MockMessengerMockRecorder) GetSegChangesIds(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSegChangesIds", reflect.TypeOf((*MockMessenger)(nil).GetSegChangesIds), arg0, arg1, arg2, arg3)
}

// GetSegs mocks base method
func (m *MockMessenger) GetSegs(arg0 context.Context, arg1 *path_mgmt.SegReq, arg2 net.Addr, arg3 uint64) (*path_mgmt.SegReply, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSegs", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*path_mgmt.SegReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSegs indicates an expected call of GetSegs
func (mr *MockMessengerMockRecorder) GetSegs(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSegs", reflect.TypeOf((*MockMessenger)(nil).GetSegs), arg0, arg1, arg2, arg3)
}

// GetTRC mocks base method
func (m *MockMessenger) GetTRC(arg0 context.Context, arg1 *cert_mgmt.TRCReq, arg2 net.Addr, arg3 uint64) (*cert_mgmt.TRC, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTRC", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*cert_mgmt.TRC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTRC indicates an expected call of GetTRC
func (mr *MockMessengerMockRecorder) GetTRC(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTRC", reflect.TypeOf((*MockMessenger)(nil).GetTRC), arg0, arg1, arg2, arg3)
}

// ListenAndServe mocks base method
func (m *MockMessenger) ListenAndServe() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ListenAndServe")
}

// ListenAndServe indicates an expected call of ListenAndServe
func (mr *MockMessengerMockRecorder) ListenAndServe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAndServe", reflect.TypeOf((*MockMessenger)(nil).ListenAndServe))
}

// RequestChainIssue mocks base method
func (m *MockMessenger) RequestChainIssue(arg0 context.Context, arg1 *cert_mgmt.ChainIssReq, arg2 net.Addr, arg3 uint64) (*cert_mgmt.ChainIssRep, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestChainIssue", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*cert_mgmt.ChainIssRep)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestChainIssue indicates an expected call of RequestChainIssue
func (mr *MockMessengerMockRecorder) RequestChainIssue(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestChainIssue", reflect.TypeOf((*MockMessenger)(nil).RequestChainIssue), arg0, arg1, arg2, arg3)
}

// SendAck mocks base method
func (m *MockMessenger) SendAck(arg0 context.Context, arg1 *ack.Ack, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAck", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAck indicates an expected call of SendAck
func (mr *MockMessengerMockRecorder) SendAck(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAck", reflect.TypeOf((*MockMessenger)(nil).SendAck), arg0, arg1, arg2, arg3)
}

// SendCertChain mocks base method
func (m *MockMessenger) SendCertChain(arg0 context.Context, arg1 *cert_mgmt.Chain, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCertChain", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCertChain indicates an expected call of SendCertChain
func (mr *MockMessengerMockRecorder) SendCertChain(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCertChain", reflect.TypeOf((*MockMessenger)(nil).SendCertChain), arg0, arg1, arg2, arg3)
}

// SendChainIssueReply mocks base method
func (m *MockMessenger) SendChainIssueReply(arg0 context.Context, arg1 *cert_mgmt.ChainIssRep, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendChainIssueReply", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendChainIssueReply indicates an expected call of SendChainIssueReply
func (mr *MockMessengerMockRecorder) SendChainIssueReply(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendChainIssueReply", reflect.TypeOf((*MockMessenger)(nil).SendChainIssueReply), arg0, arg1, arg2, arg3)
}

// SendIfId mocks base method
func (m *MockMessenger) SendIfId(arg0 context.Context, arg1 *ifid.IFID, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendIfId", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendIfId indicates an expected call of SendIfId
func (mr *MockMessengerMockRecorder) SendIfId(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendIfId", reflect.TypeOf((*MockMessenger)(nil).SendIfId), arg0, arg1, arg2, arg3)
}

// SendIfStateInfos mocks base method
func (m *MockMessenger) SendIfStateInfos(arg0 context.Context, arg1 *path_mgmt.IFStateInfos, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendIfStateInfos", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendIfStateInfos indicates an expected call of SendIfStateInfos
func (mr *MockMessengerMockRecorder) SendIfStateInfos(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendIfStateInfos", reflect.TypeOf((*MockMessenger)(nil).SendIfStateInfos), arg0, arg1, arg2, arg3)
}

// SendSeg mocks base method
func (m *MockMessenger) SendSeg(arg0 context.Context, arg1 *seg.PathSegment, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSeg", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendSeg indicates an expected call of SendSeg
func (mr *MockMessengerMockRecorder) SendSeg(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSeg", reflect.TypeOf((*MockMessenger)(nil).SendSeg), arg0, arg1, arg2, arg3)
}

// SendSegChangesIdReply mocks base method
func (m *MockMessenger) SendSegChangesIdReply(arg0 context.Context, arg1 *path_mgmt.SegChangesIdReply, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSegChangesIdReply", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendSegChangesIdReply indicates an expected call of SendSegChangesIdReply
func (mr *MockMessengerMockRecorder) SendSegChangesIdReply(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSegChangesIdReply", reflect.TypeOf((*MockMessenger)(nil).SendSegChangesIdReply), arg0, arg1, arg2, arg3)
}

// SendSegChangesReply mocks base method
func (m *MockMessenger) SendSegChangesReply(arg0 context.Context, arg1 *path_mgmt.SegChangesReply, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSegChangesReply", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendSegChangesReply indicates an expected call of SendSegChangesReply
func (mr *MockMessengerMockRecorder) SendSegChangesReply(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSegChangesReply", reflect.TypeOf((*MockMessenger)(nil).SendSegChangesReply), arg0, arg1, arg2, arg3)
}

// SendSegReply mocks base method
func (m *MockMessenger) SendSegReply(arg0 context.Context, arg1 *path_mgmt.SegReply, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSegReply", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendSegReply indicates an expected call of SendSegReply
func (mr *MockMessengerMockRecorder) SendSegReply(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSegReply", reflect.TypeOf((*MockMessenger)(nil).SendSegReply), arg0, arg1, arg2, arg3)
}

// SendSegSync mocks base method
func (m *MockMessenger) SendSegSync(arg0 context.Context, arg1 *path_mgmt.SegSync, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSegSync", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendSegSync indicates an expected call of SendSegSync
func (mr *MockMessengerMockRecorder) SendSegSync(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSegSync", reflect.TypeOf((*MockMessenger)(nil).SendSegSync), arg0, arg1, arg2, arg3)
}

// SendTRC mocks base method
func (m *MockMessenger) SendTRC(arg0 context.Context, arg1 *cert_mgmt.TRC, arg2 net.Addr, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTRC", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTRC indicates an expected call of SendTRC
func (mr *MockMessengerMockRecorder) SendTRC(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTRC", reflect.TypeOf((*MockMessenger)(nil).SendTRC), arg0, arg1, arg2, arg3)
}

// UpdateSigner mocks base method
func (m *MockMessenger) UpdateSigner(arg0 ctrl.Signer, arg1 []infra.MessageType) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateSigner", arg0, arg1)
}

// UpdateSigner indicates an expected call of UpdateSigner
func (mr *MockMessengerMockRecorder) UpdateSigner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSigner", reflect.TypeOf((*MockMessenger)(nil).UpdateSigner), arg0, arg1)
}

// UpdateVerifier mocks base method
func (m *MockMessenger) UpdateVerifier(arg0 ctrl.SigVerifier) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateVerifier", arg0)
}

// UpdateVerifier indicates an expected call of UpdateVerifier
func (mr *MockMessengerMockRecorder) UpdateVerifier(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVerifier", reflect.TypeOf((*MockMessenger)(nil).UpdateVerifier), arg0)
}

// MockResponseWriter is a mock of ResponseWriter interface
type MockResponseWriter struct {
	ctrl     *gomock.Controller
	recorder *MockResponseWriterMockRecorder
}

// MockResponseWriterMockRecorder is the mock recorder for MockResponseWriter
type MockResponseWriterMockRecorder struct {
	mock *MockResponseWriter
}

// NewMockResponseWriter creates a new mock instance
func NewMockResponseWriter(ctrl *gomock.Controller) *MockResponseWriter {
	mock := &MockResponseWriter{ctrl: ctrl}
	mock.recorder = &MockResponseWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResponseWriter) EXPECT() *MockResponseWriterMockRecorder {
	return m.recorder
}

// SendAckReply mocks base method
func (m *MockResponseWriter) SendAckReply(arg0 context.Context, arg1 *ack.Ack) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendAckReply", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendAckReply indicates an expected call of SendAckReply
func (mr *MockResponseWriterMockRecorder) SendAckReply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendAckReply", reflect.TypeOf((*MockResponseWriter)(nil).SendAckReply), arg0, arg1)
}

// SendCertChainReply mocks base method
func (m *MockResponseWriter) SendCertChainReply(arg0 context.Context, arg1 *cert_mgmt.Chain) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCertChainReply", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCertChainReply indicates an expected call of SendCertChainReply
func (mr *MockResponseWriterMockRecorder) SendCertChainReply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCertChainReply", reflect.TypeOf((*MockResponseWriter)(nil).SendCertChainReply), arg0, arg1)
}

// SendChainIssueReply mocks base method
func (m *MockResponseWriter) SendChainIssueReply(arg0 context.Context, arg1 *cert_mgmt.ChainIssRep) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendChainIssueReply", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendChainIssueReply indicates an expected call of SendChainIssueReply
func (mr *MockResponseWriterMockRecorder) SendChainIssueReply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendChainIssueReply", reflect.TypeOf((*MockResponseWriter)(nil).SendChainIssueReply), arg0, arg1)
}

// SendSegReply mocks base method
func (m *MockResponseWriter) SendSegReply(arg0 context.Context, arg1 *path_mgmt.SegReply) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendSegReply", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendSegReply indicates an expected call of SendSegReply
func (mr *MockResponseWriterMockRecorder) SendSegReply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendSegReply", reflect.TypeOf((*MockResponseWriter)(nil).SendSegReply), arg0, arg1)
}

// SendTRCReply mocks base method
func (m *MockResponseWriter) SendTRCReply(arg0 context.Context, arg1 *cert_mgmt.TRC) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTRCReply", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTRCReply indicates an expected call of SendTRCReply
func (mr *MockResponseWriterMockRecorder) SendTRCReply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTRCReply", reflect.TypeOf((*MockResponseWriter)(nil).SendTRCReply), arg0, arg1)
}
