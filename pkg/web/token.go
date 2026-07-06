package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func Token(
	r *http.Request,
	token string,
) {
	r.Header.Set(
		constant.Authorization,
		key_value.Space(constant.Token, token),
	)
}
