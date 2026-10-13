package web

import (
	"github.com/funtimecoding/soil/pkg/inventory/constant"
	"net/http"
)

func (s *Server) queryInstance(r *http.Request) (string, error) {
	return s.service.Inventory().ResolveInstance(
		r.URL.Query().Get(constant.InstanceParameter),
	)
}
