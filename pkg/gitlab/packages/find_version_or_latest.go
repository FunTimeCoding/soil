package packages

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"log"
)

func FindVersionOrLatest(
	v []*gitlab.Package,
	version string,
) *gitlab.Package {
	if version == "" {
		log.Panicf("empty version")
	}

	if len(v) == 0 {
		return nil
	}

	if version == constant.LatestVersion {
		return FindLatest(v)
	}

	return FindVersion(v, version)
}
