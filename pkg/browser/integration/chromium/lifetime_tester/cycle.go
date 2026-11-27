//go:build browser

package lifetime_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/browser/integration/chromium/base"
	"time"
)

func (o *Tester) Cycle(s *base.Stack) {
	o.t.Helper()
	identifier := s.OpenTab(constant.ChromiumFixtureQuietRoute)
	assert.StringContains(o.t, "quiet", s.Body(identifier))
	assert.FatalOnError(o.t, s.Client.CloseTab(identifier))

	for i := 0; i < 50; i++ {
		if !s.TabAlive(identifier) {
			return
		}

		time.Sleep(100 * time.Millisecond)
	}

	o.t.Fatalf("tab %s never went away", identifier)
}
