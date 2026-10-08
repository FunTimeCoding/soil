package chromium

import (
	"context"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/chromedp/remote"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
)

func New(
	host string,
	port int,
) *Client {
	allocator, allocatorCancel := remote.NewAllocator(
		context.Background(),
		locator.New(host).Port(port).Scheme(constant.Socket).String(),
	)
	c, cancel := chromedp.NewContext(allocator)
	result := &Client{
		host:            host,
		port:            port,
		requester:       requester.New(locator.New(host).Port(port).Insecure()),
		allocator:       allocator,
		allocatorCancel: allocatorCancel,
		context:         c,
		cancel:          cancel,
		targets:         make(map[string]context.Context),
	}
	errors.PanicOnError(result.listenTargets())

	return result
}
