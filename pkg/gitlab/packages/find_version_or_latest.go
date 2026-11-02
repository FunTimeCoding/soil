package packages

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"gitlab.com/gitlab-org/api/client-go/v3"
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

	if version == soil.LatestVersion {
		return FindLatest(v)
	}

	return FindVersion(v, version)
}
