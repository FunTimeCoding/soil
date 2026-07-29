package constant

import "github.com/funtimecoding/soil/pkg/system/debian/release"

const (
	DebianPackageConfigurationDirectory = "DEBIAN"

	PreseedConfiguration = "preseed.cfg"

	DebianControlFile = "control"

	DebianPackageExtension = ".deb"

	DebianWeb   = "www.debian.org"
	DebianImage = "cdimage.debian.org"
)

var Bookworm = release.New("bookworm", 12, 1)

const ChecksumFile = "SHA256SUMS"
