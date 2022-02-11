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
// Source: github.com/containerd/containerd (interfaces: Container)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	containerd "github.com/containerd/containerd"
	cio "github.com/containerd/containerd/cio"
	containers "github.com/containerd/containerd/containers"
	types "github.com/gogo/protobuf/types"
	gomock "github.com/golang/mock/gomock"
	specs_go "github.com/opencontainers/runtime-spec/specs-go"
)

// MockContainer is a mock of Container interface.
type MockContainer struct {
	ctrl     *gomock.Controller
	recorder *MockContainerMockRecorder
}

// MockContainerMockRecorder is the mock recorder for MockContainer.
type MockContainerMockRecorder struct {
	mock *MockContainer
}

// NewMockContainer creates a new mock instance.
func NewMockContainer(ctrl *gomock.Controller) *MockContainer {
	mock := &MockContainer{ctrl: ctrl}
	mock.recorder = &MockContainerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContainer) EXPECT() *MockContainerMockRecorder {
	return m.recorder
}

// Checkpoint mocks base method.
func (m *MockContainer) Checkpoint(arg0 context.Context, arg1 string, arg2 ...containerd.CheckpointOpts) (containerd.Image, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Checkpoint", varargs...)
	ret0, _ := ret[0].(containerd.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Checkpoint indicates an expected call of Checkpoint.
func (mr *MockContainerMockRecorder) Checkpoint(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checkpoint", reflect.TypeOf((*MockContainer)(nil).Checkpoint), varargs...)
}

// Delete mocks base method.
func (m *MockContainer) Delete(arg0 context.Context, arg1 ...containerd.DeleteOpts) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockContainerMockRecorder) Delete(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockContainer)(nil).Delete), varargs...)
}

// Extensions mocks base method.
func (m *MockContainer) Extensions(arg0 context.Context) (map[string]types.Any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Extensions", arg0)
	ret0, _ := ret[0].(map[string]types.Any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Extensions indicates an expected call of Extensions.
func (mr *MockContainerMockRecorder) Extensions(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Extensions", reflect.TypeOf((*MockContainer)(nil).Extensions), arg0)
}

// ID mocks base method.
func (m *MockContainer) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockContainerMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockContainer)(nil).ID))
}

// Image mocks base method.
func (m *MockContainer) Image(arg0 context.Context) (containerd.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Image", arg0)
	ret0, _ := ret[0].(containerd.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Image indicates an expected call of Image.
func (mr *MockContainerMockRecorder) Image(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Image", reflect.TypeOf((*MockContainer)(nil).Image), arg0)
}

// Info mocks base method.
func (m *MockContainer) Info(arg0 context.Context, arg1 ...containerd.InfoOpts) (containers.Container, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Info", varargs...)
	ret0, _ := ret[0].(containers.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Info indicates an expected call of Info.
func (mr *MockContainerMockRecorder) Info(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockContainer)(nil).Info), varargs...)
}

// Labels mocks base method.
func (m *MockContainer) Labels(arg0 context.Context) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Labels indicates an expected call of Labels.
func (mr *MockContainerMockRecorder) Labels(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockContainer)(nil).Labels), arg0)
}

// NewTask mocks base method.
func (m *MockContainer) NewTask(arg0 context.Context, arg1 cio.Creator, arg2 ...containerd.NewTaskOpts) (containerd.Task, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewTask", varargs...)
	ret0, _ := ret[0].(containerd.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewTask indicates an expected call of NewTask.
func (mr *MockContainerMockRecorder) NewTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTask", reflect.TypeOf((*MockContainer)(nil).NewTask), varargs...)
}

// SetLabels mocks base method.
func (m *MockContainer) SetLabels(arg0 context.Context, arg1 map[string]string) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLabels", arg0, arg1)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetLabels indicates an expected call of SetLabels.
func (mr *MockContainerMockRecorder) SetLabels(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLabels", reflect.TypeOf((*MockContainer)(nil).SetLabels), arg0, arg1)
}

// Spec mocks base method.
func (m *MockContainer) Spec(arg0 context.Context) (*specs_go.Spec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Spec", arg0)
	ret0, _ := ret[0].(*specs_go.Spec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Spec indicates an expected call of Spec.
func (mr *MockContainerMockRecorder) Spec(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Spec", reflect.TypeOf((*MockContainer)(nil).Spec), arg0)
}

// Task mocks base method.
func (m *MockContainer) Task(arg0 context.Context, arg1 cio.Attach) (containerd.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Task", arg0, arg1)
	ret0, _ := ret[0].(containerd.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Task indicates an expected call of Task.
func (mr *MockContainerMockRecorder) Task(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Task", reflect.TypeOf((*MockContainer)(nil).Task), arg0, arg1)
}

// Update mocks base method.
func (m *MockContainer) Update(arg0 context.Context, arg1 ...containerd.UpdateContainerOpts) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockContainerMockRecorder) Update(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockContainer)(nil).Update), varargs...)
}
