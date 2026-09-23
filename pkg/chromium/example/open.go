package example

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"time"
)

func Open() {
	c := chromium.NewEnvironment()
	defer c.Close()
	c.Run(
		chromedp.Navigate(locator.New("pkg.go.dev").Path("/time").String()),
		chromedp.Sleep(1*time.Second),
	)
}
