package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"testing"
)

func TestUnsubscribeForgetsRootOnlyWhenItsLastSubscriberLeaves(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	m.AddPost("alfa")
	_, k := worker_tester.NewSink(t)
	s := service_tester.Service(t, m, k, nil)
	subscribe(t, s, m, "alfa", "kilo", "")
	subscribe(t, s, m, "alfa", "lima", "")
	assert.True(t, s.Watching("alfa"))
	_, e := s.Unsubscribe("kilo", "alfa")
	assert.FatalOnError(t, e)
	assert.True(t, s.Watching("alfa"))
	_, f := s.Unsubscribe("lima", "alfa")
	assert.FatalOnError(t, f)
	assert.False(t, s.Watching("alfa"))
}
