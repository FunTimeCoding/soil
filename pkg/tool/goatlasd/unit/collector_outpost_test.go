package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/gazetteer_tester"
	"testing"
)

func TestOutpostPlacesTheDeployedService(t *testing.T) {
	f := collector_tester.NewOutpost(t, constant.FixtureVirtualMachineNode, nil)
	collector_tester.AddService(
		f,
		"goagentd.service",
		systemConstant.ServiceOriginVendor,
		"https://apt.example.test",
	)
	v, e := outpostCollector(t, f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "goagentd.service", v[0].Name)
	assert.String(t, "echo", v[0].PlaceName)
	assert.String(t, "gooutpostd", v[0].Source)
}

func TestOutpostRefusesTheBaseSystem(t *testing.T) {
	f := collector_tester.NewOutpost(t, constant.FixtureVirtualMachineNode, nil)
	collector_tester.AddService(
		f,
		"ssh.service",
		systemConstant.ServiceOriginSystem,
		"mirror+file:/etc/apt/mirrors/debian.list",
	)
	v, e := outpostCollector(t, f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}

func TestOutpostRefusesASynthesizedUnit(t *testing.T) {
	f := collector_tester.NewOutpost(t, constant.FixtureVirtualMachineNode, nil)
	collector_tester.AddService(
		f,
		"getty@tty1.service",
		systemConstant.ServiceOriginLocal,
		"",
	)
	v, e := outpostCollector(t, f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}

func TestOutpostResolvesByHardwareAddress(t *testing.T) {
	f := collector_tester.NewOutpost(
		t,
		"nowhere",
		[]string{"02:aa:bb:cc:dd:04"},
	)
	collector_tester.AddService(
		f,
		"percona-postgresql.service",
		systemConstant.ServiceOriginLocal,
		"/run/systemd/generator/percona-postgresql.service",
	)
	v, e := outpostCollector(t, f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "echo", v[0].PlaceName)
}

func TestOutpostRefusesAnUnknownHost(t *testing.T) {
	f := collector_tester.NewOutpost(t, "nowhere", nil)
	collector_tester.AddService(
		f,
		"goagentd.service",
		systemConstant.ServiceOriginVendor,
		"x",
	)
	_, e := outpostCollector(t, f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.Error(t, e)
}

func TestOutpostCarriesThePackageAndVersion(t *testing.T) {
	f := collector_tester.NewOutpost(t, constant.FixtureVirtualMachineNode, nil)
	collector_tester.AddPackage(
		f,
		"goagentd.service",
		systemConstant.ServiceOriginVendor,
		"https://apt.example.test",
		"goagentd",
		"0.11.154",
	)
	v, e := outpostCollector(t, f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "goagentd", v[0].Package)
	assert.String(t, "0.11.154", v[0].Version)
}

func TestOutpostWithoutAPackageCarriesNone(t *testing.T) {
	f := collector_tester.NewOutpost(t, constant.FixtureVirtualMachineNode, nil)
	collector_tester.AddService(
		f,
		"goagentd.service",
		systemConstant.ServiceOriginVendor,
		"https://apt.example.test",
	)
	v, e := outpostCollector(t, f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "", v[0].Package)
	assert.String(t, "", v[0].Version)
}

func TestOutpostCollectsPastAFailedTarget(t *testing.T) {
	dead := collector_tester.NewOutpost(
		t,
		constant.FixtureVirtualMachineNode,
		nil,
	)
	dead.Failure = unexpected.Format("dial %s: refused", "alfa")
	alive := collector_tester.NewOutpost(
		t,
		constant.FixtureVirtualMachineNode,
		nil,
	)
	collector_tester.AddPackage(
		alive,
		"goagentd.service",
		systemConstant.ServiceOriginVendor,
		"https://apt.example.test",
		"goagentd",
		"0.11.154",
	)
	v, e := outpostCollector(t, dead, alive).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "0.11.154", v[0].Version)
}

func TestOutpostFailsWhenEveryTargetFails(t *testing.T) {
	dead := collector_tester.NewOutpost(
		t,
		constant.FixtureVirtualMachineNode,
		nil,
	)
	dead.Failure = unexpected.Format("dial %s: refused", "alfa")
	_, e := outpostCollector(t, dead).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.NotNil(t, e)
}
