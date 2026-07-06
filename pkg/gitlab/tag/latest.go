package tag

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"golang.org/x/mod/semver"
	"strings"
)

func Latest(v []*gitlab.Tag) *gitlab.Tag {
	result := v[0]

	for _, e := range v {
		// only consider tags with prefix
		if !strings.HasPrefix(e.Name, constant.VersionPrefix) {
			continue
		}

		if semver.Compare(e.Name, result.Name) > 0 {
			result = e
		}
	}

	return result
}
