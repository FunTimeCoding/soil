package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func TextHeader(w http.ResponseWriter) {
	w.Header().Set(constant.ContentType, constant.Text)
}
