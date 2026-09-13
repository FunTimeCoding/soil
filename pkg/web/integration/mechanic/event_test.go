//go:build browser

package mechanic

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/browser_tester"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestEventAfterRequestResetsForm(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.NotifyPath))
	b.WaitVisible("#note-input")
	b.Fill("#note-input", "hello")
	b.Click("#receipt-post")
	b.WaitCondition(
		"document.querySelector('#receipt').textContent.indexOf('received hello') >= 0",
	)
	b.WaitCondition("document.querySelector('#note-input').value === ''")
}

func TestEventOutOfBandSwap(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.NotifyPath))
	b.WaitVisible("#note-input")
	b.Fill("#note-input", "hello")
	b.Click("#receipt-post")
	b.WaitCondition(
		"document.querySelector('#oob-summary').textContent.indexOf('sent 1') >= 0",
	)
	assert.StringContains(t, "sent 1", b.Text("#oob-summary"))
}

func TestEventFailureRaisesNotification(t *testing.T) {
	address := serve(t)
	b := browser_tester.New(t)
	b.Navigate(page(address, constant.NotifyPath))
	b.WaitVisible("#fail-post")
	assert.Integer(t, 0, b.CountElements("#notifications .notification-error"))
	b.Click("#fail-post")
	b.WaitCondition(
		"document.querySelectorAll('#notifications > *').length > 0",
	)
	assert.String(t, "", b.Text("#receipt"))
}
