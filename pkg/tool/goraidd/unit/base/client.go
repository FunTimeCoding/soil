package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidd/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (s *Server) Client(token string) *client.Client {
	s.t.Helper()
	result := client.New(
		locator.New("localhost").Port(s.Served().Port).Insecure(),
		token,
	)
	s.t.Cleanup(result.Close)

	return result
}
