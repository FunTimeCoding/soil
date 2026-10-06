package argocd

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/requester"
	"github.com/funtimecoding/soil/pkg/web/requester/authorizer/bearer"
)

func New(
	host string,
	port int,
	secure bool,
	token string,
) *Client {
	l := locator.New(host).Port(port)

	if !secure {
		l.Insecure()
	}

	return &Client{
		requester: requester.New(l).
			WithAuthorizer(bearer.New(token)).
			WithHeader(constant.Accept, constant.Object),
	}
}
