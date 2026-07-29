package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"time"
)

func SetCookie(
	w http.ResponseWriter,
	k string,
	v string,
) *http.Cookie {
	result := &http.Cookie{
		Name:     k,
		Value:    v,
		HttpOnly: true,
		Secure:   true,
		Path:     constant.LocationRoot,
		Expires:  time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, result)

	return result
}
