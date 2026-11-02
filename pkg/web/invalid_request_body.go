package web

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"net/http"
)

func InvalidRequestBody(w http.ResponseWriter) {
	http.Error(w, soil.InvalidRequestBody, http.StatusBadRequest)
}
