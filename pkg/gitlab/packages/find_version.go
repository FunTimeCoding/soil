package packages

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"log"
)

func FindVersion(
	v []*gitlab.Package,
	version string,
) *gitlab.Package {
	if version == "" {
		log.Panicf("empty version")
	}

	errors.PanicOnSemver(version)

	if len(v) == 0 {
		return nil
	}

	for _, e := range v {
		errors.PanicOnSemver(e.Version)

		if e.Version == version {
			return e
		}
	}

	return nil
}
