package frame_probe

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func printIframeRoots(c *chromium.Client) {
	console.Line("=== iframe target root frames ===")

	for _, t := range c.Tabs() {
		if t.Type != constant.IframeTabType {
			continue
		}

		root := frameTree(c, t.Identifier).Frame
		console.Format(
			"target %s root frame %s match=%v %s\n",
			t.Identifier,
			root.ID,
			string(root.ID) == t.Identifier,
			t.Locator,
		)
	}
}
