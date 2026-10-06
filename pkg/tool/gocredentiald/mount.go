package gocredentiald

import (
	"github.com/funtimecoding/soil/pkg/face"
	credentialFace "github.com/funtimecoding/soil/pkg/tool/gocredentiald/face"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/model_context"
	"github.com/funtimecoding/soil/pkg/web/guard"
)

func Mount(
	s credentialFace.CredentialSource,
	r face.Reporter,
	t face.Recorder,
	g *guard.Mux,
) {
	model_context.New(s, r, t).Mount(g)
}
