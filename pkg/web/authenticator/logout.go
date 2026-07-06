package authenticator

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/request_context"
)

func (a *Authenticator) Logout(c *request_context.Context) {
	if s := c.Cookie(constant.Session); s != nil {
		a.store.Remove(s.Value)
		c.UnsetCookie(constant.Session)
	}
}
