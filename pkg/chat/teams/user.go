package teams

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/response"
	"github.com/funtimecoding/soil/pkg/chat/teams/user"
	microsoftConstant "github.com/funtimecoding/soil/pkg/microsoft/constant"
	"net/url"
)

func (c *Client) User(identifier string) (*user.User, error) {
	var result response.User

	if e := c.graph.Get(
		fmt.Sprintf(constant.TeamsUserPath, url.PathEscape(identifier)),
		url.Values{microsoftConstant.SelectFilter: {constant.TeamsUserFields}},
		&result,
	); e != nil {
		return nil, e
	}

	return toUser(&result), nil
}
