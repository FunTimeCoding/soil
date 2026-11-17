package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goitermd/unit/service_tester"
	"testing"
)

func New(
	t *testing.T,
	option ...service_tester.Option,
) *Server {
	t.Helper()

	return &Server{Tester: service_tester.New(t, option...), t: t}
}
