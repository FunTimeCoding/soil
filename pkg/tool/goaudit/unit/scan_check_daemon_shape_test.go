package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"testing"
)

func TestMissingDaemonFlaggedWhenArmed(t *testing.T) {
	v := daemonFixture(constant.FixtureClosedRun)
	v.Delete("pkg/tool/gotestd/daemon/daemon.go")
	assertConcern(t, scanOne(t, v, true), constant.MissingDaemonKey)
}

func TestMissingDaemonQuietWhenDisarmed(t *testing.T) {
	v := daemonFixture(constant.FixtureClosedRun)
	v.Delete("pkg/tool/gotestd/daemon/daemon.go")
	assertNoConcern(t, scanOne(t, v, false), constant.MissingDaemonKey)
}

func TestMissingDaemonQuietWithDaemon(t *testing.T) {
	s := scanOne(t, daemonFixture(constant.FixtureClosedRun), true)
	assertNoConcern(t, s, constant.MissingDaemonKey)
}

func TestMissingServiceBesideDaemon(t *testing.T) {
	v := daemonFixture(constant.FixtureClosedRun)
	v.Delete("pkg/tool/gotestd/service/service.go")
	assertConcern(t, scanOne(t, v, false), constant.MissingServiceKey)
}

func TestServiceBesideDaemonClean(t *testing.T) {
	s := scanOne(t, daemonFixture(constant.FixtureClosedRun), false)
	assertNoConcern(t, s, constant.MissingServiceKey)
}

func TestMixedServiceFlagged(t *testing.T) {
	v := daemonFixture(constant.FixtureClosedRun)
	v.WriteString(
		"pkg/tool/gotestd/service/book/service.go",
		"package book\n\ntype Service struct{}\n",
	)
	assertConcern(t, scanOne(t, v, false), constant.MixedServiceKey)
}

func TestResultPackageBesideServiceClean(t *testing.T) {
	v := daemonFixture(constant.FixtureClosedRun)
	v.WriteString(
		"pkg/tool/gotestd/service/sync_result/result.go",
		"package sync_result\n\ntype Result struct{}\n",
	)
	assertNoConcern(t, scanOne(t, v, false), constant.MixedServiceKey)
}

func TestSeveralServicesWithoutCodeClean(t *testing.T) {
	v := daemonFixture(constant.FixtureClosedRun)
	v.Delete("pkg/tool/gotestd/service/service.go")
	v.WriteString(
		"pkg/tool/gotestd/service/book/service.go",
		"package book\n\ntype Service struct{}\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/service/loan/service.go",
		"package loan\n\ntype Service struct{}\n",
	)
	s := scanOne(t, v, false)
	assertNoConcern(t, s, constant.MixedServiceKey)
	assertNoConcern(t, s, constant.MissingServiceKey)
}

func TestServiceImportingMapperFlagged(t *testing.T) {
	v := daemonFixture(constant.FixtureClosedRun)
	v.Delete("pkg/tool/gotestd/service/service.go")
	v.WriteString(
		"pkg/tool/gotestd/service/book/service.go",
		constant.FixtureMapperService,
	)
	assertConcern(t, scanOne(t, v, false), constant.ServiceImportsMapperKey)
}

func TestDaemonImportingMapperClean(t *testing.T) {
	v := daemonFixture(constant.FixtureClosedRun)
	v.WriteString(
		"pkg/tool/gotestd/daemon/new.go",
		"package daemon\n\nimport \"gorm.io/gorm\"\n\nfunc New(m *gorm.DB) {}\n",
	)
	assertNoConcern(t, scanOne(t, v, false), constant.ServiceImportsMapperKey)
}

func TestRunBuildingClockFlagged(t *testing.T) {
	s := scanOne(t, daemonFixture(constant.FixtureClockRun), false)
	assertConcern(t, s, constant.RunBuildsClockKey)
	assertNoConcern(t, s, constant.RunBuildsClientKey)
}

func TestRunBuildingClockQuietBeforeDaemon(t *testing.T) {
	v := daemonFixture(constant.FixtureClockRun)
	v.Delete("pkg/tool/gotestd/daemon/daemon.go")
	assertNoConcern(t, scanOne(t, v, true), constant.RunBuildsClockKey)
}

func TestRunBuildingClientFlagged(t *testing.T) {
	s := scanOne(t, daemonFixture(constant.FixtureClientRun), false)
	assertConcern(t, s, constant.RunBuildsClientKey)
}

func TestRunBuildingOnlyDaemonClean(t *testing.T) {
	s := scanOne(t, daemonFixture(constant.FixtureClosedRun), false)
	assertNoConcern(t, s, constant.RunBuildsClientKey)
	assertNoConcern(t, s, constant.RunBuildsClockKey)
}

func TestDaemonNotClosedFlagged(t *testing.T) {
	v := daemonFixture(constant.FixtureUnclosedRun)
	v.WriteString("pkg/tool/gotestd/daemon/close.go", "package daemon\n")
	assertConcern(t, scanOne(t, v, false), constant.DaemonNotClosedKey)
}

func TestDaemonClosedClean(t *testing.T) {
	v := daemonFixture(constant.FixtureClosedRun)
	v.WriteString("pkg/tool/gotestd/daemon/close.go", "package daemon\n")
	assertNoConcern(t, scanOne(t, v, false), constant.DaemonNotClosedKey)
}

func TestSurfaceHoldingDaemonFlagged(t *testing.T) {
	v := daemonFixture(constant.FixtureClosedRun)
	v.WriteString(
		"pkg/tool/gotestd/model_context/new.go",
		"package model_context\n\nimport \"example.org/pkg/tool/gotestd/daemon\"\n",
	)
	assertConcern(t, scanOne(t, v, false), constant.SurfaceHoldsDaemonKey)
}

func TestSurfaceHoldingServiceClean(t *testing.T) {
	v := daemonFixture(constant.FixtureClosedRun)
	v.WriteString(
		"pkg/tool/gotestd/model_context/new.go",
		"package model_context\n\nimport \"example.org/pkg/tool/gotestd/service\"\n",
	)
	assertNoConcern(t, scanOne(t, v, false), constant.SurfaceHoldsDaemonKey)
}

func TestMountOutsideDaemonFlagged(t *testing.T) {
	v := guardedDaemon()
	v.WriteString("pkg/tool/gotestd/mount.go", "package gotestd\n")
	assertConcern(t, scanOne(t, v, false), constant.MissingServiceMountKey)
}

func TestMountInDaemonClean(t *testing.T) {
	v := guardedDaemon()
	v.WriteString("pkg/tool/gotestd/daemon/mount.go", "package daemon\n")
	assertNoConcern(t, scanOne(t, v, false), constant.MissingServiceMountKey)
}
