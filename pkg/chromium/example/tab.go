package example

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/chromium/protocol"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func Tab() {
	c := chromium.NewEnvironment()
	defer c.Close()
	console.Format(
		"Body: %+v",
		protocol.New(c, environment.Required("CHROMIUM_EXAMPLE_TAB")).Body(),
	)
}
