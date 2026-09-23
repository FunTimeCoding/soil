package protocol

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium/history"
	"github.com/funtimecoding/soil/pkg/chromium/history/entry"
)

func (p *Protocol) History() (*history.Result, error) {
	var currentIndex int64
	var entries []*page.NavigationEntry
	e := chromedp.Run(
		p.context,
		chromedp.ActionFunc(
			func(x context.Context) error {
				var e error
				currentIndex, entries, e = page.GetNavigationHistory().Do(x)

				return e
			},
		),
	)

	if e != nil {
		return nil, e
	}

	result := history.New(currentIndex)

	for _, n := range entries {
		result.Entries = append(result.Entries, entry.New(n.Title, n.URL))
	}

	return result, nil
}
