package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"net/http"
)

func currentSessionIdentifier(r *http.Request) string {
	cookie, e := r.Cookie(constant.AuthenticationCookieName)

	if e != nil {
		return ""
	}

	return cookie.Value
}
