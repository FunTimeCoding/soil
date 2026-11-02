package lab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/image"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"strings"
)

func Registry(
	x context.Context,
	c face.GitLabClient,
	p *project.Project,
) {
	for _, r := range c.MustRegistryRepositories(x, p.Identifier, false) {
		images := c.MustImages(x, p.Identifier, r.Identifier)

		if len(images) == 0 {
			continue
		}

		if strings.HasSuffix(r.Name, "/cache") {
			for _, i := range images {
				console.Format("Image: %s\n", i.Name)
				c.MustDeleteImage(x, p.Identifier, r.Identifier, i.Name)
			}

			continue
		}

		latest := image.Latest(images)

		for _, i := range images {
			if strings.HasSuffix(i.Path, gitlabConstant.LatestSuffix) {
				continue
			}

			if i.Path == latest.Path {
				continue
			}

			console.Format("Image: %s\n", i.Name)
			c.MustDeleteImage(x, p.Identifier, r.Identifier, i.Name)
		}
	}
}
