package chromium

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"os"
	"time"
)

func (c *Client) Save(
	x context.Context,
	locator string,
	filename string,
) {
	console.Line("  Save")
	start := time.Now()
	console.Format("    Start %v\n", start.Format(constant.Micro))
	var b []byte
	c.RunContext(
		x,
		chromedp.ActionFunc(
			func(o context.Context) error {
				t2 := time.Now()
				console.Format(
					"    GetResourceTree %v\n",
					t2.Format(constant.Micro),
				)
				t, e := page.GetResourceTree().Do(o)

				if e != nil {
					console.Format(
						"    GetResourceTree fail %v: %v\n",
						time.Since(t2),
						e,
					)

					return e
				}

				console.Format("    GetResourceTree took %v\n", time.Since(t2))
				t3 := time.Now()
				b, e = page.GetResourceContent(t.Frame.ID, locator).Do(o)

				if e != nil {
					console.Format(
						"    GetResourceContent fail %v: %v\n",
						time.Since(t3),
						e,
					)

					return e
				}

				console.Format(
					"    GetResourceContent took %v\n",
					time.Since(t3),
				)

				return nil
			},
		),
	)
	errors.PanicOnError(os.WriteFile(filename, b, 0644))
	console.Format("    Complete after: %v\n", time.Since(start))
}
