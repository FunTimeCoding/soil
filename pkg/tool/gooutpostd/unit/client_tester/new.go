package client_tester

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/client"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(
	t *testing.T,
	s *base.Server,
) *client.Client {
	t.Helper()
	result := client.New(
		locator.New(webConstant.Localhost).Port(s.Served().Port).Insecure(),
		generativeConstant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result
}
