package preseed

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web"
)

func Download(
	release string,
	configurationPath string,
) {
	if system.FileExists(configurationPath) {
		fmt.Printf("Configuration exists: %s\n", configurationPath)

		return
	}

	l := Link(release)
	fmt.Printf("Locator: %s\n", l)
	web.Download(l, configurationPath)
}
