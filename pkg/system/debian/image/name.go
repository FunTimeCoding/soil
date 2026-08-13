package image

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
)

func Name(
	v *semver.Version,
	architecture string,
) string {
	return fmt.Sprintf("debian-%s-%s-netinst.iso", v.String(), architecture)
}
