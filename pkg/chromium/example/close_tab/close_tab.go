package close_tab

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func CloseTab() {
	c := chromium.NewEnvironment()
	console.Line("listing tabs via HTTP...")

	for _, t := range c.Tabs() {
		if t.Type != constant.PageTabType {
			continue
		}

		console.Format("  %s %s\n", t.Identifier, t.Title)
	}

	console.Format("listing targets via CDP...\n")

	for _, t := range c.Targets() {
		console.Format("  %s %s\n", t.TargetID, t.Title)
	}

	identifier := environment.Optional("CHROMIUM_TAB_ID")

	if identifier == "" {
		return
	}

	console.Format("closing tab %s...\n", identifier)
	e := c.CloseTab(identifier)

	if e != nil {
		console.Format("error: %s\n", e)

		return
	}

	console.Line("tab closed")
}
