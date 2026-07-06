package web_service_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/integration_test/base"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := base.New(t)
	c, e := client.NewClientWithResponses(
		locator.New(
			constant.Localhost,
		).Insecure().Port(s.ContextServer.Port).String(),
	)
	assert.FatalOnError(t, e)

	return &Tester{server: s, Client: c}
}
