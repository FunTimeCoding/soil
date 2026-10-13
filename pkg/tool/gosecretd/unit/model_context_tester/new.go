package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/unit/base"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := base.New(t)

	return &Tester{Server: s, Client: s.ContextServer.Connect(t)}
}
