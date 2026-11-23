package teams

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	microsoftConstant "github.com/funtimecoding/soil/pkg/microsoft/constant"
	"net/url"
)

func (c *Client) ChannelMessages(
	team string,
	channel string,
	limit int,
) ([]*message.Message, error) {
	limit, top := pageSize(limit)

	return c.messages(
		fmt.Sprintf(
			constant.TeamsChannelMessagesPath,
			url.PathEscape(team),
			url.PathEscape(channel),
		),
		url.Values{
			microsoftConstant.TopFilter:    {top},
			microsoftConstant.ExpandFilter: {constant.TeamsRepliesExpansion},
		},
		limit,
	)
}
