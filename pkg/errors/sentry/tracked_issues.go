package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
)

func (c *Client) TrackedIssues() ([]*issue.Issue, error) {
	var result []*issue.Issue

	if c.organization == "" {
		return result, nil
	}

	for _, p := range c.projects {
		proj, e := c.Project(c.organization, p)

		if e != nil {
			return nil, e
		}

		issues, f := c.Issues(
			c.organization,
			proj.ID,
			constant.PeriodFortnight,
		)

		if f != nil {
			return nil, f
		}

		result = append(result, issues...)
	}

	return result, nil
}
