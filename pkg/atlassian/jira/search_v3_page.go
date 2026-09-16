package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic/response"
	"github.com/funtimecoding/soil/pkg/notation"
	web "github.com/funtimecoding/soil/pkg/web/constant"
)

func (c *Client) searchV3Page(
	maximumResults int,
	nextPageToken string,
	query string,
	fields string,
	expand string,
) (*response.Search, error) {
	b := c.basic.Base().Copy().Base(constant.JiraBase).Path(
		constant.JiraSearch,
	).Set(
		web.ParameterFields,
		fields,
	).SetInteger(constant.JiraMaximumResultsKey, maximumResults).Set(
		constant.JiraNextPageTokenKey,
		nextPageToken,
	).Set(constant.JiraQueryKey, query)

	if expand != "" {
		b = b.Set(constant.JiraExpandKey, expand)
	}

	var result response.Search
	_, r, e := c.basic.Get(b.String())

	if e != nil {
		return nil, e
	}

	notation.MustDecode(r, &result, true)

	return &result, nil
}
