// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/assisted-service/internal/provider (interfaces: Provider)

// Package provider is a generated GoMock package.
package provider

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	common "github.com/openshift/assisted-service/internal/common"
	installcfg "github.com/openshift/assisted-service/internal/installcfg"
	usage "github.com/openshift/assisted-service/internal/usage"
	models "github.com/openshift/assisted-service/models"
)

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// AddPlatformToInstallConfig mocks base method.
func (m *MockProvider) AddPlatformToInstallConfig(arg0 *installcfg.InstallerConfigBaremetal, arg1 *common.Cluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPlatformToInstallConfig", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPlatformToInstallConfig indicates an expected call of AddPlatformToInstallConfig.
func (mr *MockProviderMockRecorder) AddPlatformToInstallConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPlatformToInstallConfig", reflect.TypeOf((*MockProvider)(nil).AddPlatformToInstallConfig), arg0, arg1)
}

// AreHostsSupported mocks base method.
func (m *MockProvider) AreHostsSupported(arg0 []*models.Host) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AreHostsSupported", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AreHostsSupported indicates an expected call of AreHostsSupported.
func (mr *MockProviderMockRecorder) AreHostsSupported(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AreHostsSupported", reflect.TypeOf((*MockProvider)(nil).AreHostsSupported), arg0)
}

// CleanPlatformValuesFromDBUpdates mocks base method.
func (m *MockProvider) CleanPlatformValuesFromDBUpdates(arg0 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CleanPlatformValuesFromDBUpdates", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CleanPlatformValuesFromDBUpdates indicates an expected call of CleanPlatformValuesFromDBUpdates.
func (mr *MockProviderMockRecorder) CleanPlatformValuesFromDBUpdates(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CleanPlatformValuesFromDBUpdates", reflect.TypeOf((*MockProvider)(nil).CleanPlatformValuesFromDBUpdates), arg0)
}

// IsHostSupported mocks base method.
func (m *MockProvider) IsHostSupported(arg0 *models.Host) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsHostSupported", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsHostSupported indicates an expected call of IsHostSupported.
func (mr *MockProviderMockRecorder) IsHostSupported(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsHostSupported", reflect.TypeOf((*MockProvider)(nil).IsHostSupported), arg0)
}

// Name mocks base method.
func (m *MockProvider) Name() models.PlatformType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(models.PlatformType)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockProviderMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockProvider)(nil).Name))
}

// SetPlatformUsages mocks base method.
func (m *MockProvider) SetPlatformUsages(arg0 *models.Platform, arg1 map[string]models.Usage, arg2 usage.API) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPlatformUsages", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPlatformUsages indicates an expected call of SetPlatformUsages.
func (mr *MockProviderMockRecorder) SetPlatformUsages(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPlatformUsages", reflect.TypeOf((*MockProvider)(nil).SetPlatformUsages), arg0, arg1, arg2)
}

// SetPlatformValuesInDBUpdates mocks base method.
func (m *MockProvider) SetPlatformValuesInDBUpdates(arg0 *models.Platform, arg1 map[string]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPlatformValuesInDBUpdates", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPlatformValuesInDBUpdates indicates an expected call of SetPlatformValuesInDBUpdates.
func (mr *MockProviderMockRecorder) SetPlatformValuesInDBUpdates(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPlatformValuesInDBUpdates", reflect.TypeOf((*MockProvider)(nil).SetPlatformValuesInDBUpdates), arg0, arg1)
}
