package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"net/http"
)

func setAuthenticationCookie(
	w http.ResponseWriter,
	identifier string,
) {
	http.SetCookie(
		w,
		&http.Cookie{
			Name:     constant.AuthenticationCookieName,
			Value:    identifier,
			Path:     "/",
			MaxAge:   int(constant.AuthenticationIdleTimeToLive.Seconds()),
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		},
	)
}
