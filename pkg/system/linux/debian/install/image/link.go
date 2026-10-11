package image

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
)

func Link(
	v *semver.Version,
	architecture string,
	image string,
) string {
	return fmt.Sprintf("%s/%s", DirectoryLink(v, architecture), image)
}
