package base

import "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"

func (s *Server) SendImmediate(
	name string,
	to string,
	body string,
) {
	s.t.Helper()
	s.deliver(name, to, body, client.Immediate)
}
