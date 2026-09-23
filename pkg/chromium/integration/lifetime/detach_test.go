//go:build browser

package lifetime

import (
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"github.com/funtimecoding/soil/pkg/chromium/integration/lifetime/lifetime_tester"
	"testing"
)

func TestDetachCarriesSessionOnly(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	detached := lifetime_tester.WatchKind(t, s, constant.EventKindDetached)
	x, cancel := chromedp.NewContext(
		s.Client.Context(),
		chromedp.WithTargetID(target.ID(identifier)),
	)
	var body string
	s.Client.RunContext(x, chromedp.OuterHTML(constant.BodySelector, &body))
	cancel()
	e := lifetime_tester.AwaitEvent(t, detached)
	assert.String(t, "", e.TargetIdentifier)
	assert.True(t, e.SessionIdentifier != "")
}
