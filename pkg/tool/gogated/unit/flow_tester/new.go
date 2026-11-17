package flow_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"testing"
)

func New(
	t *testing.T,
	o ...base.Option,
) *Tester {
	t.Helper()
	s := base.New(t, o...)

	return &Tester{server: s, client: s.Served().Halting(), t: t}
}
