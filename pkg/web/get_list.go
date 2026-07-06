package web

import (
	"github.com/funtimecoding/soil/pkg/strings/split"
	"net/http"
)

func GetList(
	r *http.Request,
	key string,
) []string {
	var result []string

	if s := r.URL.Query().Get(key); s != "" {
		result = split.Comma(s)
	}

	return result
}
