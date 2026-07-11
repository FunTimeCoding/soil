package form

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/url"
)

func Redirect(
	w http.ResponseWriter,
	r *http.Request,
	path string,
	message string,
	values url.Values,
) {
	q := url.Values{}

	for k, list := range values {
		for _, v := range list {
			q.Add(k, v)
		}
	}

	q.Set(constant.ErrorParameter, message)
	http.Redirect(
		w,
		r,
		fmt.Sprintf("%s?%s", path, q.Encode()),
		http.StatusSeeOther,
	)
}
