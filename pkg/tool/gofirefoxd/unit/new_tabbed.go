package unit

import (
	"github.com/funtimecoding/soil/pkg/firefox/tab"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/unit/model_context_tester"
	"testing"
)

func newTabbed(t *testing.T) *model_context_tester.Tester {
	t.Helper()
	o := model_context_tester.New(t)
	b := tab.New()
	b.Identifier = 100
	b.Locator = "https://example.com"
	b.Title = "Example"
	o.MockClient.AddTab(b)

	return o
}
