package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/base"
	"testing"
)

func New(
	t *testing.T,
	fixture string,
) *Tester {
	t.Helper()
	s := base.New(t, fixture)

	return &Tester{
		t:         t,
		server:    s,
		Client:    s.Connect(t),
		Directory: s.Directory,
	}
}
