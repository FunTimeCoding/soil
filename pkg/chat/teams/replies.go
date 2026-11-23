package teams

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	microsoftConstant "github.com/funtimecoding/soil/pkg/microsoft/constant"
	"net/url"
)

func (c *Client) Replies(
	team string,
	channel string,
	identifier string,
	limit int,
) ([]*message.Message, error) {
	limit, top := pageSize(limit)

	return c.messages(
		fmt.Sprintf(
			constant.TeamsRepliesPath,
			url.PathEscape(team),
			url.PathEscape(channel),
			url.PathEscape(identifier),
		),
		url.Values{microsoftConstant.TopFilter: {top}},
		limit,
	)
}
