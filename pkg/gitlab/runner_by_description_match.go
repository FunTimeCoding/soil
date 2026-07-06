package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
	"strings"
)

func (c *Client) RunnerByDescriptionMatch(s string) *runner.Runner {
	for _, r := range c.MustRunners(true) {
		if strings.Contains(r.Description, s) {
			return r
		}
	}

	return nil
}
