package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/option"
	"net/http"
)

func New(
	i *identity.Tool,
	w *option.Web,
	r face.Reporter,
	setup func(*guard.Mux),
) *Server {
	result := &Server{
		Mux:      http.NewServeMux(),
		Address:  w.Address,
		identity: i,
		release:  func() {},
	}
	result.Setup = func(m *http.ServeMux) {
		g, release := w.Guard(m, r)
		result.release = release
		setup(g)
	}

	return result
}
