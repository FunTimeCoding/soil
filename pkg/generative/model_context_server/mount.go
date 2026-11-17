package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func Mount(
	t *testing.T,
	d face.Mountable,
	r face.Reporter,
	o ...Option,
) *Server {
	t.Helper()

	return New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
		o...,
	)
}
