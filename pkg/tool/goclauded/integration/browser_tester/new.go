package browser_tester

import (
	"context"
	"github.com/chromedp/chromedp"
	"testing"
)

func New(t *testing.T) *Browser {
	t.Helper()
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ExecPath(
			"/Applications/Brave Browser.app/Contents/MacOS/Brave Browser",
		),
		chromedp.Headless,
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.DisableGPU,
	)
	allocator, allocatorCancel := chromedp.NewExecAllocator(
		context.Background(),
		opts...,
	)
	c, cancel := chromedp.NewContext(allocator)
	t.Cleanup(
		func() {
			cancel()
			allocatorCancel()
		},
	)

	return &Browser{T: t, Context: c, cancel: cancel}
}
