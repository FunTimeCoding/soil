package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
	"strconv"
)

func queryInteger(r *http.Request, name string) int64 {
	value := r.URL.Query().Get(name)

	if value == "" {
		return 0
	}

	result, e := strconv.ParseInt(value, 10, 64)
	errors.PanicOnError(e)

	return result
}
