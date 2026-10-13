package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/base"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := base.New(t)

	return &Tester{
		t:          t,
		server:     s,
		Client:     s.ContextServer.Connect(t),
		MockClient: s.MockClient,
	}
}
