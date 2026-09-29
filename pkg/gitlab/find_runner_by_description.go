package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
	"strings"
)

func (c *Client) FindRunnerByDescription(
	match string,
) (*runner.Runner, bool, error) {
	runners, e := c.Runners(true)

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
