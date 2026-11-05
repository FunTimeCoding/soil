package service

import (
	"github.com/funtimecoding/soil/pkg/iterm/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"time"
)

func New(
	c face.ItermClient,
	g *peer.Guard,
	keyInterval time.Duration,
) *Service {
	return &Service{client: c, guard: g, keyInterval: keyInterval}
}
