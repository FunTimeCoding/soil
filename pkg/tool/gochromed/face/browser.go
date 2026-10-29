package face

import (
	"github.com/funtimecoding/soil/pkg/chromium/history"
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/fill"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/saved"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/view"
)

type Browser interface {
	Tabs() ([]*tab.Tab, error)
	CreateTab(l string) (string, error)
	CloseTab(v *selector.Selector) (*tab.Tab, error)
	Wake(v *selector.Selector) (*tab.Tab, error)
	Navigate(
		v *selector.Selector,
		l string,
	) error
	Snapshot(v *selector.Selector) (*view.View, error)
	Click(
		v *selector.Selector,
		uid string,
		snapshot bool,
	) (*view.View, error)
	Fill(r *fill.Fill) (*view.View, error)
	Evaluate(
		v *selector.Selector,
		expression string,
	) (any, error)
	Screenshot(v *selector.Selector) (*saved.Saved, error)
	ReadBody(v *selector.Selector) (*saved.Saved, error)
	History(v *selector.Selector) (*history.Result, error)
}
