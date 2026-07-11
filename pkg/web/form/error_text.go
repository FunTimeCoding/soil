package form

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func ErrorText(r *http.Request) string {
	return r.URL.Query().Get(constant.ErrorParameter)
}
