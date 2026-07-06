package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func ObjectHeader(w http.ResponseWriter) {
	w.Header().Set(constant.ContentType, constant.Object)
}
