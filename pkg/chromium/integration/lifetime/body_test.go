//go:build browser

package lifetime

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"testing"
)

func TestBodyKeepsTabOpen(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	assert.StringContains(t, "quiet", s.Body(identifier))
	s.AssertTabAlive(identifier)
}

func TestBodyOnStalledPageKeepsTabOpen(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureStalledRoute)
	assert.StringContains(t, "stalled", s.Body(identifier))
	s.AssertTabAlive(identifier)
}
