package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
	"slices"
)

func (c *Client) RunnersByTag(tag string) []*runner.Runner {
	var result []*runner.Runner

	for _, r := range c.MustRunners(true) {
		r = c.MustRunner(r.Identifier)

		if slices.Contains(r.Tags, tag) {
			result = append(result, r)
		}
	}

	return c.enrichRunners(result)
}
