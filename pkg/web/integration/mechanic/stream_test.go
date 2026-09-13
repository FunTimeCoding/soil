//go:build browser

package mechanic

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/browser_tester"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestStreamConnects(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.StreamPath))
	b.WaitVisible("#pulse-post")
	assert.StringContains(t, "pulse 0", b.Text("#pulse"))
}

func TestStreamPushesNamedEvent(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.StreamPath))
	b.WaitVisible("#pulse-post")
	b.Click("#pulse-post")
	b.WaitCondition(
		"document.querySelector('#pulse').textContent.indexOf('pulse 1') >= 0",
	)
	assert.StringContains(t, "pulse 1", b.Text("#pulse"))
}

func TestStreamPushesSecondNamedEvent(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.StreamPath))
	b.WaitVisible("#pulse-post")
	b.Click("#pulse-post")
	assert.StringContains(t, "count 0", b.Text("#tick"))
	b.WaitCondition(
		"document.querySelector('#tick').textContent.indexOf('count 1') >= 0",
	)
	assert.StringContains(t, "count 1", b.Text("#tick"))
}
