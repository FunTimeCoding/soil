package frame_probe

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chromium"
	"github.com/funtimecoding/go-library/pkg/chromium/constant"
)

func printFrameTrees(c *chromium.Client) {
	fmt.Println("=== Page.getFrameTree per page tab ===")

	for _, t := range c.Tabs() {
		if t.Type != constant.PageTabType {
			continue
		}

		fmt.Printf("tab %s %s\n", t.Identifier, t.Locator)
		printFrame(frameTree(c, t.Identifier), 1)
	}
}
