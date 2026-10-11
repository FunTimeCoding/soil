package base

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (s *Server) BaseLocator() string {
	return locator.New(constant.Localhost).Port(s.Web.Port).Insecure().String()
}
