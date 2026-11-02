package lab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/git"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
)

func Tag(
	x context.Context,
	c face.GitLabClient,
	p *project.Project,
) {
	tags := c.MustTags(x, p.Identifier)

	if len(tags) == 0 {
		return
	}

	latest := tag.Latest(tags)

	for _, t := range tags {
		if t.Name == latest.Name {
			continue
		}

		console.Format("Delete tag: %s\n", t.Name)
		c.MustDeleteTag(x, p.Identifier, t.Name)
	}

	git.Fetch()
}
