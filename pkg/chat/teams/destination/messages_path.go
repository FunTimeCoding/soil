package destination

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"net/url"
)

func (d *Destination) MessagesPath() string {
	switch {
	case d.Chat != "":
		return fmt.Sprintf(
			constant.TeamsChatMessagesPath,
			url.PathEscape(d.Chat),
		)
	case d.Thread != "":
		return fmt.Sprintf(
			constant.TeamsRepliesPath,
			url.PathEscape(d.Team),
			url.PathEscape(d.Channel),
			url.PathEscape(d.Thread),
		)
	default:
		return fmt.Sprintf(
			constant.TeamsChannelMessagesPath,
			url.PathEscape(d.Team),
			url.PathEscape(d.Channel),
		)
	}
}
