package base

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Server) Close() {
	s.Server.Close()
	errors.PanicClose(s.database)
	s.service.Close()
	s.indexer.Close()
}
