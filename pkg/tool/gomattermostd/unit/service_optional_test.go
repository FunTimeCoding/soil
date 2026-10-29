package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"testing"
)

func TestSubscriptionsNeedGoclauded(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	m.AddPost("alfa")
	s := service_tester.Service(t, m, nil, nil)
	assert.False(t, s.Subscribing())
	assert.Count(t, 1, s.Workers())
	p, e := m.FindPost("alfa")
	assert.FatalOnError(t, e)
	_, _, f := s.Subscribe(p, "kilo", "")
	assert.True(t, not_configured.Is(f))
	_, g := s.Unsubscribe("kilo", "alfa")
	assert.True(t, not_configured.Is(g))
	_, h := s.Subscriptions("kilo")
	assert.True(t, not_configured.Is(h))
}

func TestGoclaudedBuildsTheSubscriptionWorkers(t *testing.T) {
	_, k := worker_tester.NewSink(t)
	s := service_tester.Service(
		t,
		service_tester.NewMattermost(t, "kilo"),
		k,
		nil,
	)
	assert.True(t, s.Subscribing())
	assert.Count(t, 3, s.Workers())
}

func TestMonitorNeedsItsConfiguration(t *testing.T) {
	s := service_tester.Service(
		t,
		service_tester.NewMattermost(t, "kilo"),
		nil,
		nil,
	)
	assert.False(t, s.Monitoring())
	assert.True(t, not_configured.Is(s.RunMonitor()))
}
