package gmail

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/gmail/constant"
	keyValueJoin "github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"github.com/funtimecoding/soil/pkg/system/join"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
	"strings"
)

func (c *Client) loadService() *gmail.Service {
	token := c.tokenFlow(c.credential, false)
	result, e := gmail.NewService(
		c.context,
		option.WithTokenSource(
			c.credential.TokenSource(c.context, token),
		),
	)
	errors.PanicOnError(e)
	name, _ := key_value.At(c.profile(result).EmailAddress)
	c.saveToken(
		token,
		join.Absolute(
			c.directory,
			keyValueJoin.Empty(strings.ToLower(name), constant.TokenSuffix),
		),
	)

	return result
}
