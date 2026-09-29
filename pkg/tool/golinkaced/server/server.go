package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	linkace "github.com/funtimecoding/soil/pkg/tool/golinkaced/face"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/service"
)

type Server struct {
	client   linkace.LinkAceSource
	service  *service.Service
	reporter face.Reporter
}
