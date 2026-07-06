package chromium

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/errors"
	library "github.com/funtimecoding/soil/pkg/time"
	"os"
	"time"
)

func (c *Client) Save(
	x context.Context,
	locator string,
	filename string,
) {
	fmt.Println("  Save")
	start := time.Now()
	fmt.Printf("    Start %v\n", start.Format(library.Micro))
	var b []byte
	c.RunContext(
		x,
		chromedp.ActionFunc(
			func(o context.Context) error {
				t2 := time.Now()
				fmt.Printf(
					"    GetResourceTree %v\n",
					t2.Format(library.Micro),
				)
				t, e := page.GetResourceTree().Do(o)

				if e != nil {
					fmt.Printf(
						"    GetResourceTree fail %v: %v\n",
						time.Since(t2),
						e,
					)

					return e
				}

				fmt.Printf(
					"    GetResourceTree took %v\n",
					time.Since(t2),
				)
				t3 := time.Now()
				b, e = page.GetResourceContent(t.Frame.ID, locator).Do(o)

				if e != nil {
					fmt.Printf(
						"    GetResourceContent fail %v: %v\n",
						time.Since(t3),
						e,
					)

					return e
				}

				fmt.Printf(
					"    GetResourceContent took %v\n",
					time.Since(t3),
				)

				return nil
			},
		),
	)
	errors.PanicOnError(os.WriteFile(filename, b, 0644))
	fmt.Printf("    Complete after: %v\n", time.Since(start))
}
