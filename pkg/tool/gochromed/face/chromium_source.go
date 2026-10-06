package face

import "github.com/funtimecoding/soil/pkg/chromium/tab"

type ChromiumSource interface {
	Wake(identifier string) error
	Tabs() ([]*tab.Tab, error)
	TabByHost(s string) (*tab.Tab, error)
	CreateTab(l string) (string, error)
	CloseTab(identifier string) error
	Page(identifier string) Page
}
