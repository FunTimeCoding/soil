package teams

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
	"github.com/funtimecoding/soil/pkg/chat/teams/user"
	microsoftConstant "github.com/funtimecoding/soil/pkg/microsoft/constant"
	"net/url"
)

func (c *Client) Me() (*user.User, error) {
	var result response.User

	if e := c.graph.Get(
		microsoftConstant.ProfilePath,
		url.Values{microsoftConstant.SelectFilter: {constant.TeamsUserFields}},
		&result,
	); e != nil {
		return nil, e
	}

	return toUser(&result), nil
}
