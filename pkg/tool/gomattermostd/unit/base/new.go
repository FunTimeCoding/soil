package base

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/base"
	"net/http"
	"testing"
)

func New(
	t *testing.T,
	configure func(*http.ServeMux),
	o ...Option,
) *Server {
	t.Helper()
	p := &setup{}

	for _, f := range o {
		f(p)
	}

	var k face.Client

	if p.connector {
		claude := base.New(t)
		t.Cleanup(claude.Close)
		k = claude.Client()
	}

	return serve(t, configure, k)
}
