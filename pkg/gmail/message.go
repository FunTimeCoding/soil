package gmail

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"google.golang.org/api/gmail/v1"
)

func (c *Client) Message(m *gmail.Message) *gmail.Message {
	result, e := c.service.Users.Messages.Get("me", m.Id).Do()
	errors.PanicOnError(e)

	return result
}
