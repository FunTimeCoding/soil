//go:build browser

package browser

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/browser_tester"
	"testing"
)

func TestBrowserSentinelDebug(t *testing.T) {
	b := browser_tester.New(t)
	b.Navigate("http://localhost:8583/conversations")
	b.WaitReady(".sidebar-entry")
	var sentinelMarkup string
	b.Evaluate(
		"(() => { const el = document.querySelector('[hx-trigger=\"revealed\"]'); return el ? el.outerHTML : 'NOT FOUND'; })()",
		&sentinelMarkup,
	)
	console.Format("sentinel: %s\n", sentinelMarkup)
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
	console.Format(
		"sidebar scrollHeight: %.0f, clientHeight: %.0f\n",
		sidebarHeight,
		sidebarClient,
	)
	b.ScrollToBottom(".sidebar")
	var scrollTop float64
	b.Evaluate("document.querySelector('.sidebar').scrollTop", &scrollTop)
	console.Format("after scroll, scrollTop: %.0f\n", scrollTop)
}
