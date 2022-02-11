// Copyright (c) 2021 Contributors to the Eclipse Foundation
//
// See the NOTICE file(s) distributed with this work for additional
// information regarding copyright ownership.
//
// This program and the accompanying materials are made available under the
// terms of the Eclipse Public License 2.0 which is available at
// http://www.eclipse.org/legal/epl-2.0
//
// SPDX-License-Identifier: EPL-2.0

// Code generated by MockGen. DO NOT EDIT.
// Source: ./containerm/network/net_mgr_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	types "github.com/eclipse-kanto/container-management/containerm/containers/types"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockNetworkManager is a mock of NetworkManager interface
type MockNetworkManager struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkManagerMockRecorder
}

// MockNetworkManagerMockRecorder is the mock recorder for MockNetworkManager
type MockNetworkManagerMockRecorder struct {
	mock *MockNetworkManager
}

// NewMockNetworkManager creates a new mock instance
func NewMockNetworkManager(ctrl *gomock.Controller) *MockNetworkManager {
	mock := &MockNetworkManager{ctrl: ctrl}
	mock.recorder = &MockNetworkManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkManager) EXPECT() *MockNetworkManagerMockRecorder {
	return m.recorder
}

// Manage mocks base method
func (m *MockNetworkManager) Manage(ctx context.Context, container *types.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Manage", ctx, container)
	ret0, _ := ret[0].(error)
	return ret0
}

// Manage indicates an expected call of Manage
func (mr *MockNetworkManagerMockRecorder) Manage(ctx, container interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Manage", reflect.TypeOf((*MockNetworkManager)(nil).Manage), ctx, container)
}

// Connect mocks base method
func (m *MockNetworkManager) Connect(ctx context.Context, containers *types.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", ctx, containers)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockNetworkManagerMockRecorder) Connect(ctx, containers interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockNetworkManager)(nil).Connect), ctx, containers)
}

// Disconnect mocks base method
func (m *MockNetworkManager) Disconnect(ctx context.Context, container *types.Container, force bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Disconnect", ctx, container, force)
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect
func (mr *MockNetworkManagerMockRecorder) Disconnect(ctx, container, force interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockNetworkManager)(nil).Disconnect), ctx, container, force)
}

// ReleaseNetworkResources mocks base method
func (m *MockNetworkManager) ReleaseNetworkResources(ctx context.Context, container *types.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseNetworkResources", ctx, container)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReleaseNetworkResources indicates an expected call of ReleaseNetworkResources
func (mr *MockNetworkManagerMockRecorder) ReleaseNetworkResources(ctx, container interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseNetworkResources", reflect.TypeOf((*MockNetworkManager)(nil).ReleaseNetworkResources), ctx, container)
}

// Dispose mocks base method
func (m *MockNetworkManager) Dispose(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dispose", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Dispose indicates an expected call of Dispose
func (mr *MockNetworkManagerMockRecorder) Dispose(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dispose", reflect.TypeOf((*MockNetworkManager)(nil).Dispose), ctx)
}

// Restore mocks base method
func (m *MockNetworkManager) Restore(ctx context.Context, container []*types.Container) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", ctx, container)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore
func (mr *MockNetworkManagerMockRecorder) Restore(ctx, container interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockNetworkManager)(nil).Restore), ctx, container)
}

// Initialize mocks base method
func (m *MockNetworkManager) Initialize(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize
func (mr *MockNetworkManagerMockRecorder) Initialize(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockNetworkManager)(nil).Initialize), ctx)
}
