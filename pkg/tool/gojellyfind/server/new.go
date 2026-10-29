package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	jellyfin "github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"
)

func New(
	v jellyfin.Jellyfin,
	r face.Reporter,
) *Server {
	return &Server{service: v, reporter: r}
}
