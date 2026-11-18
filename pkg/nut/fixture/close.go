package fixture

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Server) Close() {
	errors.LogClose(s.listener)
}
