package service

import (
	"github.com/funtimecoding/soil/pkg/iterm/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"time"
)

type Service struct {
	client      face.ItermClient
	guard       *peer.Guard
	keyInterval time.Duration
}
