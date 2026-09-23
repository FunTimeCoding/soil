package close_tab

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func History() {
	c := chromium.NewEnvironment()
	defer c.Close()
	identifier := environment.Required("CHROMIUM_TAB_ID")
	console.Format("reading navigation history for tab %s...\n", identifier)
	h, e := protocol.NewIdentifier(c, identifier).History()

	if e != nil {
		console.Format("error: %s\n", e)

		return
	}

	console.Format("current index: %d\n", h.CurrentIndex)

	for i, entry := range h.Entries {
		marker := "  "

		if int64(i) == h.CurrentIndex {
			marker = "→ "
		}

		console.Format(
			"%s[%d] %s\n     %s\n",
			marker,
			i,
			entry.Title,
			entry.Locator,
		)
	}
}
