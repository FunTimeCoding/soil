package close_tab

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func Wake() {
	c := chromium.NewEnvironment()
	identifier := environment.Required("CHROMIUM_TAB_ID")
	console.Format("waking tab %s...\n", identifier)
	e := c.Wake(identifier)

	if e != nil {
		console.Format("error: %s\n", e)

		return
	}

	console.Line("woke")
}
