package gmail

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"golang.org/x/oauth2"
)

func (c *Client) exchange(
	o *oauth2.Config,
	code string,
) *oauth2.Token {
	result, e := o.Exchange(c.context, code)
	errors.PanicOnError(e)

	return result
}
