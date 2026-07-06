package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/netbox"
)

func New(
	c *netbox.Client,
	r face.Reporter,
) *Server {
	return &Server{client: c, reporter: r}
}
