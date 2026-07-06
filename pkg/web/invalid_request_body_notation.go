package web

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"net/http"
)

func InvalidRequestBodyNotation(w http.ResponseWriter) {
	EncodeError(w, http.StatusBadRequest, constant.InvalidRequestBody)
}
