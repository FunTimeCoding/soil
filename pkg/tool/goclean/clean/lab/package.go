package lab

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/gitlab/packages"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func Package(
	c *gitlab.Client,
	p *project.Project,
) {
	projectPackages := c.MustPackages(p.Identifier, false)

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

			fmt.Printf("Package: %s %s\n", a.Name, a.Version)
			c.MustDeletePackage(p.Identifier, a.ID)
		}
	}
}
