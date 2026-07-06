package request

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"net/http"
)

func PrivateToken(
	r *http.Request,
	token string,
) {
	r.Header.Add(constant.PrivateTokenHeader, token)
}
