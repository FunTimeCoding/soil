package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/store"
)

type Server struct {
	store      *store.Store
	outputPath string
	reporter   face.Reporter
}
