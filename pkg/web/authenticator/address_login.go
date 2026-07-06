package authenticator

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/request_context"
)

func (a *Authenticator) AddressLogin(c *request_context.Context) string {
	address := c.Address()

	if address != a.loginAddress {
		return ""
	}

	s := c.Cookie(constant.Session)

	if s == nil {
		identifier := a.store.Create(address)
		s = c.SetCookie(constant.Session, identifier)
	}

	return s.Value
}
