package server

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gopostgresd/service"
)

type Server struct {
	service  *service.Service
	reporter face.Reporter
}
