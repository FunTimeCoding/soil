package base

import (
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/unit/service_tester"
	"testing"
)

func New(
	t *testing.T,
	o ...service_tester.Option,
) *Server {
	t.Helper()

	return &Server{Tester: service_tester.New(t, o...), t: t}
}
