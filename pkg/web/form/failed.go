package form

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func Failed(
	r *http.Request,
	identifier string,
) string {
	if r.URL.Query().Get(constant.FormParameter) != identifier {
		return ""
	}

	return ErrorText(r)
}
