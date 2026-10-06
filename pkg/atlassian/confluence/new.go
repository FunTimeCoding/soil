package confluence

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	host string,
	user string,
	token string,
	o ...Option,
) *Client {
	errors.FatalOnEmpty(host, "host")
	errors.FatalOnEmpty(user, "user")
	errors.FatalOnEmpty(token, "token")
	result := &Client{context: context.Background(), host: host}

	for _, f := range o {
		f(result)
	}

	result.basic = basic.New(locator.New(result.host), user, token)

	return result
}
