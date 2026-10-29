package ansible_tester

import (
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/model/run"
	"github.com/funtimecoding/soil/pkg/provision/types/host_event"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
)

func (o *Tester) TriggerHosts(events []*host_event.Event) []*run.Run {
	o.t.Helper()
	response := make(chan *trigger.Result, 1)

	if e := o.Service.Trigger(
		trigger.Request{
			Parameters: map[string]any{constant.DownstreamHosts: events},
			Response:   response,
		},
	); e != nil {
		o.t.Fatal(e)
	}
	result := <-response

	if result.Error != nil {
		o.t.Fatal(result.Error)
	}

	return result.Value.([]*run.Run)
}
