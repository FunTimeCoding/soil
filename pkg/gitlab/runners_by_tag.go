package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
	"slices"
)

func (c *Client) RunnersByTag(tag string) ([]*runner.Runner, error) {
	var result []*runner.Runner
	runners, e := c.Runners(true)

	if e != nil {
		return nil, e
	}

	for _, r := range runners {
		detail, f := c.Runner(r.Identifier)

		if f != nil {
			return nil, f
		}

		if slices.Contains(detail.Tags, tag) {
			result = append(result, detail)
		}
	}

	return result, nil
}
