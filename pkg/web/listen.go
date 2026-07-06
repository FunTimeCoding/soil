package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func Listen(m *http.ServeMux) {
	ListenAddress(m, constant.ListenAddress)
}
