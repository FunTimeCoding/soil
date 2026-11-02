package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"net/http"
	"testing"
)

func New(
	t *testing.T,
	configure func(*http.ServeMux),
) *Server {
	t.Helper()
	claude := base.New(t)
	t.Cleanup(claude.Close)

	return serve(t, configure, claude.Client(t))
}
