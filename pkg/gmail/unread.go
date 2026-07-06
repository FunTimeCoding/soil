package gmail

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"google.golang.org/api/gmail/v1"
)

func (c *Client) Unread() *gmail.ListMessagesResponse {
	result, e := c.service.Users.Messages.List(
		"me",
	).Q("is:unread").MaxResults(10).Do()
	errors.PanicOnError(e)

	return result
}
