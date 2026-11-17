package client_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/client"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(
	t *testing.T,
	s *base.Server,
	token string,
) *client.Client {
	t.Helper()
	result := client.New(
		locator.New(webConstant.Localhost).Port(s.Served().Port).Insecure(),
		token,
	)
	t.Cleanup(result.Close)

	return result
}
