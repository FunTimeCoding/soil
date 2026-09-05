//go:build browser

package browser

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/browser_tester"
	"testing"
)

func TestBrowserSmoke(t *testing.T) {
	b := browser_tester.New(t)
	b.Navigate("http://localhost:8583/conversations")
	var title string
	b.Evaluate("document.title", &title)
	console.Format("title: %q\n", title)
	var m string
	b.Evaluate("document.body.innerHTML.substring(0, 500)", &m)
	console.Format("body: %s\n", m)
}
