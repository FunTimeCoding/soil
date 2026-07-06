package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/semver"
)

func LatestTags() {
	c := github.NewEnvironment()

	for _, p := range c.MustPackages(constant.LibraryNamespace) {
		var tags []string

		for _, v := range c.MustPackageVersions(p.Name) {
			tags = append(tags, v.Tags...)
		}

		latest := semver.Latest(tags)

		if latest == "" {
			fmt.Printf("%s: no semver tags\n", p.Name)

			continue
		}

		fmt.Printf("%s: %s\n", p.Name, latest)
	}
}
