package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/unit/base"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := base.New(t)

	return &Tester{t: t, server: s, Client: s.Connect(t), Store: s.Store}
}
