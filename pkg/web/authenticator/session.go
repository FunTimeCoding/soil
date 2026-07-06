package authenticator

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/request_context"
)

func (a *Authenticator) Session(c *request_context.Context) string {
	if s := c.Cookie(constant.Session); s != nil {
		return s.Value
	}

	return ""
}
