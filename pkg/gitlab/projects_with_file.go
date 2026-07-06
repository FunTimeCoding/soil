package gitlab

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"strings"
)

func (c *Client) ProjectsWithFile(
	path string,
	caseInsensitive bool,
) []*project.Project {
	var result []*project.Project

	if caseInsensitive {
		path = strings.ToLower(path)
	}

	for _, p := range c.MustProjects() {
		if c.verbose {
			fmt.Printf("Project: %s\n", p.Raw.NameWithNamespace)
		}

		for _, n := range c.MustTree(p.Identifier) {
			if path == n.Path ||
				(caseInsensitive && path == strings.ToLower(n.Path)) {
				result = append(result, p)

				break
			}
		}
	}

	return result
}
