package example

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func Tab() {
	c := chromium.NewEnvironment()
	defer c.Close()
	t := c.TabByHost(environment.Required("CHROMIUM_EXAMPLE_TAB"))

	if t == nil {
		panic("tab not found")
	}

	console.Format("Body: %+v", c.Body(t.Identifier))
}
