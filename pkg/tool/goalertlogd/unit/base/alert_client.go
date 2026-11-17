package base

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (s *Server) AlertClient() *client.Client {
	s.t.Helper()
	result := client.New(
		locator.New(webConstant.Localhost).Insecure().Port(s.Served().Port),
		generativeConstant.ModelContextTestToken,
	)
	s.t.Cleanup(result.Close)

	return result
}
