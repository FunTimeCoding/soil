package request

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
	"net/url"
)

func PostForm(
	path string,
	form url.Values,
) *Request {
	return New(http.MethodPost, path).WithBody(
		constant.FormEncoded,
		[]byte(form.Encode()),
	)
}
