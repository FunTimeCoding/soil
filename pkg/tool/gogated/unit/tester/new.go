package tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := base.New(t)

	return &Tester{server: s, client: s.Web.Halting()}
}
