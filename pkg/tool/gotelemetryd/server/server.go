package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

type Server struct {
	store    *store.Store
	reporter face.Reporter
}
