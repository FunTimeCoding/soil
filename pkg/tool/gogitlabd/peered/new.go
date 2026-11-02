package peered

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	c face.GitLabClient,
	g *peer.Guard,
) *Peered {
	return &Peered{client: c, guard: g}
}
