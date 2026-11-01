package form

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func PageError(r *http.Request) string {
	if r.URL.Query().Has(constant.FormParameter) {
		return ""
	}

	return ErrorText(r)
}
