package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	atlasConstant "github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/service_tester"
	"testing"
)

func TestTheClientReadsCollectedPlacements(t *testing.T) {
	v := service_tester.New(t)
	v.SeedInventory()
	collector_tester.AddProcess(v.Process, atlasConstant.FixtureService, true)
	collector_tester.AddPod(
		v.Kubernetes,
		"alloy-a",
		"alloy",
		"delta",
		map[string]string{"app": "alloy"},
	)
	_, e := v.Service.Collect(t.Context())
	assert.FatalOnError(t, e)
	c := client_tester.New(t, v.Service, constant.ModelContextTestToken)
	placements, f := c.ListPlacements(
		t.Context(),
		atlasConstant.FixtureVirtualMachineNode,
		"",
		"",
	)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, placements)
	assert.String(t, "relay", placements[0].Name)
	places, g := c.ListPlaces(t.Context())
	assert.FatalOnError(t, g)
	assert.Count(t, 2, places)
}

func TestTheClientFiltersSightings(t *testing.T) {
	v := service_tester.New(t)
	v.SeedInventory()
	collector_tester.AddLease(
		v.Lease,
		"delta",
		"198.51.100.8",
		"02:aa:bb:cc:dd:08",
		true,
	)
	collector_tester.AddLease(
		v.Lease,
		"juliet",
		"198.51.100.233",
		"02:aa:bb:cc:dd:33",
		false,
	)
	_, e := v.Service.Collect(t.Context())
	assert.FatalOnError(t, e)
	c := client_tester.New(t, v.Service, constant.ModelContextTestToken)
	all, f := c.ListSightings(t.Context(), "", false)
	assert.FatalOnError(t, f)
	assert.Count(t, 2, all)
	unclaimed, g := c.ListSightings(t.Context(), "", true)
	assert.FatalOnError(t, g)
	assert.Count(t, 1, unclaimed)
	assert.String(t, "198.51.100.233", unclaimed[0].Address)
}

func TestTheClientReachesTheDaemon(t *testing.T) {
	c := client_tester.New(
		t,
		service_tester.New(t).Service,
		constant.ModelContextTestToken,
	)
	assert.FatalOnError(t, c.Reachable(t.Context()))
}

func TestTheClientRefusesAWrongToken(t *testing.T) {
	c := client_tester.New(t, service_tester.New(t).Service, "wrong")
	assert.Error(t, c.Reachable(t.Context()))
	_, e := c.ListPlaces(t.Context())
	assert.Error(t, e)
}
