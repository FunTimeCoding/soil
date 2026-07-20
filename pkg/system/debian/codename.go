package debian

import (
	"github.com/coreos/go-semver/semver"
	"github.com/funtimecoding/soil/pkg/system/debian/constant"
	"log"
)

func Codename(v *semver.Version) string {
	switch v.Major {
	case constant.Bookworm.Major:
		return constant.Bookworm.Codename
	default:
		log.Panicf("unsupported major version %d", v.Major)
	}

	return ""
}
