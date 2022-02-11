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

package sysinfo

import (
	"testing"

	"github.com/eclipse-kanto/container-management/containerm/pkg/testutil"
	"github.com/eclipse-kanto/container-management/containerm/registry"
	"github.com/eclipse-kanto/container-management/containerm/version"
)

func TestRegistryInit(t *testing.T) {
	arg := &registry.ServiceRegistryContext{}

	got, _ := registryInit(arg)
	expectedSysInfoMgr := got.(*systemInfoMgr)
	testutil.AssertEqual(t, expectedSysInfoMgr.mgrVersionInfo.ProjectVersion, version.ProjectVersion)
	testutil.AssertEqual(t, expectedSysInfoMgr.mgrVersionInfo.BuildTime, version.BuildTime)
	testutil.AssertEqual(t, expectedSysInfoMgr.mgrVersionInfo.APIVersion, version.APIVersion)
	testutil.AssertEqual(t, expectedSysInfoMgr.mgrVersionInfo.GitCommit, version.GitCommit)

}
