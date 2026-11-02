package face

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
)

type ChromiumSource interface {
	Reachable(x context.Context) error
	Wake(identifier string) error
	Tabs() ([]*tab.Tab, error)
	TabByHost(s string) (*tab.Tab, error)
	CreateTab(l string) (string, error)
	CloseTab(identifier string) error
	Page(identifier string) Page
}
