package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/mock_inventory_source"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	"testing"
)

func TestLabelPlacesOnTheObjectCarryingIt(t *testing.T) {
	f := mock_inventory_source.New()
	f.AddDevice(72, "kilo", collector_tester.Labels("service.lighting", "mesh"))
	v := collector_tester.CollectLabels(t, f)
	assert.Count(t, 1, v)
	assert.String(t, "lighting", v[0].Name)
	assert.String(t, "kilo", v[0].PlaceName)
	assert.String(t, "dcim.device", v[0].PlaceKind)
	assert.String(t, "netbox", v[0].Source)
}

func TestLabelPlacesVirtualMachinesToo(t *testing.T) {
	f := mock_inventory_source.New()
	f.AddMachine(
		2,
		constant.FixtureVirtualMachineNode,
		collector_tester.Labels("service.ingress", "front ingress"),
	)
	v := collector_tester.CollectLabels(t, f)
	assert.Count(t, 1, v)
	assert.String(t, "ingress", v[0].Name)
	assert.String(t, "virtualization.virtualmachine", v[0].PlaceKind)
}

func TestLabelTakesOneRowPerResponsibility(t *testing.T) {
	f := mock_inventory_source.New()
	f.AddDevice(
		72,
		"kilo",
		collector_tester.Labels(
			"service.lighting",
			"mesh",
			"service.scenes",
			"evening",
			"owner",
			"admin",
		),
	)
	v := collector_tester.CollectLabels(t, f)
	assert.Count(t, 2, v)
}

func TestLabelSkipsObjectsWithoutServiceLabels(t *testing.T) {
	f := mock_inventory_source.New()
	f.AddDevice(41, "delta", collector_tester.Labels("owner", "admin"))
	f.AddDevice(70, "lima", nil)
	assert.Count(t, 0, collector_tester.CollectLabels(t, f))
}
