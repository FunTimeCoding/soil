package confluence

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/search_result"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (c *Client) Search(
	query string,
	a ...any,
) ([]*search_result.Result, error) {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	body, e := c.basic.Get(
		locator.New(c.host).Base(constant.OldBase).Path(
			constant.Search,
		).Set(constant.Query, query).String(),
	)

	if e != nil {
		return nil, e
	}

	var result *response.Search
	notation.MustDecode(body, &result, false)

	return search_result.NewSlice(result.Results), nil
}
