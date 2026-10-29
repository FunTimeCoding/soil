package service

import (
	"github.com/funtimecoding/soil/pkg/linkace/face"
	"github.com/funtimecoding/soil/pkg/peer"
)

type Service struct {
	client face.Client
	guard  *peer.Guard
}
