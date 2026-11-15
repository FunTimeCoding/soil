package web_tester

import "net/http"

func Serve(
	w http.ResponseWriter,
	_ *http.Request,
) {
	w.WriteHeader(http.StatusOK)
}
