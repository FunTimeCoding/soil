package service

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

func New(
	c face.SentryClient,
	g *peer.Guard,
	organization string,
) *Service {
	return &Service{client: c, guard: g, organization: organization}
}
