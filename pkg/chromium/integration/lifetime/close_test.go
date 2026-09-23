//go:build browser

package lifetime

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"testing"
)

func TestClientCloseAfterAcquireKeepsTabOpen(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	assert.StringContains(t, "quiet", s.Body(identifier))
	s.Client.Close()
	s.AssertTabAlive(identifier)
}

func TestClientCloseAfterAcquireOnBusyPageKeepsTabOpen(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureBusyRoute)
	assert.StringContains(t, "busy", s.Body(identifier))
	s.Client.Close()
	s.AssertTabAlive(identifier)
}

func TestClientCloseLeavesAcquiredContextUncancelled(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	x := s.Client.TargetContext(identifier)
	assert.StringContains(t, "quiet", s.Body(identifier))
	s.Client.Close()
	assert.Nil(t, x.Err())
	s.AssertTabAlive(identifier)
}

func TestClientCloseWithoutAcquireKeepsTabOpen(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	fresh := chromium.New("localhost", s.Port)
	fresh.Close()
	s.AssertTabAlive(identifier)
}
