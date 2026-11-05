package service

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Service struct {
	client face.JellyfinClient
	guard  *peer.Guard
}
