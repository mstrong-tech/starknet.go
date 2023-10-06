// Code generated by MockGen. DO NOT EDIT.
// Source: account.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	felt "github.com/NethermindEth/juno/core/felt"
	rpc "github.com/NethermindEth/starknet.go/rpc"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountInterface is a mock of AccountInterface interface.
type MockAccountInterface struct {
	ctrl     *gomock.Controller
	recorder *MockAccountInterfaceMockRecorder
}

// MockAccountInterfaceMockRecorder is the mock recorder for MockAccountInterface.
type MockAccountInterfaceMockRecorder struct {
	mock *MockAccountInterface
}

// NewMockAccountInterface creates a new mock instance.
func NewMockAccountInterface(ctrl *gomock.Controller) *MockAccountInterface {
	mock := &MockAccountInterface{ctrl: ctrl}
	mock.recorder = &MockAccountInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountInterface) EXPECT() *MockAccountInterfaceMockRecorder {
	return m.recorder
}

// Call mocks base method.
func (m *MockAccountInterface) Call(ctx context.Context, call rpc.FunctionCall) ([]*felt.Felt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", ctx, call)
	ret0, _ := ret[0].([]*felt.Felt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call indicates an expected call of Call.
func (mr *MockAccountInterfaceMockRecorder) Call(ctx, call interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockAccountInterface)(nil).Call), ctx, call)
}

// TransactionHash mocks base method.
func (m *MockAccountInterface) TransactionHash(calls rpc.FunctionCall, txDetails rpc.TxDetails) (*felt.Felt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionHash", calls, txDetails)
	ret0, _ := ret[0].(*felt.Felt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionHash indicates an expected call of TransactionHash.
func (mr *MockAccountInterfaceMockRecorder) TransactionHash(calls, txDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionHash", reflect.TypeOf((*MockAccountInterface)(nil).TransactionHash), calls, txDetails)
}