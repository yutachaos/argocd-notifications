// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/argoproj-labs/argocd-notifications/triggers (interfaces: Trigger)

// Package mocks is a generated GoMock package.
package mocks

import (
	notifiers "github.com/argoproj-labs/argocd-notifications/notifiers"
	gomock "github.com/golang/mock/gomock"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	reflect "reflect"
)

// MockTrigger is a mock of Trigger interface
type MockTrigger struct {
	ctrl     *gomock.Controller
	recorder *MockTriggerMockRecorder
}

// MockTriggerMockRecorder is the mock recorder for MockTrigger
type MockTriggerMockRecorder struct {
	mock *MockTrigger
}

// NewMockTrigger creates a new mock instance
func NewMockTrigger(ctrl *gomock.Controller) *MockTrigger {
	mock := &MockTrigger{ctrl: ctrl}
	mock.recorder = &MockTriggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrigger) EXPECT() *MockTriggerMockRecorder {
	return m.recorder
}

// FormatNotification mocks base method
func (m *MockTrigger) FormatNotification(arg0 *unstructured.Unstructured, arg1 map[string]string) (*notifiers.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FormatNotification", arg0, arg1)
	ret0, _ := ret[0].(*notifiers.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FormatNotification indicates an expected call of FormatNotification
func (mr *MockTriggerMockRecorder) FormatNotification(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FormatNotification", reflect.TypeOf((*MockTrigger)(nil).FormatNotification), arg0, arg1)
}

// Triggered mocks base method
func (m *MockTrigger) Triggered(arg0 *unstructured.Unstructured) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Triggered", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Triggered indicates an expected call of Triggered
func (mr *MockTriggerMockRecorder) Triggered(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Triggered", reflect.TypeOf((*MockTrigger)(nil).Triggered), arg0)
}
