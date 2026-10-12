package base

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func (s *Server) Connector(t *testing.T) *connector.Client {
	t.Helper()

	return connector.New(
		locator.New(webConstant.Localhost).Insecure().Port(s.Port).String(),
		false,
		generativeConstant.ModelContextTestToken,
	)
}
