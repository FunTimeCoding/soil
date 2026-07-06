package close_tab

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func History() {
	c := chromium.NewEnvironment()
	identifier := environment.Required("CHROMIUM_TAB_ID")
	fmt.Printf("reading navigation history for tab %s...\n", identifier)
	h, e := c.History(identifier)

	if e != nil {
		fmt.Printf("error: %s\n", e)

		return
	}

	fmt.Printf("current index: %d\n", h.CurrentIndex)

	for i, entry := range h.Entries {
		marker := "  "

		if int64(i) == h.CurrentIndex {
			marker = "→ "
		}

		fmt.Printf("%s[%d] %s\n     %s\n", marker, i, entry.Title, entry.URL)
	}
}
