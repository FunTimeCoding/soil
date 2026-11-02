package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
	"strings"
)

func (c *Client) FindRunnerByDescription(
	x context.Context,
	match string,
) (*runner.Runner, bool, error) {
	runners, e := c.Runners(x, true)

	if e != nil {
		return nil, false, e
	}

	for _, r := range runners {
		if strings.Contains(r.Description, match) {
			return r, true, nil
		}
	}

	return nil, false, nil
}
