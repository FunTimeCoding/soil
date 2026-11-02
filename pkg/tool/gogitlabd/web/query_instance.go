package web

import (
	inventoryConstant "github.com/funtimecoding/soil/pkg/inventory/constant"
	"net/http"
)

func (s *Server) queryInstance(r *http.Request) (string, error) {
	return s.inventory.ResolveInstance(
		r.URL.Query().Get(inventoryConstant.InstanceParameter),
	)
}
