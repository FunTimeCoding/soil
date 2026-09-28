package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/mock_outpost_source"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/gazetteer_tester"
	"testing"
)

func TestOutpostPlacesTheDeployedService(t *testing.T) {
	f := mock_outpost_source.New(constant.FixtureVirtualMachineNode, nil)
	f.Add(
		"goagentd.service",
		systemConstant.ServiceOriginVendor,
		"https://apt.example.test",
	)
	v, e := outpostCollector(f).Collect(
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
	f := mock_outpost_source.New(constant.FixtureVirtualMachineNode, nil)
	f.Add(
		"ssh.service",
		systemConstant.ServiceOriginSystem,
		"mirror+file:/etc/apt/mirrors/debian.list",
	)
	v, e := outpostCollector(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}

func TestOutpostRefusesASynthesizedUnit(t *testing.T) {
	f := mock_outpost_source.New(constant.FixtureVirtualMachineNode, nil)
	f.Add("getty@tty1.service", systemConstant.ServiceOriginLocal, "")
	v, e := outpostCollector(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v)
}

func TestOutpostResolvesByHardwareAddress(t *testing.T) {
	f := mock_outpost_source.New("nowhere", []string{"02:aa:bb:cc:dd:04"})
	f.Add(
		"percona-postgresql.service",
		systemConstant.ServiceOriginLocal,
		"/run/systemd/generator/percona-postgresql.service",
	)
	v, e := outpostCollector(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "echo", v[0].PlaceName)
}

func TestOutpostRefusesAnUnknownHost(t *testing.T) {
	f := mock_outpost_source.New("nowhere", nil)
	f.Add("goagentd.service", systemConstant.ServiceOriginVendor, "x")
	_, e := outpostCollector(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.Error(t, e)
}

func TestOutpostCarriesThePackageAndVersion(t *testing.T) {
	f := mock_outpost_source.New(constant.FixtureVirtualMachineNode, nil)
	f.AddPackage(
		"goagentd.service",
		systemConstant.ServiceOriginVendor,
		"https://apt.example.test",
		"goagentd",
		"0.11.154",
	)
	v, e := outpostCollector(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "goagentd", v[0].Package)
	assert.String(t, "0.11.154", v[0].Version)
}

func TestOutpostWithoutAPackageCarriesNone(t *testing.T) {
	f := mock_outpost_source.New(constant.FixtureVirtualMachineNode, nil)
	f.Add(
		"goagentd.service",
		systemConstant.ServiceOriginVendor,
		"https://apt.example.test",
	)
	v, e := outpostCollector(f).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "", v[0].Package)
	assert.String(t, "", v[0].Version)
}

func TestOutpostCollectsPastAFailedTarget(t *testing.T) {
	dead := mock_outpost_source.New(constant.FixtureVirtualMachineNode, nil)
	dead.Fail(unexpected.Format("dial %s: refused", "alfa"))
	alive := mock_outpost_source.New(constant.FixtureVirtualMachineNode, nil)
	alive.AddPackage(
		"goagentd.service",
		systemConstant.ServiceOriginVendor,
		"https://apt.example.test",
		"goagentd",
		"0.11.154",
	)
	v, e := outpostCollector(dead, alive).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)
	assert.String(t, "0.11.154", v[0].Version)
}

func TestOutpostFailsWhenEveryTargetFails(t *testing.T) {
	dead := mock_outpost_source.New(constant.FixtureVirtualMachineNode, nil)
	dead.Fail(unexpected.Format("dial %s: refused", "alfa"))
	_, e := outpostCollector(dead).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.NotNil(t, e)
}
