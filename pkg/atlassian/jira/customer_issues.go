package jira

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic/issue/issues"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/customer"
	"slices"
)

func (c *Client) CustomerIssues(all bool) ([]*customer.Issue, error) {
	v, r, e := c.service.Request.Gets(
		c.context,
		&models.ServiceRequestOptionScheme{},
		0,
		10,
	)

	if e != nil {
		return nil, e
	}

	list := issues.FromResponseScheme(r)
	var result []*customer.Issue

	for _, i := range customer.NewSlice(v.Values) {
		b := list.Key(i.Key)

		if b == nil {
			continue
		}

		i.Title = b.Title
		i.Body = b.Body

		if all {
			result = append(result, i)
		} else if !slices.Contains(constant.ServiceDeskDone, i.Status) {
			result = append(result, i)
		}
	}

	return result, nil
}
