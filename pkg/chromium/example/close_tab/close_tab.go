package close_tab

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func CloseTab() {
	c := chromium.NewEnvironment()
	fmt.Println("listing tabs via HTTP...")

	for _, t := range c.Tabs() {
		if t.Type != constant.PageTabType {
			continue
		}

		fmt.Printf("  %s %s\n", t.Identifier, t.Title)
	}

	fmt.Printf("listing targets via CDP...\n")

	for _, t := range c.Targets() {
		fmt.Printf("  %s %s\n", t.TargetID, t.Title)
	}

	identifier := environment.Optional("CHROMIUM_TAB_ID")

	if identifier == "" {
		return
	}

	fmt.Printf("closing tab %s...\n", identifier)
	e := c.CloseTab(identifier)

	if e != nil {
		fmt.Printf("error: %s\n", e)

		return
	}

	fmt.Println("closed")
}
