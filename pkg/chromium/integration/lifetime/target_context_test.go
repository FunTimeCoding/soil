//go:build browser

package lifetime

import (
	"context"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/chromium/integration/base"
	"testing"
)

func TestCancelDerivedTargetContextKeepsTab(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	x, cancel := context.WithCancel(s.Client.TargetContext(identifier))
	var body string
	s.Client.RunContext(x, chromedp.OuterHTML(constant.BodySelector, &body))
	assert.StringContains(t, "quiet", body)
	cancel()
	s.AssertTabAlive(identifier)
}

func TestCancelAcquiredTargetContextTakesTab(t *testing.T) {
	s := base.New(t)
	identifier := s.OpenTab(constant.FixtureQuietRoute)
	x, cancel := chromedp.NewContext(
		s.Client.Context(),
		chromedp.WithTargetID(target.ID(identifier)),
	)
	var body string
	s.Client.RunContext(x, chromedp.OuterHTML(constant.BodySelector, &body))
	assert.StringContains(t, "quiet", body)
	cancel()
	s.AssertTabGone(identifier)
}
