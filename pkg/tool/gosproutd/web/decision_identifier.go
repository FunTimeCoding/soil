package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"net/http"
	"strconv"
)

func decisionIdentifier(
	w http.ResponseWriter,
	r *http.Request,
) (uint, bool) {
	identifier, e := strconv.ParseUint(
		r.URL.Query().Get(constant.IdentifierParameter),
		10,
		64,
	)

	if e != nil {
		http.Error(w, "invalid identifier", http.StatusBadRequest)

		return 0, false
	}

	return uint(identifier), true
}
