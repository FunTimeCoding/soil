package service

import (
	"github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Service struct {
	client face.TeamsClient
	guard  *peer.Guard
}
