package example

import (
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"log"
	"strings"
	"time"
)

func Open() {
	c := chromium.NewEnvironment()
	defer c.Close()
	var result string
	c.Run(
		chromedp.Navigate(
			locator.New("pkg.go.dev").Path("/time").String(),
		),
		//chromedp.Text(
		//	`.Documentation-overview`,
		//	&result,
		//	chromedp.NodeVisible,
		//),
		chromedp.Sleep(1*time.Second),
	)
	log.Println(strings.TrimSpace(result))
}
