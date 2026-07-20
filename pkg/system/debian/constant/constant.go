package constant

import "github.com/funtimecoding/soil/pkg/system/debian/release"

const (
	PackageConfigurationDirectory = "DEBIAN"

	PreseedConfiguration = "preseed.cfg"

	ControlFile = "control"

	PackageExtension = ".deb"

	Web   = "www.debian.org"
	Image = "cdimage.debian.org"
)

var Bookworm = release.New("bookworm", 12, 1)
