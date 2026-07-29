package confluence

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/user"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) User() (*user.User, error) {
	r, e := c.basic.GetPath(constant.ConfluenceUser)

	if e != nil {
		return nil, e
	}

	var result *response.User
	notation.MustDecode(r, &result, false)

	return user.New(result), nil
}
