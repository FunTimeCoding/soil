package lab

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
)

func Tag(
	c *gitlab.Client,
	p *project.Project,
) {
	tags := c.MustTags(p.Identifier)

	if len(tags) == 0 {
		return
	}

	latest := tag.Latest(tags)

	for _, t := range tags {
		if t.Name == latest.Name {
			continue
		}

		fmt.Printf("Delete tag: %s\n", t.Name)
		c.MustDeleteTag(p.Identifier, t.Name)
	}

	git.Fetch()
}
