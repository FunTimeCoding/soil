package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/base"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := base.New(t)
	s.SkipUnreachable(t)
	t.Cleanup(s.Close)
	c := s.Connect(t)
	t.Cleanup(c.Close)

	return &Tester{t: t, Client: c, base: s}
}
