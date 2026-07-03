package frame_probe

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/chromium"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func printTargets(c *chromium.Client) {
	fmt.Println("=== Target.getTargets ===")

	for _, t := range c.Targets() {
		fmt.Println(notation.MarshalIndent(t))
	}
}
