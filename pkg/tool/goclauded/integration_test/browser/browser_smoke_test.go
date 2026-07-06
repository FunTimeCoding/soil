//go:build browser

package browser

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/browser_tester"
	"testing"
)

func TestBrowserSmoke(t *testing.T) {
	b := browser_tester.New(t)
	b.Navigate("http://localhost:8583/conversations")
	var title string
	b.Evaluate("document.title", &title)
	fmt.Printf("title: %q\n", title)
	var m string
	b.Evaluate("document.body.innerHTML.substring(0, 500)", &m)
	fmt.Printf("body: %s\n", m)
}
