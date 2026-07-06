package hub

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/tag"
	"github.com/funtimecoding/soil/pkg/system"
)

func Tag(
	c *github.Client,
	namespace string,
	repository string,
) {
	tags := c.MustTags(namespace, repository)
	latestTag := tag.Latest(tags)

	if latestTag == nil {
		system.Exitf(
			1,
			"could not find latest tag for %s/%s\n",
			namespace,
			repository,
		)

		return
	}

	for _, t := range tags {
		if t.Name == latestTag.Name {
			continue
		}

		fmt.Printf("Delete tag: %s\n", *t.Name)
		c.MustDeleteTag(namespace, repository, *t.Name)
	}

	git.Fetch()
}
