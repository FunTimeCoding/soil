package lifecycle

import "github.com/funtimecoding/soil/pkg/lifecycle/server"

func (l *Lifecycle) WithServer(s *server.Server) *Lifecycle {
	l.component = append(l.component, s)

	return l
}
