package protocol

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/history"
	"github.com/funtimecoding/soil/pkg/chromium/history/entry"
)

func (p *Protocol) History() (*history.Result, error) {
	h, e := run(p, chromedp.NavigationEntries())

	if e != nil {
		return nil, e
	}

	result := history.New(h.CurrentIndex)

	for _, n := range h.Entries {
		result.Entries = append(result.Entries, entry.New(n.Title, n.URL))
	}

	return result, nil
}
