package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func DefaultServer(h http.Handler) *http.Server {
	return Server(h, AddressPort(constant.ListenPort))
}
