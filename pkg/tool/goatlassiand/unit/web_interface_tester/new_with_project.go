package web_interface_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/base"
	"testing"
)

func NewWithProject(
	t *testing.T,
	project []string,
) *Tester {
	t.Helper()
	s := base.NewWithProject(t, project)

	return &Tester{t: t, server: s}
}
