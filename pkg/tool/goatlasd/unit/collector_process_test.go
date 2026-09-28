package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/collector/process"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/mock_process_source"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/gazetteer_tester"
	"testing"
)

func TestProcessPlacesOnlyRunningEntries(t *testing.T) {
	f := mock_process_source.New()
	f.Add("gosublimed", true)
	f.Add("goflightd", false)
	f.Add("goqueryd", true)
	v, e := process.New(f, "delta").Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, v)
	assert.String(t, "gosublimed", v[0].Name)
	assert.String(t, "goqueryd", v[1].Name)
}

func TestProcessPlacesOnTheConfiguredMachine(t *testing.T) {
	f := mock_process_source.New()
	f.Add("gosublimed", true)
	v, e := process.New(f, constant.FixtureVirtualMachineNode).Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.String(t, "echo", v[0].PlaceName)
	assert.String(t, "virtualization.virtualmachine", v[0].PlaceKind)
	assert.String(t, "goprocessd", v[0].Source)
}

func TestProcessRefusesAnUnknownMachine(t *testing.T) {
	f := mock_process_source.New()
	f.Add("gosublimed", true)
	_, e := process.New(f, "nowhere").Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.Error(t, e)
}

func TestProcessCarriesNoScope(t *testing.T) {
	f := mock_process_source.New()
	f.Add("gosublimed", true)
	v, e := process.New(f, "delta").Collect(
		context.Background(),
		gazetteer_tester.NewGazetteer(),
	)
	assert.FatalOnError(t, e)
	assert.String(t, "", v[0].Scope)
}
