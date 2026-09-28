package server

import "github.com/funtimecoding/soil/pkg/tool/gogated/constant"

func (s *Server) defaultSource() string {
	if s.service.DirectoryConfigured() {
		return constant.SourceDirectory
	}

	return constant.SourceLocal
}
