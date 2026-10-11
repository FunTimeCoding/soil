package aptly

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/basic"
)

func New(
	host string,
	port int,
	insecure bool,
	username string,
	password string,
) *Client {
	l := locator.New(host).Port(port)

	if insecure {
		l.Insecure()
	}

	result := requester.New(l)

	if username != "" {
		result.WithAuthorizer(basic.New(username, password))
	}

	return &Client{requester: result}
}
