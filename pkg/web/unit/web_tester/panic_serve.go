package web_tester

import "net/http"

func PanicServe(
	http.ResponseWriter,
	*http.Request,
) {
	panic("store failed")
}
