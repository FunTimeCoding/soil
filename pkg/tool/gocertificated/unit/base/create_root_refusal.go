package base

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"

func (s *Server) CreateRootRefusal() string {
	s.t.Helper()

	return s.Tools().MustCallToolError(constant.CreateAuthority, rootArgument())
}
