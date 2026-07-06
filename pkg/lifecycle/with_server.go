package lifecycle

import "github.com/funtimecoding/soil/pkg/lifecycle/server"

func WithServer(s *server.Server) Option {
	return func(l *Lifecycle) {
		l.component = append(l.component, s)
	}
}
