package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/types/host_event"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/ansible_tester"
	"testing"
)

func TestHostEventsLimitPlaybooksToLiveHosts(t *testing.T) {
	o := ansible_tester.New(t, nil)
	o.WaitForRun()
	runs := o.TriggerHosts(
		[]*host_event.Event{
			host_event.New("alfa", provisionConstant.HostUpdate),
			host_event.New("bravo", provisionConstant.HostDelete),
			host_event.New("charlie", provisionConstant.HostReplace),
		},
	)
	assert.Count(t, 1, runs)
	assert.String(t, "success", runs[0].Status)
	arguments, okay := o.Captured(constant.FixtureArguments)
	assert.True(t, okay)
	assert.String(t, "--limit\nalfa,charlie\nbravo.yaml\n", arguments)
}

func TestHostEventsReachingNoLiveHostSkipRun(t *testing.T) {
	o := ansible_tester.New(t, nil)
	o.WaitForRun()
	runs := o.TriggerHosts(
		[]*host_event.Event{
			host_event.New("delta", provisionConstant.HostDelete),
		},
	)
	assert.Count(t, 0, runs)
	recent, e := o.Service.Recent(10)
	errors.PanicOnError(e)
	assert.Count(t, 1, recent)
}
