package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func ListenAsynchronous(m *http.ServeMux) *http.Server {
	s := Server(m, AddressPort(constant.ListenPort))
	ServeAsynchronous(s)

	return s
}
