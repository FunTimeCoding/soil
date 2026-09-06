package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (s *Server) address(path string) string {
	return locator.New(constant.Localhost).Insecure().Port(s.Port).Path(
		path,
	).String()
}
