package sentry

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
)

func (c *Client) AllIssues() ([]*issue.Issue, error) {
	organizations, e := c.Organizations()

	if e != nil {
		return nil, e
	}

	var result []*issue.Issue

	for _, o := range organizations {
		projects, f := c.OrganizationProjects(o.Slug)

		if f != nil {
			return nil, f
		}

		for _, p := range projects {
			issues, g := c.Issues(o.Slug, p.ID, constant.PeriodFortnight)

			if g != nil {
				return nil, g
			}

			result = append(result, issues...)
		}
	}

	return result, nil
}
