package confluence

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/search_result"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (c *Client) Search(
	query string,
	a ...any,
) ([]*search_result.Result, error) {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	var result *response.Search

	if e := c.basic.Get(
		locator.New(c.host).Base(constant.ConfluenceOldBase).Path(
			constant.ConfluenceSearch,
		).Set(constant.ConfluenceQuery, query).String(),
		&result,
	); e != nil {
		return nil, e
	}

	return search_result.NewSlice(result.Results), nil
}
