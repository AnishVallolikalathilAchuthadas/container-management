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
// Source: containerm/ctr/ctrd_log_mgr.go

// Package mocks is a generated GoMock package.
package ctrd

import (
	reflect "reflect"

	types "github.com/eclipse-kanto/container-management/containerm/containers/types"
	logger "github.com/eclipse-kanto/container-management/containerm/logger"
	gomock "github.com/golang/mock/gomock"
)

// MockcontainerLogsManager is a mock of containerLogsManager interface.
type MockcontainerLogsManager struct {
	ctrl     *gomock.Controller
	recorder *MockcontainerLogsManagerMockRecorder
}

// MockcontainerLogsManagerMockRecorder is the mock recorder for MockcontainerLogsManager.
type MockcontainerLogsManagerMockRecorder struct {
	mock *MockcontainerLogsManager
}

// NewMockcontainerLogsManager creates a new mock instance.
func NewMockcontainerLogsManager(ctrl *gomock.Controller) *MockcontainerLogsManager {
	mock := &MockcontainerLogsManager{ctrl: ctrl}
	mock.recorder = &MockcontainerLogsManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcontainerLogsManager) EXPECT() *MockcontainerLogsManagerMockRecorder {
	return m.recorder
}

// GetLogDriver mocks base method.
func (m *MockcontainerLogsManager) GetLogDriver(c *types.Container) (logger.LogDriver, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogDriver", c)
	ret0, _ := ret[0].(logger.LogDriver)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogDriver indicates an expected call of GetLogDriver.
func (mr *MockcontainerLogsManagerMockRecorder) GetLogDriver(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogDriver", reflect.TypeOf((*MockcontainerLogsManager)(nil).GetLogDriver), c)
}
