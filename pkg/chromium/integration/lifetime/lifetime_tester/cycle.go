//go:build browser

package lifetime_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"testing"
	"time"
)

func Cycle(
	t *testing.T,
	s *base.Stack,
) {
	t.Helper()
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	assert.StringContains(t, "quiet", s.Body(identifier))
	assert.FatalOnError(t, s.Client.CloseTab(identifier))

	for i := 0; i < 50; i++ {
		if !s.TabAlive(identifier) {
			return
		}

		time.Sleep(100 * time.Millisecond)
	}

	t.Fatalf("tab %s never went away", identifier)
}
