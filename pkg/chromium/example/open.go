package example

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"time"
)

func Open() {
	c := chromium.NewEnvironment()
	defer c.Close()
	errors.PanicOnError(
		c.Run(
			chromedp.Navigate(locator.New("pkg.go.dev").Path("/time").String()),
			chromedp.Sleep(1*time.Second),
		),
	)
}
