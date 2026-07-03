package frame_probe

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chromium"
	"github.com/funtimecoding/go-library/pkg/chromium/constant"
)

func printIframeRoots(c *chromium.Client) {
	fmt.Println("=== iframe target root frames ===")

	for _, t := range c.Tabs() {
		if t.Type != constant.IframeTabType {
			continue
		}

		root := frameTree(c, t.Identifier).Frame
		fmt.Printf(
			"target %s root frame %s match=%v %s\n",
			t.Identifier,
			root.ID,
			string(root.ID) == t.Identifier,
			t.Locator,
		)
	}
}
