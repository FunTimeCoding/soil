package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
	"slices"
)

func (c *Client) TagRunners(
	x context.Context,
	tag string,
) ([]*runner.Runner, error) {
	var result []*runner.Runner
	runners, e := c.Runners(x, true)

	if e != nil {
		return nil, e
	}

	for _, r := range runners {
		detail, f := c.Runner(x, r.Identifier)

		if f != nil {
			return nil, f
		}

		if slices.Contains(detail.Tags, tag) {
			result = append(result, detail)
		}
	}

	return result, nil
}
