package base

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (s *Server) MemoryClient() *client.Client {
	s.t.Helper()
	result := client.New(
		locator.New(webConstant.Localhost).Insecure().Port(s.Port),
		generativeConstant.ModelContextTestToken,
	)
	s.t.Cleanup(result.Close)

	return result
}
