package package_server

import (
	"log"
	"net/http"
)

func httpFail(
	w http.ResponseWriter,
	message string,
	e error,
) {
	log.Printf("%s: %v", message, e)
	http.Error(w, message, http.StatusInternalServerError)
}
