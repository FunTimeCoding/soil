package chromium

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func (c *Client) History(identifier string) (*HistoryResult, error) {
	t := c.AcquireTarget(identifier)
	var currentIndex int64
	var entries []*page.NavigationEntry
	e := chromedp.Run(
		t,
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

	result := &HistoryResult{CurrentIndex: currentIndex}

	for _, entry := range entries {
		result.Entries = append(
			result.Entries,
			&HistoryEntry{Title: entry.Title, URL: entry.URL},
		)
	}

	return result, nil
}
