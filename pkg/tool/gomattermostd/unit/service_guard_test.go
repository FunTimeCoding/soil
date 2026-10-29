package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"testing"
)

func TestAnUnavailableMattermostClosesTheGuard(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	m.AddPost("alfa")
	r := memory.New()
	s := service_tester.Reported(t, m, nil, nil, r)
	m.Failure = unavailable.Format("mattermost maintenance")
	_, e := s.FindPost("alfa")
	assert.True(t, unavailable.Is(e))
	m.Failure = nil
	_, f := s.FindPost("alfa")
	assert.True(t, unavailable.Is(f))
	assert.Count(t, 0, r.Events())
}

func TestAQuietRefusalLeavesTheGuardOpen(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	m.AddPost("alfa")
	r := memory.New()
	s := service_tester.Reported(t, m, nil, nil, r)
	_, e := s.FindPost("missing")
	assert.True(t, not_found.Is(e))
	p, f := s.FindPost("alfa")
	assert.FatalOnError(t, f)
	assert.String(t, "alfa", p.Id)
	assert.Count(t, 0, r.Events())
}
