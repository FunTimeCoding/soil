package release

import (
	"fmt"
	"github.com/coreos/go-semver/semver"
)

func (r *Release) Version() *semver.Version {
	if r.version != nil {
		return r.version
	}

	r.version = semver.New(fmt.Sprintf("%d.%d.0", r.Major, r.Minor))

	return r.version
}
