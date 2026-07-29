package client

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (c *Client) Subject(r *http.Request) string {
	cookie, e := r.Cookie(constant.AuthorizationSubjectCookie)

	if e != nil {
		return ""
	}

	b, e := c.decrypt(cookie.Value)

	if e != nil {
		return ""
	}

	return string(b)
}
