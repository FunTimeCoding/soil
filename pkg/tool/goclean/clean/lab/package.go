package lab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/packages"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func Package(
	x context.Context,
	c face.GitLabClient,
	p *project.Project,
) {
	projectPackages := c.MustPackages(x, p.Identifier, false)

	if len(projectPackages) > 0 {
		latest := packages.Latest(projectPackages)

		for _, a := range projectPackages {
			var isLatest bool

			for _, inner := range latest {
				if inner.Name == a.Name &&
					inner.Version == a.Version {
					isLatest = true
				}
			}

			if isLatest {
				continue
			}

			console.Format("Package: %s %s\n", a.Name, a.Version)
			c.MustDeletePackage(x, p.Identifier, a.ID)
		}
	}
}
