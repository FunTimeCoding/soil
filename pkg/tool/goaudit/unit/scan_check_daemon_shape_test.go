package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/unit/scan_tester"
	"testing"
)

func TestMissingDaemonFlaggedWhenArmed(t *testing.T) {
	v := scan_tester.Daemon(constant.FixtureClosedRun)
	v.Delete("pkg/tool/gotestd/daemon/daemon.go")
	scan_tester.New(t).Concern(
		constant.MissingDaemonKey,
		scan_tester.New(t).One(v, true),
	)
}

func TestMissingDaemonQuietWhenDisarmed(t *testing.T) {
	v := scan_tester.Daemon(constant.FixtureClosedRun)
	v.Delete("pkg/tool/gotestd/daemon/daemon.go")
	scan_tester.New(t).NoConcern(
		constant.MissingDaemonKey,
		scan_tester.New(t).One(v, false),
	)
}

func TestMissingDaemonQuietWithDaemon(t *testing.T) {
	s := scan_tester.New(t).One(
		scan_tester.Daemon(constant.FixtureClosedRun),
		true,
	)
	scan_tester.New(t).NoConcern(constant.MissingDaemonKey, s)
}

func TestMissingServiceBesideDaemon(t *testing.T) {
	v := scan_tester.Daemon(constant.FixtureClosedRun)
	v.Delete("pkg/tool/gotestd/service/service.go")
	scan_tester.New(t).Concern(
		constant.MissingServiceKey,
		scan_tester.New(t).One(v, false),
	)
}

func TestServiceBesideDaemonClean(t *testing.T) {
	s := scan_tester.New(t).One(
		scan_tester.Daemon(constant.FixtureClosedRun),
		false,
	)
	scan_tester.New(t).NoConcern(constant.MissingServiceKey, s)
}

func TestMixedServiceFlagged(t *testing.T) {
	v := scan_tester.Daemon(constant.FixtureClosedRun)
	v.WriteString(
		"pkg/tool/gotestd/service/book/service.go",
		"package book\n\ntype Service struct{}\n",
	)
	scan_tester.New(t).Concern(
		constant.MixedServiceKey,
		scan_tester.New(t).One(v, false),
	)
}

func TestResultPackageBesideServiceClean(t *testing.T) {
	v := scan_tester.Daemon(constant.FixtureClosedRun)
	v.WriteString(
		"pkg/tool/gotestd/service/sync_result/result.go",
		"package sync_result\n\ntype Result struct{}\n",
	)
	scan_tester.New(t).NoConcern(
		constant.MixedServiceKey,
		scan_tester.New(t).One(v, false),
	)
}

func TestSeveralServicesWithoutCodeClean(t *testing.T) {
	v := scan_tester.Daemon(constant.FixtureClosedRun)
	v.Delete("pkg/tool/gotestd/service/service.go")
	v.WriteString(
		"pkg/tool/gotestd/service/book/service.go",
		"package book\n\ntype Service struct{}\n",
	)
	v.WriteString(
		"pkg/tool/gotestd/service/loan/service.go",
		"package loan\n\ntype Service struct{}\n",
	)
	s := scan_tester.New(t).One(v, false)
	scan_tester.New(t).NoConcern(constant.MixedServiceKey, s)
	scan_tester.New(t).NoConcern(constant.MissingServiceKey, s)
}

func TestServiceImportingMapperFlagged(t *testing.T) {
	v := scan_tester.Daemon(constant.FixtureClosedRun)
	v.Delete("pkg/tool/gotestd/service/service.go")
	v.WriteString(
		"pkg/tool/gotestd/service/book/service.go",
		constant.FixtureMapperService,
	)
	scan_tester.New(t).Concern(
		constant.ServiceImportsMapperKey,
		scan_tester.New(t).One(v, false),
	)
}

func TestDaemonImportingMapperClean(t *testing.T) {
	v := scan_tester.Daemon(constant.FixtureClosedRun)
	v.WriteString(
		"pkg/tool/gotestd/daemon/new.go",
		"package daemon\n\nimport \"gorm.io/gorm\"\n\nfunc New(m *gorm.DB) {}\n",
	)
	scan_tester.New(t).NoConcern(
		constant.ServiceImportsMapperKey,
		scan_tester.New(t).One(v, false),
	)
}

func TestRunBuildingClockFlagged(t *testing.T) {
	s := scan_tester.New(t).One(
		scan_tester.Daemon(constant.FixtureClockRun),
		false,
	)
	scan_tester.New(t).Concern(constant.RunBuildsClockKey, s)
	scan_tester.New(t).NoConcern(constant.RunBuildsClientKey, s)
}

func TestRunBuildingClockQuietBeforeDaemon(t *testing.T) {
	v := scan_tester.Daemon(constant.FixtureClockRun)
	v.Delete("pkg/tool/gotestd/daemon/daemon.go")
	scan_tester.New(t).NoConcern(
		constant.RunBuildsClockKey,
		scan_tester.New(t).One(v, true),
	)
}

func TestRunBuildingClientFlagged(t *testing.T) {
	s := scan_tester.New(t).One(
		scan_tester.Daemon(constant.FixtureClientRun),
		false,
	)
	scan_tester.New(t).Concern(constant.RunBuildsClientKey, s)
}

func TestRunBuildingOnlyDaemonClean(t *testing.T) {
	s := scan_tester.New(t).One(
		scan_tester.Daemon(constant.FixtureClosedRun),
		false,
	)
	scan_tester.New(t).NoConcern(constant.RunBuildsClientKey, s)
	scan_tester.New(t).NoConcern(constant.RunBuildsClockKey, s)
}

func TestDaemonNotClosedFlagged(t *testing.T) {
	v := scan_tester.Daemon(constant.FixtureUnclosedRun)
	v.WriteString("pkg/tool/gotestd/daemon/close.go", "package daemon\n")
	scan_tester.New(t).Concern(
		constant.DaemonNotClosedKey,
		scan_tester.New(t).One(v, false),
	)
}

func TestDaemonClosedClean(t *testing.T) {
	v := scan_tester.Daemon(constant.FixtureClosedRun)
	v.WriteString("pkg/tool/gotestd/daemon/close.go", "package daemon\n")
	scan_tester.New(t).NoConcern(
		constant.DaemonNotClosedKey,
		scan_tester.New(t).One(v, false),
	)
}

func TestSurfaceHoldingDaemonFlagged(t *testing.T) {
	v := scan_tester.Daemon(constant.FixtureClosedRun)
	v.WriteString(
		"pkg/tool/gotestd/model_context/new.go",
		"package model_context\n\nimport \"example.org/pkg/tool/gotestd/daemon\"\n",
	)
	scan_tester.New(t).Concern(
		constant.SurfaceHoldsDaemonKey,
		scan_tester.New(t).One(v, false),
	)
}

func TestSurfaceHoldingServiceClean(t *testing.T) {
	v := scan_tester.Daemon(constant.FixtureClosedRun)
	v.WriteString(
		"pkg/tool/gotestd/model_context/new.go",
		"package model_context\n\nimport \"example.org/pkg/tool/gotestd/service\"\n",
	)
	scan_tester.New(t).NoConcern(
		constant.SurfaceHoldsDaemonKey,
		scan_tester.New(t).One(v, false),
	)
}

func TestMountOutsideDaemonFlagged(t *testing.T) {
	v := scan_tester.GuardedDaemon()
	v.WriteString("pkg/tool/gotestd/mount.go", "package gotestd\n")
	scan_tester.New(t).Concern(
		constant.MissingServiceMountKey,
		scan_tester.New(t).One(v, false),
	)
}

func TestMountInDaemonClean(t *testing.T) {
	v := scan_tester.GuardedDaemon()
	v.WriteString("pkg/tool/gotestd/daemon/mount.go", "package daemon\n")
	scan_tester.New(t).NoConcern(
		constant.MissingServiceMountKey,
		scan_tester.New(t).One(v, false),
	)
}
