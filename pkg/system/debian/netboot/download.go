package netboot

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web"
)

func Download(
	release string,
	architecture string,
	archivePath string,
) {
	if system.FileExists(archivePath) {
		fmt.Printf("Archive exists: %s\n", archivePath)

		return
	}

	l := Link(release, architecture)
	fmt.Printf("Locator: %s\n", l)
	web.Download(l, archivePath)
}
