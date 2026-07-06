package hub

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/github"
	"github.com/funtimecoding/soil/pkg/github/image"
)

func Image(
	c *github.Client,
	namespace string,
	repository string,
) {
	for _, p := range c.MustPackages(namespace) {
		if p.Repository != repository {
			continue
		}

		for _, v := range c.MustPackageVersions(p.Name) {
			if image.HasLatest(v) {
				continue
			}

			fmt.Printf("Delete image: %s@%s\n", p.Name, v.Digest)
			c.MustDeletePackageVersion(p.Name, v.Identifier)
		}
	}
}
