package base

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func (s *Server) Client(t *testing.T) *client.Client {
	t.Helper()
	result := client.New(
		locator.New(webConstant.Localhost).Insecure().Port(s.Port).String(),
		false,
		generativeConstant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result
}
