package form

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/url"
)

func RedirectFor(
	w http.ResponseWriter,
	r *http.Request,
	path string,
	identifier string,
	message string,
	values url.Values,
) {
	carried := url.Values{constant.FormParameter: {identifier}}

	for k, list := range values {
		carried[k] = list
	}

	Redirect(w, r, path, message, carried)
}
