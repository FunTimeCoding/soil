package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store"
)

type Server struct {
	store    *store.Store
	reporter face.Reporter
}
