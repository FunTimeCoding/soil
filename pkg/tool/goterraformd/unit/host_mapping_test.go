package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/types/host_event"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/host_mapping"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/types/plan"
	"testing"
)

func TestHostEventsFromPlan(t *testing.T) {
	p := fixturePlan()
	assert.Any(
		t,
		[]*host_event.Event{
			host_event.New("alfa", "update"),
			host_event.New("bravo", "update"),
			host_event.New("foxtrot", "delete"),
			host_event.New("golf", "replace"),
			host_event.New("hotel", "update"),
			host_event.New("lima", "update"),
		},
		host_mapping.New(p, constant.HostAttribute).Events(p.ResourceChanges),
	)
}

func TestPlanAddressesSkipReadsAndNoOps(t *testing.T) {
	assert.Strings(
		t,
		[]string{
			"module.storage.example_disk.charlie",
			"example_network.delta",
			"module.compute.proxmox_download_file.image",
			"example_role.echo",
			"module.compute.proxmox_virtual_environment_vm.golf",
			"module.compute.example_disk.juliett",
			"module.compute.example_disk.kilo",
			"module.compute.proxmox_virtual_environment_vm.foxtrot",
		},
		host_mapping.Addresses(fixturePlan()),
	)
}

func TestEmptyPlanHasNoHosts(t *testing.T) {
	p := plan.New()
	assert.Any(
		t,
		[]*host_event.Event{},
		host_mapping.New(p, constant.HostAttribute).Events(p.ResourceChanges),
	)
}
