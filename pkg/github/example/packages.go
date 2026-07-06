package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/constant"
)

func Packages() {
	c := github.NewEnvironment()
	f := constant.Format

	for _, p := range c.MustPackages(constant.Namespace) {
		fmt.Printf("Package: %s\n", p.Format(f))

		for _, v := range c.MustPackageVersions(p.Name) {
			fmt.Printf("  Image: %s\n", v.Format(f))
		}
	}
}
