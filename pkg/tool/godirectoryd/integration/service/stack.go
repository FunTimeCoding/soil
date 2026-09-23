package service

import (
	"github.com/funtimecoding/soil/pkg/directory"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
)

func stack() *service.Service {
	return service.New(directory.NewEnvironment())
}
