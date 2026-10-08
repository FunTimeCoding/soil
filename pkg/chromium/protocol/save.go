package protocol

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/time/constant"
	"os"
	"time"
)

func (p *Protocol) Save(
	locator string,
	filename string,
) error {
	console.Line("  Save")
	start := time.Now()
	console.Format("    Start %v\n", start.Format(constant.Micro))
	b, e := run(
		p,
		func(
			o context.Context,
			s *chromedp.Target,
		) ([]byte, error) {
			t2 := time.Now()
			console.Format(
				"    GetResourceTree %v\n",
				t2.Format(constant.Micro),
			)
			t, e := cdp.Call(o, s, page.GetResourceTree, cdp.Empty{})

			if e != nil {
				console.Format(
					"    GetResourceTree fail %v: %v\n",
					time.Since(t2),
					e,
				)

				return nil, e
			}

			console.Format("    GetResourceTree took %v\n", time.Since(t2))
			t3 := time.Now()
			r, e := cdp.Call(
				o,
				s,
				page.GetResourceContent,
				page.GetResourceContentParams{
					FrameID: t.FrameTree.Frame.ID,
					URL:     locator,
				},
			)

			if e != nil {
				console.Format(
					"    GetResourceContent fail %v: %v\n",
					time.Since(t3),
					e,
				)

				return nil, e
			}

			console.Format("    GetResourceContent took %v\n", time.Since(t3))

			return r.Content, nil
		},
	)

	if e != nil {
		return e
	}

	if f := os.WriteFile(filename, b, 0644); f != nil {
		return f
	}

	console.Format("    Complete after: %v\n", time.Since(start))

	return nil
}
