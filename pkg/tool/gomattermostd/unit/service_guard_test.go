package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"testing"
)

func TestUnavailableMattermostClosesGuard(t *testing.T) {
	o := service_tester.New(t)
	o.Mattermost.AddPost("alfa")
	s := o.Service()
	o.Mattermost.Failure = unavailable.Format("mattermost maintenance")
	_, e := s.FindPost("alfa")
	assert.True(t, unavailable.Is(e))
	o.Mattermost.Failure = nil
	_, f := s.FindPost("alfa")
	assert.True(t, unavailable.Is(f))
	assert.Count(t, 0, o.Reporter.Events())
}

func TestQuietRefusalLeavesGuardOpen(t *testing.T) {
	o := service_tester.New(t)
	o.Mattermost.AddPost("alfa")
	s := o.Service()
	_, e := s.FindPost("missing")
	assert.True(t, not_found.Is(e))
	p, f := s.FindPost("alfa")
	assert.FatalOnError(t, f)
	assert.String(t, "alfa", p.Id)
	assert.Count(t, 0, o.Reporter.Events())
}
