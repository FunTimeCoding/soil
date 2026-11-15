package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/firefox/tab"
	"testing"
)

func NewTabbed(t *testing.T) *Tester {
	t.Helper()
	o := New(t)
	b := tab.New()
	b.Identifier = 100
	b.Locator = "https://example.com"
	b.Title = "Example"
	o.MockClient.AddTab(b)

	return o
}
