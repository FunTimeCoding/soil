package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic/response"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/notation"
	web "github.com/funtimecoding/soil/pkg/web/constant"
)

func (c *Client) searchV3Page(
	maximumResults int,
	nextPageToken string,
	query string,
) (*response.Search, error) {
	var result response.Search
	status, r, e := c.basic.Get(
		c.basic.Base().Copy().Base(constant.JiraBase).Path(constant.JiraSearch).Set(
			web.ParameterFields,
			constant.JiraAllFields,
		).SetInteger(constant.JiraMaximumResultsKey, maximumResults).Set(
			constant.JiraNextPageTokenKey,
			nextPageToken,
		).Set(constant.JiraQueryKey, query).Set(
			constant.JiraExpandKey,
			constant.JiraChangelogExpand,
		).String(),
	)

	if e != nil {
		return nil, e
	}

	notation.MustDecode(r, &result, true)

	if false {
		console.Format("Response: %d %s\n", status, r)
	}

	return &result, nil
}
