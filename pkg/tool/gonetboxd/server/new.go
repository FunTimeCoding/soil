package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store"
)

func New(
	c *netbox.Client,
	s *store.Store,
	r face.Reporter,
) *Server {
	return &Server{client: c, store: s, reporter: r}
}
