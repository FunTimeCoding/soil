package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func ClientAddress(r *http.Request) string {
	if s := r.Header.Get(constant.ForwardedFor); s != "" {
		return s
	}

	if s := r.Header.Get(constant.RealAddress); s != "" {
		return s
	}

	return r.RemoteAddr
}
