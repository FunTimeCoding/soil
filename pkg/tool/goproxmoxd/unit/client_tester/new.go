package client_tester

import (
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/client"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(t *testing.T) (*client.Client, *base.Server) {
	t.Helper()
	s := base.New(t)
	result := client.New(
		locator.New(webConstant.Localhost).Insecure().Port(s.Port),
		generativeConstant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result, s
}
