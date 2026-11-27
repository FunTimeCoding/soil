package destination

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"net/url"
)

func (d *Destination) MembersPath() string {
	if d.Chat != "" {
		return fmt.Sprintf(
			constant.TeamsChatMembersPath,
			url.PathEscape(d.Chat),
		)
	}

	return fmt.Sprintf(
		constant.TeamsChannelMembersPath,
		url.PathEscape(d.Team),
		url.PathEscape(d.Channel),
	)
}
