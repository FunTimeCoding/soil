package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/web/parameter"
)

func (c *Client) searchV3Page(
	maximumResults int,
	nextPageToken string,
	query string,
) (*response.Search, error) {
	var result response.Search
	status, r, e := c.basic.Get(
		c.basic.Base().Copy().Base(constant.Base).Path(constant.Search).Set(
			parameter.Fields,
			constant.AllFields,
		).SetInteger(constant.MaximumResultsKey, maximumResults).Set(
			constant.NextPageTokenKey,
			nextPageToken,
		).Set(constant.QueryKey, query).Set(
			constant.ExpandKey,
			constant.ChangelogExpand,
		).String(),
	)

	if e != nil {
		return nil, e
	}

	notation.MustDecode(r, &result, true)

	if false {
		fmt.Printf("Response: %d %s\n", status, r)
	}

	return &result, nil
}
