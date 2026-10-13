package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := base.New(t)
	t.Cleanup(s.Close)
	c := s.Connect(t)
	t.Cleanup(c.Close)

	return &Tester{Client: c, t: t, base: s}
}
