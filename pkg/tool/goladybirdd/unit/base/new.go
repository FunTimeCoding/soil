package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/unit/service_tester"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()

	return &Server{Tester: service_tester.New(t), t: t}
}
