package face

import (
	"context"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/event"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/face"
)

type ChromiumClient interface {
	AcquireTarget(identifier string) (context.Context, error)
	Activate(targetIdentifier string)
	CloseTab(identifier string) error
	Context() context.Context
	CreateTab(l string) (string, error)
	MustHostTab(s string) *tab.Tab
	MustTabs() []*tab.Tab
	Run(steps ...chromedp.Action[chromedp.Void]) error
	RunContext(
		o context.Context,
		steps ...chromedp.Action[chromedp.Void],
	) error
	HostTab(s string) (*tab.Tab, error)
	Tabs() ([]*tab.Tab, error)
	TargetContext(identifier string) (context.Context, error)
	TargetCount() int
	TargetIdentifiers() []string
	Targets() ([]*target.Info, error)
	Wake(identifier string) error
	Watch(observe func(*event.Event)) error
	Address() string
	face.Reachable
	Close()
}
