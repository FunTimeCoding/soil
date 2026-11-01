package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"net/http"
	"strconv"
)

func (s *Server) reorder(
	w http.ResponseWriter,
	r *http.Request,
) {
	if e := r.ParseForm(); e != nil {
		http.Error(w, "invalid form", http.StatusBadRequest)

		return
	}

	values := r.Form[constant.ItemParameter]
	var identifiers []uint

	for _, v := range values {
		n, e := strconv.ParseUint(v, 10, 64)

		if e != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)

			return
		}

		identifiers = append(identifiers, uint(n))
	}

	s.service.Reorder(identifiers)
	w.WriteHeader(http.StatusNoContent)
}
