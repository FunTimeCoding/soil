package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
)

func New(
	c netbox.NetboxSource,
	s *store.Store,
	r face.Reporter,
) *Server {
	return &Server{client: c, store: s, reporter: r}
}
