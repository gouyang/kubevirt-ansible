// Automatically generated by MockGen. DO NOT EDIT!
// Source: client.go

package cmdclient

import (
	gomock "github.com/golang/mock/gomock"

	v1 "kubevirt.io/kubevirt/pkg/api/v1"
	api "kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/api"
)

// Mock of LauncherClient interface
type MockLauncherClient struct {
	ctrl     *gomock.Controller
	recorder *_MockLauncherClientRecorder
}

// Recorder for MockLauncherClient (not exported)
type _MockLauncherClientRecorder struct {
	mock *MockLauncherClient
}

func NewMockLauncherClient(ctrl *gomock.Controller) *MockLauncherClient {
	mock := &MockLauncherClient{ctrl: ctrl}
	mock.recorder = &_MockLauncherClientRecorder{mock}
	return mock
}

func (_m *MockLauncherClient) EXPECT() *_MockLauncherClientRecorder {
	return _m.recorder
}

func (_m *MockLauncherClient) SyncVirtualMachine(vmi *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "SyncVirtualMachine", vmi)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLauncherClientRecorder) SyncVirtualMachine(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SyncVirtualMachine", arg0)
}

func (_m *MockLauncherClient) SyncMigrationTarget(vmi *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "SyncMigrationTarget", vmi)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLauncherClientRecorder) SyncMigrationTarget(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SyncMigrationTarget", arg0)
}

func (_m *MockLauncherClient) ShutdownVirtualMachine(vmi *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "ShutdownVirtualMachine", vmi)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLauncherClientRecorder) ShutdownVirtualMachine(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ShutdownVirtualMachine", arg0)
}

func (_m *MockLauncherClient) KillVirtualMachine(vmi *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "KillVirtualMachine", vmi)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLauncherClientRecorder) KillVirtualMachine(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "KillVirtualMachine", arg0)
}

func (_m *MockLauncherClient) MigrateVirtualMachine(vmi *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "MigrateVirtualMachine", vmi)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLauncherClientRecorder) MigrateVirtualMachine(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MigrateVirtualMachine", arg0)
}

func (_m *MockLauncherClient) DeleteDomain(vmi *v1.VirtualMachineInstance) error {
	ret := _m.ctrl.Call(_m, "DeleteDomain", vmi)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLauncherClientRecorder) DeleteDomain(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteDomain", arg0)
}

func (_m *MockLauncherClient) GetDomain() (*api.Domain, bool, error) {
	ret := _m.ctrl.Call(_m, "GetDomain")
	ret0, _ := ret[0].(*api.Domain)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockLauncherClientRecorder) GetDomain() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDomain")
}

func (_m *MockLauncherClient) Ping() error {
	ret := _m.ctrl.Call(_m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLauncherClientRecorder) Ping() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ping")
}

func (_m *MockLauncherClient) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockLauncherClientRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}
