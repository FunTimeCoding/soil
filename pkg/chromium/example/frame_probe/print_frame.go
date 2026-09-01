package frame_probe

import (
	"github.com/chromedp/cdproto/page"
	"github.com/funtimecoding/soil/pkg/console"
	"strings"
)

func printFrame(
	t *page.FrameTree,
	depth int,
) {
	console.Format(
		"%sframe %s %s\n",
		strings.Repeat("  ", depth),
		t.Frame.ID,
		t.Frame.URL,
	)

	for _, child := range t.ChildFrames {
		printFrame(child, depth+1)
	}
}
