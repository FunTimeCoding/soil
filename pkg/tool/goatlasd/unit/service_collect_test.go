package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/collector_tester"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/service_tester"
	"testing"
)

func TestCollectKeepsSourcesAfterAFailingOne(t *testing.T) {
	v := service_tester.New(t)
	v.SeedInventory()
	v.Kubernetes.Failure = errors.New("pod list refused")
	collector_tester.AddProcess(v.Process, constant.FixtureService, true)
	outcomes, e := v.Service.Collect(t.Context())
	assert.FatalOnError(t, e)
	assert.Count(t, 4, outcomes)
	assert.Error(t, outcomes[0].Error)
	assert.FatalOnError(t, outcomes[2].Error)
	assert.Integer(t, 1, outcomes[2].Count)
	placed, f := v.Service.MatchingPlacements("", constant.SourceProcess, "")
	assert.FatalOnError(t, f)
	assert.Count(t, 1, placed)
	assert.String(t, "echo", placed[0].PlaceName)
}

func TestCollectKeepsTheFailingSourcePlacements(t *testing.T) {
	v := service_tester.New(t)
	v.SeedInventory()
	collector_tester.AddProcess(v.Process, constant.FixtureService, true)
	_, e := v.Service.Collect(t.Context())
	assert.FatalOnError(t, e)
	v.Process.Failure = errors.New("process list refused")
	outcomes, f := v.Service.Collect(t.Context())
	assert.FatalOnError(t, f)
	assert.Error(t, outcomes[2].Error)
	placed, g := v.Service.MatchingPlacements("", constant.SourceProcess, "")
	assert.FatalOnError(t, g)
	assert.Count(t, 1, placed)
}

func TestCollectFailsWithoutTheInventory(t *testing.T) {
	v := service_tester.New(t)
	v.Netbox.Failure = errors.New("device list refused")
	_, e := v.Service.Collect(t.Context())
	assert.Error(t, e)
}

func TestCollectSightsLeases(t *testing.T) {
	v := service_tester.New(t)
	v.SeedInventory()
	collector_tester.AddLease(
		v.Lease,
		"delta",
		"198.51.100.8",
		"02:aa:bb:cc:dd:08",
		true,
	)
	outcomes, e := v.Service.Collect(t.Context())
	assert.FatalOnError(t, e)
	last := outcomes[len(outcomes)-1]
	assert.String(t, "gopnsensed", last.Source)
	assert.True(t, last.Sightings)
	assert.Integer(t, 1, last.Count)
	sighted, f := v.Service.MatchingSightings("delta", false)
	assert.FatalOnError(t, f)
	assert.Count(t, 1, sighted)
}
