package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	jellyfin "github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"
)

func New(
	c jellyfin.JellyfinSource,
	r face.Reporter,
) *Server {
	return &Server{client: c, reporter: r}
}
