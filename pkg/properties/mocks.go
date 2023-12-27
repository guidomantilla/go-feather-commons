// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/properties/types.go
//
// Generated by this command:
//
//	mockgen -source ../pkg/properties/types.go -destination ../pkg/properties/mocks.go -package=properties
//
// Package properties is a generated GoMock package.
package properties

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockProperties is a mock of Properties interface.
type MockProperties struct {
	ctrl     *gomock.Controller
	recorder *MockPropertiesMockRecorder
}

// MockPropertiesMockRecorder is the mock recorder for MockProperties.
type MockPropertiesMockRecorder struct {
	mock *MockProperties
}

// NewMockProperties creates a new mock instance.
func NewMockProperties(ctrl *gomock.Controller) *MockProperties {
	mock := &MockProperties{ctrl: ctrl}
	mock.recorder = &MockPropertiesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProperties) EXPECT() *MockPropertiesMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockProperties) Add(property, value string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", property, value)
}

// Add indicates an expected call of Add.
func (mr *MockPropertiesMockRecorder) Add(property, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockProperties)(nil).Add), property, value)
}

// AsMap mocks base method.
func (m *MockProperties) AsMap() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsMap")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// AsMap indicates an expected call of AsMap.
func (mr *MockPropertiesMockRecorder) AsMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsMap", reflect.TypeOf((*MockProperties)(nil).AsMap))
}

// Get mocks base method.
func (m *MockProperties) Get(property string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", property)
	ret0, _ := ret[0].(string)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockPropertiesMockRecorder) Get(property any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProperties)(nil).Get), property)
}

// MockPropertySource is a mock of PropertySource interface.
type MockPropertySource struct {
	ctrl     *gomock.Controller
	recorder *MockPropertySourceMockRecorder
}

// MockPropertySourceMockRecorder is the mock recorder for MockPropertySource.
type MockPropertySourceMockRecorder struct {
	mock *MockPropertySource
}

// NewMockPropertySource creates a new mock instance.
func NewMockPropertySource(ctrl *gomock.Controller) *MockPropertySource {
	mock := &MockPropertySource{ctrl: ctrl}
	mock.recorder = &MockPropertySourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPropertySource) EXPECT() *MockPropertySourceMockRecorder {
	return m.recorder
}

// AsMap mocks base method.
func (m *MockPropertySource) AsMap() map[string]any {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsMap")
	ret0, _ := ret[0].(map[string]any)
	return ret0
}

// AsMap indicates an expected call of AsMap.
func (mr *MockPropertySourceMockRecorder) AsMap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsMap", reflect.TypeOf((*MockPropertySource)(nil).AsMap))
}

// Get mocks base method.
func (m *MockPropertySource) Get(property string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", property)
	ret0, _ := ret[0].(string)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockPropertySourceMockRecorder) Get(property any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPropertySource)(nil).Get), property)
}
