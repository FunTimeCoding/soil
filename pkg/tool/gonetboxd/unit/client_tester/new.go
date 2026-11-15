package client_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(
	t *testing.T,
	s face.NetBoxClient,
) *client.Client {
	t.Helper()
	result := client.New(
		locator.New(webConstant.Localhost).Insecure().Port(
			base.New(t, s, memory.New()).Port,
		),
		generativeConstant.ModelContextTestToken,
	)
	t.Cleanup(result.Close)

	return result
}
