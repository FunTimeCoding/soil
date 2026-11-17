package base

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (s *Server) Client() *client.Client {
	s.t.Helper()
	result := client.New(
		locator.New("localhost").Port(s.Served().Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	s.t.Cleanup(result.Close)

	return result
}
