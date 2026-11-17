package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goagentd/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (s *Server) AgentClient(token string) *client.Client {
	s.t.Helper()
	result := client.New(
		locator.New("localhost").Port(s.Served().Port).Insecure(),
		token,
	)
	s.t.Cleanup(result.Close)

	return result
}
