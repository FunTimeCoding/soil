package face

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/snapshot"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
)

type ChromiumSource interface {
	Wake(identifier string) error
	Tabs() []*tab.Tab
	TabByHost(s string) *tab.Tab
	AcquireTarget(identifier string) context.Context
	TargetContext(identifier string) context.Context
	RunContext(
		o context.Context,
		a ...chromedp.Action,
	)
	CreateTab(l string) (string, error)
	CloseTab(identifier string) error
	Navigate(
		x context.Context,
		l string,
	) error
	Body(identifier string) string
	Snapshot(x context.Context) ([]*snapshot.Node, error)
	Screenshot(x context.Context) ([]byte, error)
	Evaluate(
		x context.Context,
		expression string,
		result any,
	) error
	ClickNode(
		x context.Context,
		backendNodeIdentifier int64,
	) error
	FillNode(
		x context.Context,
		backendNodeIdentifier int64,
		value string,
		direct bool,
	) error
	History(identifier string) (*chromium.HistoryResult, error)
}
