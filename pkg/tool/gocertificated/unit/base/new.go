package base

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/service_tester"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	option ...service_tester.Option,
) *Server {
	t.Helper()
	o := service_tester.New(t, option...)

	return &Server{
		Forge:    o.Forge,
		Reporter: o.Reporter,
		LitePath: filepath.Join(t.TempDir(), "gocertificated.db"),
		clock:    o.Clock,
		t:        t,
	}
}
