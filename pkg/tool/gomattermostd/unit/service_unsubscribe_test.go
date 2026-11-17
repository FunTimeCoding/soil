package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"testing"
)

func TestUnsubscribeForgetsRootOnlyWhenItsLastSubscriberLeaves(t *testing.T) {
	_, k := worker_tester.NewSink(t)
	o := service_tester.New(t, service_tester.WithConnector(k))
	o.Mattermost.AddPost("alfa")
	s := o.Service()
	o.Subscribe("alfa", "kilo", "")
	o.Subscribe("alfa", "lima", "")
	assert.True(t, s.Watching("alfa"))
	_, e := s.Unsubscribe("kilo", "alfa")
	assert.FatalOnError(t, e)
	assert.True(t, s.Watching("alfa"))
	_, f := s.Unsubscribe("lima", "alfa")
	assert.FatalOnError(t, f)
	assert.False(t, s.Watching("alfa"))
}
