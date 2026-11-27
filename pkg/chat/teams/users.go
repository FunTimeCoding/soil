package teams

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
	"github.com/funtimecoding/soil/pkg/chat/teams/user"
	microsoftConstant "github.com/funtimecoding/soil/pkg/microsoft/constant"
	"net/url"
	"strings"
)

func (c *Client) Users(
	query string,
	limit int,
) ([]*user.User, error) {
	limit, top := pageSize(limit)
	items, e := c.graph.Items(
		constant.TeamsUsersPath,
		url.Values{
			microsoftConstant.MatchFilter: {
				fmt.Sprintf(
					constant.TeamsUserMatch,
					strings.ReplaceAll(query, "'", "''"),
				),
			},
			microsoftConstant.SelectFilter: {constant.TeamsUserFields},
			microsoftConstant.TopFilter:    {top},
		},
		limit,
	)

	if e != nil {
		return nil, e
	}

	decoded, e := decodeAll[response.User](items)

	if e != nil {
		return nil, e
	}

	result := make([]*user.User, 0, len(decoded))

	for _, u := range decoded {
		result = append(result, toUser(&u))
	}

	return result, nil
}
