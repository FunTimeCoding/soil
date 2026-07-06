package chromium

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	port int,
) *Client {
	allocator, allocatorCancel := chromedp.NewRemoteAllocator(
		context.Background(),
		locator.New(host).Port(port).Scheme(constant.Socket).String(),
	)
	c, cancel := chromedp.NewContext(allocator)
	// Force CDP websocket initialization so Browser is non-nil for
	// methods that access it directly (e.g. CloseTab).
	_, e := chromedp.Targets(c)
	errors.PanicOnError(e)
	result := &Client{
		host:            host,
		port:            port,
		allocator:       allocator,
		allocatorCancel: allocatorCancel,
		context:         c,
		cancel:          cancel,
		targets:         make(map[string]context.Context),
	}
	result.listenTargets()

	return result
}
