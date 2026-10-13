package base

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Stack) Close() {
	s.Client.Close()
	errors.LogOnError(s.process.Kill())
	s.Site.CloseClientConnections()
	s.Site.Close()
	s.pool.Close()
}
