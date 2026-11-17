package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"
	"testing"
)

func New(
	t *testing.T,
	o ...service_tester.Option,
) *Server {
	t.Helper()
	result := &Server{Tester: service_tester.New(t, o...), t: t}
	t.Cleanup(result.Close)

	return result
}
