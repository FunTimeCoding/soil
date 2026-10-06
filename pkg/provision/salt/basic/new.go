package basic

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/provision/salt/session"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	base *locator.Locator,
	user string,
	password string,
	eauth string,
) *Client {
	s := session.New(newRequester(base), user, password, eauth)
	errors.PanicOnError(s.Renew())

	return &Client{
		requester: newRequester(base).
			WithAuthorizer(s).
			WithHeader(constant.Accept, constant.Object),
	}
}
