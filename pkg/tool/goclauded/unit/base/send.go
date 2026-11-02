package base

import "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"

func (s *Server) Send(
	name string,
	to string,
	body string,
) {
	s.t.Helper()
	s.deliver(name, to, body, client.Queued)
}
