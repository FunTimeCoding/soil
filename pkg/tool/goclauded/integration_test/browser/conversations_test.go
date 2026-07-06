//go:build browser

package browser

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/browser_tester"
	"testing"
	"time"
)

func TestConversationsSidebarInfiniteScroll(t *testing.T) {
	b := browser_tester.New(t)
	b.Navigate("http://localhost:8583/conversations")
	b.WaitReady(".sidebar-entry")
	initial := b.CountElements(".sidebar-entry")
	assert.True(t, initial > 0)
	fmt.Printf("initial entries: %d\n", initial)
	b.ScrollToBottom(".sidebar")
	time.Sleep(2 * time.Second)
	after := b.CountElements(".sidebar-entry")
	fmt.Printf("after scroll: %d\n", after)
	assert.True(t, after > initial)
}

func TestConversationsSidebarFilter(t *testing.T) {
	b := browser_tester.New(t)
	b.Navigate("http://localhost:8583/conversations")
	b.WaitReady(".sidebar-entry")
	total := b.CountElements(".sidebar-entry")
	b.Evaluate(
		"document.querySelector('.sidebar-filter').value = 'goclauded'; document.querySelector('.sidebar-filter').dispatchEvent(new Event('input'))",
		nil,
	)
	time.Sleep(500 * time.Millisecond)
	var visible int
	b.Evaluate(
		"document.querySelectorAll('.sidebar-entry:not([style*=\"display: none\"])').length",
		&visible,
	)
	assert.True(t, visible > 0)
	assert.True(t, visible < total)
}
