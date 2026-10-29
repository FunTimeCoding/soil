package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	jellyfin "github.com/funtimecoding/soil/pkg/tool/gojellyfind/face"
)

type Server struct {
	service  jellyfin.Jellyfin
	reporter face.Reporter
}
