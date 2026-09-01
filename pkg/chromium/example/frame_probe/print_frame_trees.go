package frame_probe

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func printFrameTrees(c *chromium.Client) {
	console.Line("=== Page.getFrameTree per page tab ===")

	for _, t := range c.Tabs() {
		if t.Type != constant.PageTabType {
			continue
		}

		console.Format("tab %s %s\n", t.Identifier, t.Locator)
		printFrame(frameTree(c, t.Identifier), 1)
	}
}
