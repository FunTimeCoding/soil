package request

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"net/http"
)

func JobToken(
	r *http.Request,
	token string,
) {
	r.Header.Add(constant.JobTokenHeader, token)
}
