package checksum

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/soil/pkg/system/debian/image"
)

func Link(
	v *semver.Version,
	architecture string,
) string {
	return fmt.Sprintf(
		"%s/%s",
		image.DirectoryLink(v, architecture),
		File,
	)
}
