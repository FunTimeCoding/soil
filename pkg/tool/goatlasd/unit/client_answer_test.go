package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	"testing"
)

func TestClientReadsCollectedPlacements(t *testing.T) {
	s := base.New(t)
	s.SeedInventory()
	collector_tester.AddProcess(s.Process, constant.FixtureService, true)
	collector_tester.AddPod(
		s.Kubernetes,
		"alloy-a",
		"alloy",
		"delta",
		map[string]string{"app": "alloy"},
	)
	_, e := s.Service.Collect(t.Context())
	assert.FatalOnError(t, e)
	c := client_tester.New(t, s, generativeConstant.ModelContextTestToken)
	placements, f := c.ListPlacements(
		t.Context(),
		constant.FixtureVirtualMachineNode,
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

func TestClientFiltersSightings(t *testing.T) {
	s := base.New(t)
	s.SeedInventory()
	collector_tester.AddLease(
		s.Lease,
		"delta",
		"198.51.100.8",
		"02:aa:bb:cc:dd:08",
		true,
	)
	collector_tester.AddLease(
		s.Lease,
		"juliet",
		"198.51.100.233",
		"02:aa:bb:cc:dd:33",
		false,
	)
	_, e := s.Service.Collect(t.Context())
	assert.FatalOnError(t, e)
	c := client_tester.New(t, s, generativeConstant.ModelContextTestToken)
	all, f := c.ListSightings(t.Context(), "", false)
	assert.FatalOnError(t, f)
	assert.Count(t, 2, all)
	unclaimed, g := c.ListSightings(t.Context(), "", true)
	assert.FatalOnError(t, g)
	assert.Count(t, 1, unclaimed)
	assert.String(t, "198.51.100.233", unclaimed[0].Address)
}

func TestClientReachesDaemon(t *testing.T) {
	c := client_tester.New(
		t,
		base.New(t),
		generativeConstant.ModelContextTestToken,
	)
	assert.FatalOnError(t, c.Reachable(t.Context()))
}

func TestClientRefusesWrongToken(t *testing.T) {
	c := client_tester.New(t, base.New(t), "wrong")
	assert.Error(t, c.Reachable(t.Context()))
	_, e := c.ListPlaces(t.Context())
	assert.Error(t, e)
}
