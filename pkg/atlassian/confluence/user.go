package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/user"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) User() (*user.User, error) {
	var result *response.User

	if e := c.basic.GetPath(constant.ConfluenceUser, &result); e != nil {
		return nil, e
	}

	return user.New(result), nil
}
