// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/scionproto/scion/go/lib/snet (interfaces: Conn,PacketDispatcherService,Network,PacketConn,Path,Router)

// Package mock_snet is a generated GoMock package.
package mock_snet

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	addr "github.com/scionproto/scion/go/lib/addr"
	overlay "github.com/scionproto/scion/go/lib/overlay"
	snet "github.com/scionproto/scion/go/lib/snet"
	spath "github.com/scionproto/scion/go/lib/spath"
	net "net"
	reflect "reflect"
	time "time"
)

// MockConn is a mock of Conn interface
type MockConn struct {
	ctrl     *gomock.Controller
	recorder *MockConnMockRecorder
}

// MockConnMockRecorder is the mock recorder for MockConn
type MockConnMockRecorder struct {
	mock *MockConn
}

// NewMockConn creates a new mock instance
func NewMockConn(ctrl *gomock.Controller) *MockConn {
	mock := &MockConn{ctrl: ctrl}
	mock.recorder = &MockConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockConn) EXPECT() *MockConnMockRecorder {
	return m.recorder
}

// BindAddr mocks base method
func (m *MockConn) BindAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// BindAddr indicates an expected call of BindAddr
func (mr *MockConnMockRecorder) BindAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindAddr", reflect.TypeOf((*MockConn)(nil).BindAddr))
}

// Close mocks base method
func (m *MockConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockConnMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockConn)(nil).Close))
}

// LocalAddr mocks base method
func (m *MockConn) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr
func (mr *MockConnMockRecorder) LocalAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockConn)(nil).LocalAddr))
}

// Read mocks base method
func (m *MockConn) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockConnMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockConn)(nil).Read), arg0)
}

// ReadFrom mocks base method
func (m *MockConn) ReadFrom(arg0 []byte) (int, net.Addr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFrom", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(net.Addr)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadFrom indicates an expected call of ReadFrom
func (mr *MockConnMockRecorder) ReadFrom(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFrom", reflect.TypeOf((*MockConn)(nil).ReadFrom), arg0)
}

// ReadFromSCION mocks base method
func (m *MockConn) ReadFromSCION(arg0 []byte) (int, *snet.Addr, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFromSCION", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(*snet.Addr)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ReadFromSCION indicates an expected call of ReadFromSCION
func (mr *MockConnMockRecorder) ReadFromSCION(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFromSCION", reflect.TypeOf((*MockConn)(nil).ReadFromSCION), arg0)
}

// RemoteAddr mocks base method
func (m *MockConn) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr
func (mr *MockConnMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockConn)(nil).RemoteAddr))
}

// SVC mocks base method
func (m *MockConn) SVC() addr.HostSVC {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SVC")
	ret0, _ := ret[0].(addr.HostSVC)
	return ret0
}

// SVC indicates an expected call of SVC
func (mr *MockConnMockRecorder) SVC() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SVC", reflect.TypeOf((*MockConn)(nil).SVC))
}

// SetDeadline mocks base method
func (m *MockConn) SetDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline
func (mr *MockConnMockRecorder) SetDeadline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadline", reflect.TypeOf((*MockConn)(nil).SetDeadline), arg0)
}

// SetReadDeadline mocks base method
func (m *MockConn) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline
func (mr *MockConnMockRecorder) SetReadDeadline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockConn)(nil).SetReadDeadline), arg0)
}

// SetWriteDeadline mocks base method
func (m *MockConn) SetWriteDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline
func (mr *MockConnMockRecorder) SetWriteDeadline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockConn)(nil).SetWriteDeadline), arg0)
}

// Write mocks base method
func (m *MockConn) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockConnMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockConn)(nil).Write), arg0)
}

// WriteTo mocks base method
func (m *MockConn) WriteTo(arg0 []byte, arg1 net.Addr) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTo", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteTo indicates an expected call of WriteTo
func (mr *MockConnMockRecorder) WriteTo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTo", reflect.TypeOf((*MockConn)(nil).WriteTo), arg0, arg1)
}

// WriteToSCION mocks base method
func (m *MockConn) WriteToSCION(arg0 []byte, arg1 *snet.Addr) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteToSCION", arg0, arg1)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteToSCION indicates an expected call of WriteToSCION
func (mr *MockConnMockRecorder) WriteToSCION(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteToSCION", reflect.TypeOf((*MockConn)(nil).WriteToSCION), arg0, arg1)
}

// MockPacketDispatcherService is a mock of PacketDispatcherService interface
type MockPacketDispatcherService struct {
	ctrl     *gomock.Controller
	recorder *MockPacketDispatcherServiceMockRecorder
}

// MockPacketDispatcherServiceMockRecorder is the mock recorder for MockPacketDispatcherService
type MockPacketDispatcherServiceMockRecorder struct {
	mock *MockPacketDispatcherService
}

// NewMockPacketDispatcherService creates a new mock instance
func NewMockPacketDispatcherService(ctrl *gomock.Controller) *MockPacketDispatcherService {
	mock := &MockPacketDispatcherService{ctrl: ctrl}
	mock.recorder = &MockPacketDispatcherServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPacketDispatcherService) EXPECT() *MockPacketDispatcherServiceMockRecorder {
	return m.recorder
}

// RegisterTimeout mocks base method
func (m *MockPacketDispatcherService) RegisterTimeout(arg0 addr.IA, arg1 *addr.AppAddr, arg2 *overlay.OverlayAddr, arg3 addr.HostSVC, arg4 time.Duration) (snet.PacketConn, uint16, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterTimeout", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(snet.PacketConn)
	ret1, _ := ret[1].(uint16)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// RegisterTimeout indicates an expected call of RegisterTimeout
func (mr *MockPacketDispatcherServiceMockRecorder) RegisterTimeout(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterTimeout", reflect.TypeOf((*MockPacketDispatcherService)(nil).RegisterTimeout), arg0, arg1, arg2, arg3, arg4)
}

// MockNetwork is a mock of Network interface
type MockNetwork struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkMockRecorder
}

// MockNetworkMockRecorder is the mock recorder for MockNetwork
type MockNetworkMockRecorder struct {
	mock *MockNetwork
}

// NewMockNetwork creates a new mock instance
func NewMockNetwork(ctrl *gomock.Controller) *MockNetwork {
	mock := &MockNetwork{ctrl: ctrl}
	mock.recorder = &MockNetworkMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetwork) EXPECT() *MockNetworkMockRecorder {
	return m.recorder
}

// DialSCIONWithBindSVC mocks base method
func (m *MockNetwork) DialSCIONWithBindSVC(arg0 string, arg1, arg2, arg3 *snet.Addr, arg4 addr.HostSVC, arg5 time.Duration) (snet.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DialSCIONWithBindSVC", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(snet.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DialSCIONWithBindSVC indicates an expected call of DialSCIONWithBindSVC
func (mr *MockNetworkMockRecorder) DialSCIONWithBindSVC(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DialSCIONWithBindSVC", reflect.TypeOf((*MockNetwork)(nil).DialSCIONWithBindSVC), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ListenSCIONWithBindSVC mocks base method
func (m *MockNetwork) ListenSCIONWithBindSVC(arg0 string, arg1, arg2 *snet.Addr, arg3 addr.HostSVC, arg4 time.Duration) (snet.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenSCIONWithBindSVC", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(snet.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListenSCIONWithBindSVC indicates an expected call of ListenSCIONWithBindSVC
func (mr *MockNetworkMockRecorder) ListenSCIONWithBindSVC(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenSCIONWithBindSVC", reflect.TypeOf((*MockNetwork)(nil).ListenSCIONWithBindSVC), arg0, arg1, arg2, arg3, arg4)
}

// MockPacketConn is a mock of PacketConn interface
type MockPacketConn struct {
	ctrl     *gomock.Controller
	recorder *MockPacketConnMockRecorder
}

// MockPacketConnMockRecorder is the mock recorder for MockPacketConn
type MockPacketConnMockRecorder struct {
	mock *MockPacketConn
}

// NewMockPacketConn creates a new mock instance
func NewMockPacketConn(ctrl *gomock.Controller) *MockPacketConn {
	mock := &MockPacketConn{ctrl: ctrl}
	mock.recorder = &MockPacketConnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPacketConn) EXPECT() *MockPacketConnMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockPacketConn) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockPacketConnMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPacketConn)(nil).Close))
}

// ReadFrom mocks base method
func (m *MockPacketConn) ReadFrom(arg0 *snet.SCIONPacket, arg1 *overlay.OverlayAddr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFrom", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadFrom indicates an expected call of ReadFrom
func (mr *MockPacketConnMockRecorder) ReadFrom(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFrom", reflect.TypeOf((*MockPacketConn)(nil).ReadFrom), arg0, arg1)
}

// SetDeadline mocks base method
func (m *MockPacketConn) SetDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDeadline indicates an expected call of SetDeadline
func (mr *MockPacketConnMockRecorder) SetDeadline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDeadline", reflect.TypeOf((*MockPacketConn)(nil).SetDeadline), arg0)
}

// SetReadDeadline mocks base method
func (m *MockPacketConn) SetReadDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetReadDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetReadDeadline indicates an expected call of SetReadDeadline
func (mr *MockPacketConnMockRecorder) SetReadDeadline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadDeadline", reflect.TypeOf((*MockPacketConn)(nil).SetReadDeadline), arg0)
}

// SetWriteDeadline mocks base method
func (m *MockPacketConn) SetWriteDeadline(arg0 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWriteDeadline", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWriteDeadline indicates an expected call of SetWriteDeadline
func (mr *MockPacketConnMockRecorder) SetWriteDeadline(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWriteDeadline", reflect.TypeOf((*MockPacketConn)(nil).SetWriteDeadline), arg0)
}

// WriteTo mocks base method
func (m *MockPacketConn) WriteTo(arg0 *snet.SCIONPacket, arg1 *overlay.OverlayAddr) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteTo indicates an expected call of WriteTo
func (mr *MockPacketConnMockRecorder) WriteTo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTo", reflect.TypeOf((*MockPacketConn)(nil).WriteTo), arg0, arg1)
}

// MockPath is a mock of Path interface
type MockPath struct {
	ctrl     *gomock.Controller
	recorder *MockPathMockRecorder
}

// MockPathMockRecorder is the mock recorder for MockPath
type MockPathMockRecorder struct {
	mock *MockPath
}

// NewMockPath creates a new mock instance
func NewMockPath(ctrl *gomock.Controller) *MockPath {
	mock := &MockPath{ctrl: ctrl}
	mock.recorder = &MockPathMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPath) EXPECT() *MockPathMockRecorder {
	return m.recorder
}

// Destination mocks base method
func (m *MockPath) Destination() addr.IA {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destination")
	ret0, _ := ret[0].(addr.IA)
	return ret0
}

// Destination indicates an expected call of Destination
func (mr *MockPathMockRecorder) Destination() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destination", reflect.TypeOf((*MockPath)(nil).Destination))
}

// OverlayNextHop mocks base method
func (m *MockPath) OverlayNextHop() *overlay.OverlayAddr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OverlayNextHop")
	ret0, _ := ret[0].(*overlay.OverlayAddr)
	return ret0
}

// OverlayNextHop indicates an expected call of OverlayNextHop
func (mr *MockPathMockRecorder) OverlayNextHop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OverlayNextHop", reflect.TypeOf((*MockPath)(nil).OverlayNextHop))
}

// Path mocks base method
func (m *MockPath) Path() *spath.Path {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Path")
	ret0, _ := ret[0].(*spath.Path)
	return ret0
}

// Path indicates an expected call of Path
func (mr *MockPathMockRecorder) Path() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Path", reflect.TypeOf((*MockPath)(nil).Path))
}

// MockRouter is a mock of Router interface
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// AllRoutes mocks base method
func (m *MockRouter) AllRoutes(arg0 context.Context, arg1 addr.IA) ([]snet.Path, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllRoutes", arg0, arg1)
	ret0, _ := ret[0].([]snet.Path)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllRoutes indicates an expected call of AllRoutes
func (mr *MockRouterMockRecorder) AllRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllRoutes", reflect.TypeOf((*MockRouter)(nil).AllRoutes), arg0, arg1)
}

// LocalIA mocks base method
func (m *MockRouter) LocalIA() addr.IA {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalIA")
	ret0, _ := ret[0].(addr.IA)
	return ret0
}

// LocalIA indicates an expected call of LocalIA
func (mr *MockRouterMockRecorder) LocalIA() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalIA", reflect.TypeOf((*MockRouter)(nil).LocalIA))
}

// Route mocks base method
func (m *MockRouter) Route(arg0 context.Context, arg1 addr.IA) (snet.Path, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Route", arg0, arg1)
	ret0, _ := ret[0].(snet.Path)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Route indicates an expected call of Route
func (mr *MockRouterMockRecorder) Route(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Route", reflect.TypeOf((*MockRouter)(nil).Route), arg0, arg1)
}
