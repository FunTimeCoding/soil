package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/user"
	"github.com/funtimecoding/soil/pkg/errors"
	rawUser "github.com/opsgenie/opsgenie-go-sdk-v2/user"
)

func (c *Client) Users() []*user.User {
	result, e := c.userClient.User.List(c.context, &rawUser.ListRequest{})
	errors.PanicOnError(e)

	return user.NewSlice(result.Users)
}
