package guard

import "net/http"

func reachable(
	w http.ResponseWriter,
	_ *http.Request,
) {
	w.WriteHeader(http.StatusNoContent)
}
