package base

import (
	generative "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/connector"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func (s *Server) Connector(t *testing.T) *connector.Client {
	t.Helper()

	return connector.New(
		locator.New(constant.Localhost).Insecure().Port(s.Port).String(),
		false,
		generative.ModelContextTestToken,
	)
}
