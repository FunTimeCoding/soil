package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func Tab() {
	c := chromium.NewEnvironment()
	defer c.Close()
	t := c.TabByHost(environment.Required("CHROMIUM_EXAMPLE_TAB"))

	if t == nil {
		panic("tab not found")
	}

	fmt.Printf("Body: %+v", c.Body(t.Identifier))
}
