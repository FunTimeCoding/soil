package sentry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/soil/pkg/errors/sentry/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/issue"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"slices"
)

func (c *Client) IssuesSimple(verbose bool) ([]*issue.Issue, error) {
	if o := environment.Optional(constant.OrganizationEnvironment); o != "" {
		var projects []response.Project

		if p := environment.Optional(constant.ProjectEnvironment); p != "" {
			names := split.Comma(p)
			all, e := c.OrganizationProjects(o)

			if e != nil {
				return nil, e
			}

			for _, j := range all {
				if slices.Contains(names, j.Name) {
					projects = append(projects, j)
				}
			}
		} else {
			all, e := c.OrganizationProjects(o)

			if e != nil {
				return nil, e
			}

			projects = append(projects, all...)
		}

		var result []*issue.Issue

		for _, p := range projects {
			if verbose {
				fmt.Printf("Project: %s\n", p.Name)
			}

			issues, e := c.Issues(o, p.ID, constant.PeriodFortnight)

			if e != nil {
				return nil, e
			}

			result = append(result, issues...)
		}

		return result, nil
	}

	return c.AllIssues()
}
