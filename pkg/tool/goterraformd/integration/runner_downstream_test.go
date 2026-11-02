package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	kubernetesMock "github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/provision/client/mock_client"
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/types/host_event"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/unit/terraform_tester"
	"testing"
)

func TestApplySendsHostEventsDownstream(t *testing.T) {
	d := mock_client.New("alfa:8080")
	defer d.Close()
	terraform_tester.NewWithDownstream(
		t,
		kubernetesMock.New(),
		constant.FixturePlan,
		d,
	)
	waitForCalls(t, d, 1)
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

func TestARestartingDownstreamIsHeldOffAndNotCaptured(t *testing.T) {
	d := mock_client.New("alfa:8080")
	defer d.Close()
	d.Fail(d.PostTrigger, unavailable.Format(provisionConstant.RunnerStopped))
	s := terraform_tester.NewWithDownstream(
		t,
		kubernetesMock.New(),
		constant.FixturePlan,
		d,
	)
	waitForCalls(t, d, 1)
	response := make(chan *trigger.Result, 1)
	assert.FatalOnError(
		t,
		s.Service.Trigger(trigger.Request{Response: response}),
	)
	assert.FatalOnError(t, (<-response).Error)
	assert.Integer(t, 1, d.Calls())
	assert.Count(t, 0, s.Reporter.Events())
}
