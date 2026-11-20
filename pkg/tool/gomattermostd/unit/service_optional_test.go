package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"testing"
)

func TestSubscriptionsNeedGoclauded(t *testing.T) {
	o := service_tester.New(t)
	o.Mattermost.AddPost("alfa")
	s := o.Service()
	assert.False(t, s.Subscribing())
	p, e := o.Mattermost.FindPost("alfa")
	assert.FatalOnError(t, e)
	_, _, f := s.Subscribe(p, "kilo", "", "")
	assert.True(t, not_configured.Is(f))
	_, g := s.Unsubscribe("kilo", "alfa")
	assert.True(t, not_configured.Is(g))
	_, h := s.Subscriptions("kilo")
	assert.True(t, not_configured.Is(h))
}

func TestGoclaudedTurnsSubscriptionsOn(t *testing.T) {
	_, k := worker_tester.New(t).Sink()
	s := service_tester.New(t, service_tester.WithConnector(k)).Service()
	assert.True(t, s.Subscribing())
}

func TestMonitorNeedsItsConfiguration(t *testing.T) {
	s := service_tester.New(t).Service()
	assert.False(t, s.Monitoring())
	assert.True(t, not_configured.Is(s.RunMonitor()))
}
