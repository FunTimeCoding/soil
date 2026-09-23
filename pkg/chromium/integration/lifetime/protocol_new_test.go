//go:build browser

package lifetime

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"github.com/funtimecoding/soil/pkg/chromium/integration/lifetime/lifetime_tester"
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"testing"
)

func TestProtocolReadsBorrowedTab(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	p := protocol.New(s.Client, constant.FixtureQuietRoute)
	assert.StringContains(t, "quiet", p.Body())
	assert.Integer(t, 1, s.Client.TargetCount())
	s.AssertTabAlive(identifier)
}

func TestProtocolAbsentTabLeavesClientUsable(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	assert.String(
		t,
		"tab not found",
		lifetime_tester.ProtocolAbsent(t, s.Client),
	)
	assert.Integer(t, 0, s.Client.TargetCount())
	assert.NotNil(t, s.Client.TabByHost(constant.FixtureQuietRoute))
	s.AssertTabAlive(identifier)
}
