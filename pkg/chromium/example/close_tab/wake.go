package close_tab

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/system/environment"
)

func Wake() {
	c := chromium.NewEnvironment()
	identifier := environment.Required("CHROMIUM_TAB_ID")
	fmt.Printf("waking tab %s...\n", identifier)
	e := c.Wake(identifier)

	if e != nil {
		fmt.Printf("error: %s\n", e)

		return
	}

	fmt.Println("woke")
}
