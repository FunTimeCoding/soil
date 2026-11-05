package service

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/sublime/face"
)

type Service struct {
	client face.SublimeClient
	guard  *peer.Guard
}
