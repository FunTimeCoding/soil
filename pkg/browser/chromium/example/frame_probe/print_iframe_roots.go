package frame_probe

import (
	"github.com/funtimecoding/soil/pkg/browser/chromium"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func printIframeRoots(c *chromium.Client) {
	console.Line("=== iframe target root frames ===")

	for _, t := range c.MustTabs() {
		if t.Type != constant.ChromiumIframeTabType {
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
