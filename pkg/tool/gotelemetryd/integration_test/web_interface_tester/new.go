package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/integration_test/base"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := base.New(t)

	return &Tester{
		t:      t,
		server: s,
		base: locator.New(
			constant.Localhost,
		).Insecure().Port(s.ContextServer.Port).String(),
	}
}
