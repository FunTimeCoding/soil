package tag

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/google/go-github/v89/github"
	"golang.org/x/mod/semver"
	"strings"
)

func Latest(v []*github.RepositoryTag) *github.RepositoryTag {
	var tags []string

	for _, t := range v {
		// Only consider tags with prefix
		if strings.HasPrefix(*t.Name, constant.VersionPrefix) {
			tags = append(tags, *t.Name)
		}
	}

	semver.Sort(tags)
	var latest *github.RepositoryTag
	index := len(tags) - 1

	for _, t := range v {
		if *t.Name == tags[index] {
			latest = t
		}
	}

	return latest
}
