package release

import (
	"github.com/coreos/go-semver/semver"
	library "github.com/funtimecoding/soil/pkg/semver"
)

func LatestCompatible(
	v []*Release,
	other *semver.Version,
) *Release {
	for _, e := range v {
		s := semver.New(library.Trim(e.Name))

		if s.Major == other.Major && s.Minor == other.Minor {
			return e
		}
	}

	return nil
}
