package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/constant"
)

func Packages() {
	c := github.NewEnvironment()
	f := constant.Format

	for _, p := range c.MustPackages(constant.Namespace) {
		console.Format("Package: %s\n", p.Format(f))

		for _, v := range c.MustPackageVersions(p.Name) {
			console.Format("  Image: %s\n", v.Format(f))
		}
	}
}
