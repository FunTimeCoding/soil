package base

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Stack) Close() {
	errors.LogOnError(s.process.Kill())
	s.Site.CloseClientConnections()
	s.Site.Close()
}
