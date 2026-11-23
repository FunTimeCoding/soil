package teams

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/message"
	microsoftConstant "github.com/funtimecoding/soil/pkg/microsoft/constant"
	"net/url"
)

func (c *Client) ChatMessages(
	chat string,
	limit int,
) ([]*message.Message, error) {
	limit, top := pageSize(limit)

	return c.messages(
		fmt.Sprintf(constant.TeamsChatMessagesPath, url.PathEscape(chat)),
		url.Values{
			microsoftConstant.TopFilter:   {top},
			microsoftConstant.OrderFilter: {constant.TeamsMessageOrder},
		},
		limit,
	)
}
