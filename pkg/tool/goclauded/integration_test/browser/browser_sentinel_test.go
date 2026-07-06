//go:build browser

package browser

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/browser_tester"
	"testing"
)

func TestBrowserSentinelDebug(t *testing.T) {
	b := browser_tester.New(t)
	b.Navigate("http://localhost:8583/conversations")
	b.WaitReady(".sidebar-entry")
	var sentinelHTML string
	b.Evaluate(
		"(() => { const el = document.querySelector('[hx-trigger=\"revealed\"]'); return el ? el.outerHTML : 'NOT FOUND'; })()",
		&sentinelHTML,
	)
	fmt.Printf("sentinel: %s\n", sentinelHTML)
	var sidebarHeight float64
	b.Evaluate(
		"document.querySelector('.sidebar').scrollHeight",
		&sidebarHeight,
	)
	var sidebarClient float64
	b.Evaluate(
		"document.querySelector('.sidebar').clientHeight",
		&sidebarClient,
	)
	fmt.Printf(
		"sidebar scrollHeight: %.0f, clientHeight: %.0f\n",
		sidebarHeight,
		sidebarClient,
	)
	b.ScrollToBottom(".sidebar")
	var scrollTop float64
	b.Evaluate("document.querySelector('.sidebar').scrollTop", &scrollTop)
	fmt.Printf("after scroll, scrollTop: %.0f\n", scrollTop)
}
