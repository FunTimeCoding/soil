//go:build browser

package lifetime

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"testing"
)

func TestSharedClientReuseKeepsTab(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)

	for i := 0; i < constant.FixtureSharedCycleCount; i++ {
		assert.StringContains(t, "quiet", s.Body(identifier))
	}

	assert.Integer(t, 1, s.Client.TargetCount())
	s.AssertTabAlive(identifier)
}
