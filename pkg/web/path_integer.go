package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"net/http"
	"strconv"
)

func PathInteger(
	w http.ResponseWriter,
	r *http.Request,
	name string,
) (int, bool) {
	result, e := strconv.Atoi(r.PathValue(name))

	if e != nil {
		http.Error(w, join.Space("invalid", name), http.StatusBadRequest)

		return 0, false
	}

	return result, true
}
