package service

import (
	"github.com/funtimecoding/soil/pkg/github/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	c face.GitHubClient,
	g *peer.Guard,
	owner string,
) *Service {
	return &Service{client: c, guard: g, owner: owner}
}
