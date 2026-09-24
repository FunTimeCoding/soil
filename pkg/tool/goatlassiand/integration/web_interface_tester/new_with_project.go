package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/integration/base"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func NewWithProject(
	t *testing.T,
	project []string,
) *Tester {
	t.Helper()
	s := base.NewWithProject(t, project)

	return &Tester{
		t:      t,
		server: s,
		base: locator.New(constant.Localhost).Insecure().Port(
			s.Port,
		).String(),
	}
}
