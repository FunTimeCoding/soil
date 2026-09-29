package gitlab

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"strings"
)

func (c *Client) ProjectsWithFile(
	path string,
	caseInsensitive bool,
) ([]*project.Project, error) {
	var result []*project.Project

	if caseInsensitive {
		path = strings.ToLower(path)
	}

	projects, e := c.Projects()

	if e != nil {
		return nil, e
	}

	for _, p := range projects {
		if c.verbose {
			console.Format("Project: %s\n", p.Raw.NameWithNamespace)
		}

		nodes, f := c.Tree(p.Identifier, "", "", false, 0)

		if f != nil {
			return nil, f
		}

		for _, n := range nodes {
			if path == n.Path ||
				(caseInsensitive && path == strings.ToLower(n.Path)) {
				result = append(result, p)

				break
			}
		}
	}

	return result, nil
}
