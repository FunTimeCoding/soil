package confluence

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/basic"
	"github.com/funtimecoding/soil/pkg/errors"
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

	result.basic = basic.New(result.host, user, token, result.verbose)

	return result
}
