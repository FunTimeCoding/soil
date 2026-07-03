package frame_probe

import (
	"fmt"
	"github.com/chromedp/cdproto/page"
	"strings"
)

func printFrame(
	t *page.FrameTree,
	depth int,
) {
	fmt.Printf(
		"%sframe %s %s\n",
		strings.Repeat("  ", depth),
		t.Frame.ID,
		t.Frame.URL,
	)

	for _, child := range t.ChildFrames {
		printFrame(child, depth+1)
	}
}
