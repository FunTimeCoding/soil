package image

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func DirectoryLink(
	v *semver.Version,
	architecture string,
) string {
	return locator.New(constant.DebianImage).Path(
		"/cdimage/release/%s/%s/iso-cd",
		v.String(),
		architecture,
	).String()
}
