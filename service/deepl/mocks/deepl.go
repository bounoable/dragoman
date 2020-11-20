// Code generated by MockGen. DO NOT EDIT.
// Source: deepl.go

// Package mock_deepl is a generated GoMock package.
package mock_deepl

import (
	context "context"
	deepl "github.com/bounoable/deepl"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Translate mocks base method
func (m *MockClient) Translate(ctx context.Context, text string, targetLang deepl.Language, opts ...deepl.TranslateOption) (string, deepl.Language, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, text, targetLang}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Translate", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(deepl.Language)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Translate indicates an expected call of Translate
func (mr *MockClientMockRecorder) Translate(ctx, text, targetLang interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, text, targetLang}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Translate", reflect.TypeOf((*MockClient)(nil).Translate), varargs...)
}
