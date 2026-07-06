package web

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"net/http"
)

func InvalidRequestBody(w http.ResponseWriter) {
	http.Error(w, constant.InvalidRequestBody, http.StatusBadRequest)
}
