package face

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/history"
	"github.com/funtimecoding/soil/pkg/chromium/snapshot"
	"time"
)

type ProtocolClient interface {
	Body() (string, error)
	ClickNode(backendNodeIdentifier int64) error
	ClickQuery(s string) error
	ClickSearch(s string) error
	ClickTrusted(s string) error
	EnterText(
		s string,
		text string,
	) error
	Evaluate(
		expression string,
		result any,
	) error
	EvaluatePromise(
		expression string,
		result any,
	) error
	FillNode(
		backendNodeIdentifier int64,
		value string,
		direct bool,
	) error
	FindNode(
		s string,
		index int,
	) (*chromedp.Node, bool, error)
	Focused() (map[string]any, error)
	HasNodes(s string) (bool, error)
	History() (*history.Result, error)
	Location() (string, error)
	MustBody() string
	MustClickQuery(s string)
	MustClickSearch(s string)
	MustEnterText(
		s string,
		text string,
	)
	MustFindNode(
		s string,
		index int,
	) (*chromedp.Node, bool)
	MustHasNodes(s string) bool
	MustOuter(s string) string
	MustSave(
		locator string,
		filename string,
	)
	MustWaitVisible(s string)
	Navigate(l string) error
	NeedReload(locator string) bool
	Outer(s string) (string, error)
	PrintNode(
		s string,
		attribute []string,
	) error
	Reload() error
	Save(
		locator string,
		filename string,
	) error
	Screenshot() ([]byte, error)
	SetValue(
		s string,
		value string,
	) error
	Snapshot() ([]*snapshot.Node, error)
	WaitVisible(s string) error
	WithTimeout(d time.Duration) ProtocolClient
}
