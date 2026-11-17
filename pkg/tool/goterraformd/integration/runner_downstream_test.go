package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	kubernetesMock "github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/provision/client/mock_client"
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/types/host_event"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/unit/service_tester"
	"testing"
)

func TestApplySendsHostEventsDownstream(t *testing.T) {
	d := mock_client.New(t, "alfa:8080")
	s := base.New(
		t,
		kubernetesMock.New(),
		service_tester.WithPlan(constant.FixturePlan),
		service_tester.WithDownstream(d),
	)
	s.Start()
	s.WaitForRuns(1)
	assert.Eventually(t, func() bool { return d.Calls() >= 1 })
	triggered := d.Triggered()
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
		triggered[0].Hosts,
	)
	assert.Count(t, 8, triggered[0].Addresses)
}

func TestRestartingDownstreamIsHeldOffAndNotCaptured(t *testing.T) {
	d := mock_client.New(t, "alfa:8080")
	d.Fail(d.PostTrigger, unavailable.Format(provisionConstant.RunnerStopped))
	s := base.New(
		t,
		kubernetesMock.New(),
		service_tester.WithPlan(constant.FixturePlan),
		service_tester.WithDownstream(d),
	)
	s.Start()
	s.WaitForRuns(1)
	assert.Eventually(t, func() bool { return d.Calls() >= 1 })
	s.TriggerServed()
	s.WaitForRuns(2)
	assert.Integer(t, 1, d.Calls())
	assert.Count(t, 0, s.Reporter.Events())
}
