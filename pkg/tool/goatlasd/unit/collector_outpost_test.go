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

func TestOutpostPlacesDeployedService(t *testing.T) {
	c := collector_tester.New(t)
	f := c.OutpostClient(constant.FixtureVirtualMachineNode, nil)
	collector_tester.AddService(
		f,
		"goagentd.service",
		systemConstant.ServiceOriginVendor,
		"https://apt.example.test",
	)
	v, e := c.Outpost(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "goagentd.service", v[0].Name)
	assert.String(t, "echo", v[0].PlaceName)
	assert.String(t, "gooutpostd", v[0].Source)
}

func TestOutpostRefusesBaseSystem(t *testing.T) {
	c := collector_tester.New(t)
	f := c.OutpostClient(constant.FixtureVirtualMachineNode, nil)
	collector_tester.AddService(
		f,
		"ssh.service",
		systemConstant.ServiceOriginSystem,
		"mirror+file:/etc/apt/mirrors/debian.list",
	)
	v, e := c.Outpost(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}

func TestOutpostRefusesSynthesizedUnit(t *testing.T) {
	c := collector_tester.New(t)
	f := c.OutpostClient(constant.FixtureVirtualMachineNode, nil)
	collector_tester.AddService(
		f,
		"getty@tty1.service",
		systemConstant.ServiceOriginLocal,
		"",
	)
	v, e := c.Outpost(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}

func TestOutpostResolvesByHardwareAddress(t *testing.T) {
	c := collector_tester.New(t)
	f := c.OutpostClient("nowhere", []string{"02:aa:bb:cc:dd:04"})
	collector_tester.AddService(
		f,
		"percona-postgresql.service",
		systemConstant.ServiceOriginLocal,
		"/run/systemd/generator/percona-postgresql.service",
	)
	v, e := c.Outpost(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "echo", v[0].PlaceName)
}

func TestOutpostRefusesUnknownHost(t *testing.T) {
	c := collector_tester.New(t)
	f := c.OutpostClient("nowhere", nil)
	collector_tester.AddService(
		f,
		"goagentd.service",
		systemConstant.ServiceOriginVendor,
		"x",
	)
	_, e := c.Outpost(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.Error(t, e)
}

func TestOutpostCarriesPackageAndVersion(t *testing.T) {
	c := collector_tester.New(t)
	f := c.OutpostClient(constant.FixtureVirtualMachineNode, nil)
	collector_tester.AddPackage(
		f,
		"goagentd.service",
		systemConstant.ServiceOriginVendor,
		"https://apt.example.test",
		"goagentd",
		"0.11.154",
	)
	v, e := c.Outpost(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "goagentd", v[0].Package)
	assert.String(t, "0.11.154", v[0].Version)
}

func TestOutpostWithoutPackageCarriesNone(t *testing.T) {
	c := collector_tester.New(t)
	f := c.OutpostClient(constant.FixtureVirtualMachineNode, nil)
	collector_tester.AddService(
		f,
		"goagentd.service",
		systemConstant.ServiceOriginVendor,
		"https://apt.example.test",
	)
	v, e := c.Outpost(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "", v[0].Package)
	assert.String(t, "", v[0].Version)
}

func TestOutpostCollectsPastFailedTarget(t *testing.T) {
	c := collector_tester.New(t)
	dead := c.OutpostClient(constant.FixtureVirtualMachineNode, nil)
	dead.Failure = unexpected.Format("dial %s: refused", "alfa")
	alive := c.OutpostClient(constant.FixtureVirtualMachineNode, nil)
	collector_tester.AddPackage(
		alive,
		"goagentd.service",
		systemConstant.ServiceOriginVendor,
		"https://apt.example.test",
		"goagentd",
		"0.11.154",
	)
	v, e := c.Outpost(dead, alive).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "0.11.154", v[0].Version)
}

func TestOutpostFailsWhenEveryTargetFails(t *testing.T) {
	c := collector_tester.New(t)
	dead := c.OutpostClient(constant.FixtureVirtualMachineNode, nil)
	dead.Failure = unexpected.Format("dial %s: refused", "alfa")
	_, e := c.Outpost(dead).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.NotNil(t, e)
}
