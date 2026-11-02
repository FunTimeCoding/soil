package base

import (
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (s *Server) BaseLocator() string {
	return locator.New(webConstant.Localhost).Port(s.Web.Port).Insecure().String()
}
