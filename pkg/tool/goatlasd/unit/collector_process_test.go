package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/process"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/gazetteer_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/peer_tester"
	"testing"
)

func TestProcessPlacesOnlyRunningEntries(t *testing.T) {
	f := collector_tester.NewProcess(t)
	collector_tester.AddProcess(f, "gosublimed", true)
	collector_tester.AddProcess(f, "goflightd", false)
	collector_tester.AddProcess(f, "goqueryd", true)
	v, e := process.New(f, peer_tester.NewGuard(t, f), "delta").Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, v)
	assert.String(t, "gosublimed", v[0].Name)
	assert.String(t, "goqueryd", v[1].Name)
}

func TestProcessPlacesOnTheConfiguredMachine(t *testing.T) {
	f := collector_tester.NewProcess(t)
	collector_tester.AddProcess(f, "gosublimed", true)
	v, e := process.New(
		f,
		peer_tester.NewGuard(t, f),
		constant.FixtureVirtualMachineNode,
	).Collect(context.Background(), gazetteer_tester.NewGazetteer())
	assert.FatalOnError(t, e)
	assert.String(t, "echo", v[0].PlaceName)
	assert.String(t, "virtualization.virtualmachine", v[0].PlaceKind)
	assert.String(t, "goprocessd", v[0].Source)
}

func TestProcessRefusesAnUnknownMachine(t *testing.T) {
	f := collector_tester.NewProcess(t)
	collector_tester.AddProcess(f, "gosublimed", true)
	_, e := process.New(f, peer_tester.NewGuard(t, f), "nowhere").Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.Error(t, e)
}

func TestProcessCarriesNoScope(t *testing.T) {
	f := collector_tester.NewProcess(t)
	collector_tester.AddProcess(f, "gosublimed", true)
	v, e := process.New(f, peer_tester.NewGuard(t, f), "delta").Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.String(t, "", v[0].Scope)
}
