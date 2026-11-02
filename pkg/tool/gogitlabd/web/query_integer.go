package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"net/http"
	"strconv"
)

func queryInteger(
	r *http.Request,
	name string,
) int64 {
	value := r.URL.Query().Get(name)

	if value == "" {
		return 0
	}

	result, e := strconv.ParseInt(value, 10, 64)

	if e != nil {
		errors.PanicOnError(
			validation.New("%s is not a number: %s", name, value),
		)
	}

	return result
}
