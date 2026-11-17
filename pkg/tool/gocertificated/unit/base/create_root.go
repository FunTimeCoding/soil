package base

import "github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"

func (s *Server) CreateRoot() string {
	s.t.Helper()

	return s.Tools().MustCallTool(constant.CreateAuthority, rootArgument())
}
